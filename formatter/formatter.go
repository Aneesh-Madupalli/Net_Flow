package formatter

import "fmt"

// FormatSpeed formats bytes per second into human-readable format
// Format rules:
// - 0 B/s -> 0 KB/s
// - 1-999 B/s -> ### B/s
// - 1-999 KB/s -> ### KB/s
// - 1.0-999.9 MB/s -> #.0 MB/s
// - ≥ 1 GB/s -> #.00 GB/s
func FormatSpeed(bps uint64) string {
	const (
		KB = 1024
		MB = 1024 * 1024
		GB = 1024 * 1024 * 1024
	)

	if bps == 0 {
		return "0 KB/s"
	}

	if bps < KB {
		return fmt.Sprintf("%d B/s", bps)
	}

	if bps < MB {
		return fmt.Sprintf("%d KB/s", bps/KB)
	}

	if bps < GB {
		mbs := float64(bps) / float64(MB)
		return fmt.Sprintf("%.1f MB/s", mbs)
	}

	gbs := float64(bps) / float64(GB)
	return fmt.Sprintf("%.2f GB/s", gbs)
}

// FormatTooltip formats download and upload speeds for tooltip display
// Format: "↓ 12.4 MB/s    ↑ 1.8 MB/s"
func FormatTooltip(downloadBps, uploadBps uint64) string {
	downloadStr := FormatSpeed(downloadBps)
	uploadStr := FormatSpeed(uploadBps)
	return fmt.Sprintf("↓ %s    ↑ %s", downloadStr, uploadStr)
}

// FormatTooltipError formats error state for tooltip
func FormatTooltipError() string {
	return "↓ -- KB/s    ↑ -- KB/s"
}

