package network

import (
	"context"
	"time"

	"github.com/shirou/gopsutil/v3/net"
)

// Monitor handles network interface monitoring
type Monitor struct {
	lastBytesRecv uint64
	lastBytesSent uint64
	lastTime      time.Time
	interfaceName string
}

// Stats represents network statistics
type Stats struct {
	BytesRecv uint64
	BytesSent uint64
	Time      time.Time
}

// NewMonitor creates a new network monitor
func NewMonitor() (*Monitor, error) {
	m := &Monitor{
		lastTime: time.Now(),
	}

	// Get initial stats to detect active interface
	stats, err := m.getNetworkStats()
	if err != nil {
		return nil, err
	}

	m.lastBytesRecv = stats.BytesRecv
	m.lastBytesSent = stats.BytesSent
	m.lastTime = stats.Time

	return m, nil
}

// getNetworkStats retrieves network statistics from the primary active interface
func (m *Monitor) getNetworkStats() (*Stats, error) {
	interfaces, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}

	if len(interfaces) == 0 {
		return nil, &NoInterfaceError{Message: "no network interfaces found"}
	}

	var totalRecv, totalSent uint64
	var activeInterface string

	// Sum all non-loopback interfaces
	for _, iface := range interfaces {
		// Skip loopback interfaces
		if iface.Name == "lo" || iface.Name == "Loopback" || iface.Name == "lo0" {
			continue
		}

		// Prefer interfaces with actual traffic
		if iface.BytesRecv > 0 || iface.BytesSent > 0 {
			totalRecv += iface.BytesRecv
			totalSent += iface.BytesSent
			if activeInterface == "" {
				activeInterface = iface.Name
			}
		}
	}

	// If no active interface found, use the first non-loopback interface
	if activeInterface == "" {
		for _, iface := range interfaces {
			if iface.Name != "lo" && iface.Name != "Loopback" && iface.Name != "lo0" {
				totalRecv = iface.BytesRecv
				totalSent = iface.BytesSent
				activeInterface = iface.Name
				break
			}
		}
	}

	// If still no interface found, return error
	if activeInterface == "" {
		return nil, &NoInterfaceError{Message: "no active network interface found"}
	}

	m.interfaceName = activeInterface

	return &Stats{
		BytesRecv: totalRecv,
		BytesSent: totalSent,
		Time:      time.Now(),
	}, nil
}

// NoInterfaceError represents an error when no network interface is available
type NoInterfaceError struct {
	Message string
}

func (e *NoInterfaceError) Error() string {
	return e.Message
}

// GetSpeeds calculates download and upload speeds in bytes per second
func (m *Monitor) GetSpeeds() (downloadBps uint64, uploadBps uint64, err error) {
	currentStats, err := m.getNetworkStats()
	if err != nil {
		return 0, 0, err
	}

	// Calculate time delta
	timeDelta := currentStats.Time.Sub(m.lastTime).Seconds()
	if timeDelta <= 0 {
		timeDelta = 1.0 // Default to 1 second if time hasn't advanced
	}

	// Calculate byte deltas with rollover protection
	var recvDelta, sentDelta uint64

	// Handle counter rollover (if current < last, assume rollover)
	if currentStats.BytesRecv >= m.lastBytesRecv {
		recvDelta = currentStats.BytesRecv - m.lastBytesRecv
	} else {
		// Rollover occurred, calculate assuming 64-bit counter
		recvDelta = (^uint64(0) - m.lastBytesRecv) + currentStats.BytesRecv + 1
	}

	if currentStats.BytesSent >= m.lastBytesSent {
		sentDelta = currentStats.BytesSent - m.lastBytesSent
	} else {
		// Rollover occurred
		sentDelta = (^uint64(0) - m.lastBytesSent) + currentStats.BytesSent + 1
	}

	// Calculate speeds (bytes per second)
	if timeDelta > 0 {
		downloadBps = uint64(float64(recvDelta) / timeDelta)
		uploadBps = uint64(float64(sentDelta) / timeDelta)
	}

	// Update last values
	m.lastBytesRecv = currentStats.BytesRecv
	m.lastBytesSent = currentStats.BytesSent
	m.lastTime = currentStats.Time

	return downloadBps, uploadBps, nil
}

// GetInterfaceName returns the name of the active network interface
func (m *Monitor) GetInterfaceName() string {
	return m.interfaceName
}

// StartMonitoring starts continuous monitoring with a callback
func (m *Monitor) StartMonitoring(ctx context.Context, interval time.Duration, callback func(downloadBps, uploadBps uint64, err error)) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			downloadBps, uploadBps, err := m.GetSpeeds()
			callback(downloadBps, uploadBps, err)
		}
	}
}

