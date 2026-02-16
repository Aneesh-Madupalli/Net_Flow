package main

import (
	"context"
	_ "embed"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"netflow/network"
	"netflow/tray"
// )

// //go:embed public/netflow.ico
// var iconICO []byte

// const (
// 	updateIntervalSeconds = 1
// )

// func main() {
// 	// Create context for graceful shutdown
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	// Handle OS signals for graceful shutdown
// 	sigChan := make(chan os.Signal, 1)
// 	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

// 	// Initialize network monitor
// 	monitor, err := network.NewMonitor()
// 	if err != nil {
// 		log.Fatalf("Failed to initialize network monitor: %v", err)
// 	}

// 	// Create tray instance with embedded icon
// 	trayInstance := tray.NewTray(iconICO)

// 	// Set up tray callbacks
// 	trayInstance.SetOnReady(func() {
// 		log.Println("NetFlow started successfully")
		
// 		// Initial update
// 		downloadBps, uploadBps, err := monitor.GetSpeeds()
// 		if err != nil {
// 			log.Printf("Warning: Error getting initial speeds: %v", err)
// 			trayInstance.UpdateTooltip(0, 0, err)
// 		} else {
// 			trayInstance.UpdateTooltip(downloadBps, uploadBps, nil)
// 		}
		
// 		// Start monitoring in background
// 		go func() {
// 			ticker := time.NewTicker(time.Duration(updateIntervalSeconds) * time.Second)
// 			defer ticker.Stop()

// 			// Wait a bit before first real update to get accurate delta
// 			time.Sleep(time.Duration(updateIntervalSeconds) * time.Second)

// 			for {
// 				select {
// 				case <-ctx.Done():
// 					return
// 				case <-ticker.C:
// 					downloadBps, uploadBps, err := monitor.GetSpeeds()
// 					if err != nil {
// 						// Log error but don't spam
// 						select {
// 						case <-time.After(5 * time.Second):
// 							log.Printf("Error getting speeds: %v", err)
// 						default:
// 						}
// 						trayInstance.UpdateTooltip(0, 0, err)
// 					} else {
// 						trayInstance.UpdateTooltip(downloadBps, uploadBps, nil)
// 					}
// 				}
			}
		}()
	})

	trayInstance.SetOnExit(func() {
		log.Println("NetFlow shutting down...")
		cancel()
	})

	// Handle OS signals
	go func() {
		<-sigChan
		log.Println("Received shutdown signal")
		cancel()
		trayInstance.UpdateTooltip(0, 0, nil)
		// Give a moment for cleanup
		time.Sleep(100 * time.Millisecond)
		os.Exit(0)
	}()

	// Run tray (this blocks until quit)
	trayInstance.Run()
}

