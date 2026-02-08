package tray

import (
	"sync"

	"github.com/getlantern/systray"
	"netflow/formatter"
)

// Tray manages the system tray icon and tooltip
type Tray struct {
	iconData     []byte
	menuItemQuit *systray.MenuItem
	onReady      func()
	onExit       func()

	lastDown uint64
	lastUp   uint64
	lastErr  error
	mu       sync.RWMutex
}

// NewTray creates a new tray instance. iconData should be PNG or ICO bytes for the tray icon.
func NewTray(iconData []byte) *Tray {
	return &Tray{iconData: iconData}
}

// SetOnReady sets the callback for when tray is ready
func (t *Tray) SetOnReady(callback func()) {
	t.onReady = callback
}

// SetOnExit sets the callback for when user quits
func (t *Tray) SetOnExit(callback func()) {
	t.onExit = callback
}

// Run starts the system tray application
func (t *Tray) Run() {
	systray.Run(t.onTrayReady, t.onTrayExit)
}

// onTrayReady is called when the tray is ready
func (t *Tray) onTrayReady() {
	t.applyDisplay(0, 0, nil)

	systray.AddMenuItem("NetFlow", "").Disable()
	systray.AddSeparator()
	t.menuItemQuit = systray.AddMenuItem("Quit", "Exit NetFlow")

	go t.handleMenuClicks()

	if t.onReady != nil {
		t.onReady()
	}
}

// onTrayExit is called when the tray exits
func (t *Tray) onTrayExit() {
	if t.onExit != nil {
		t.onExit()
	}
}

// handleMenuClicks handles menu item clicks
func (t *Tray) handleMenuClicks() {
	for range t.menuItemQuit.ClickedCh {
		systray.Quit()
		return
	}
}

// applyDisplay sets icon and tooltip from current speed data (icon only: app icon, speeds on hover)
func (t *Tray) applyDisplay(downloadBps, uploadBps uint64, err error) {
	var speedStr string
	if err != nil {
		speedStr = formatter.FormatTooltipError()
	} else {
		speedStr = formatter.FormatTooltip(downloadBps, uploadBps)
	}
	if len(t.iconData) > 0 {
		systray.SetIcon(t.iconData)
	}
	systray.SetTitle("")
	systray.SetTooltip("NetFlow\n" + speedStr)
}

// UpdateTooltip updates the tooltip with new speed information
func (t *Tray) UpdateTooltip(downloadBps, uploadBps uint64, err error) {
	t.mu.Lock()
	t.lastDown, t.lastUp, t.lastErr = downloadBps, uploadBps, err
	t.mu.Unlock()
	t.applyDisplay(downloadBps, uploadBps, err)
}
