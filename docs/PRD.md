# Product Requirements Document (PRD)
## NetFlow

**Version:** 1.0  
**Date:** 2025  
**Status:** Final - Implemented  
**Category:** System Utility / Network Monitoring  
**Target Platforms:** Windows, macOS, Linux  
**Primary Language:** Go (Golang)

---

## 1. Executive Summary

### 1.1 Product Overview

**NetFlow** is a lightweight, cross-platform system tray utility that displays real-time network **download and upload speeds** directly in the system tray tooltip. It runs silently in the background, providing instant network visibility without disrupting the user's workflow.

NetFlow is designed to feel **native to the operating system**, requiring no configuration, no UI windows, and minimal system resources.

### 1.2 Problem Statement

Users often need to quickly verify whether their network is:

* Active
* Slow
* Uploading or downloading heavily

Existing solutions are:

* Heavy applications
* Browser-dependent
* Platform-specific
* Distracting or over-engineered

There is a need for a **simple, reliable, always-available network speed indicator**.

### 1.3 Solution

NetFlow provides:

* Real-time upload & download speed
* System tray integration
* Minimal resource usage
* Cross-platform compatibility
* Zero configuration

---

## 2. Objectives & Goals

### 2.1 Primary Objectives

1. Show accurate real-time download & upload speeds
2. Remain invisible unless the user checks it
3. Use minimal CPU and memory
4. Work consistently across operating systems
5. Require no user configuration

### 2.2 Success Criteria

* Tooltip updates every 1 second
* Speed accuracy within ±2%
* CPU usage < 1% idle, < 3% active
* Memory usage < 20 MB
* Startup time < 2 seconds
* Zero crashes during normal usage

---

## 3. Target Users

### 3.1 Primary Users

* Developers
* Remote workers
* Content creators
* Gamers
* IT professionals
* Power users

### 3.2 User Needs

* Quick glance network status
* No distractions
* Trustworthy data
* Always available
* No setup required

---

## 4. Scope

### 4.1 In Scope (v1)

* System tray icon
* Tooltip displaying download & upload speed
* Automatic network interface detection
* Cross-platform support
* Exit option

### 4.2 Out of Scope (v1)

* Graphs or charts
* Historical data
* Network diagnostics
* Alerts or notifications
* Settings UI
* Multiple interface display
* Mobile support

---

## 5. Functional Requirements

### 5.1 System Tray Integration

* NetFlow runs without opening any window
* Displays an icon in the system tray / menu bar
* No Dock or taskbar window

### 5.2 Tooltip Display

* Tooltip shows real-time speeds in one line:

```
↓ 12.4 MB/s    ↑ 1.8 MB/s
```

* Tooltip updates every 1 second
* Uses OS default tooltip styling

### 5.3 Speed Measurement

* Uses cumulative byte counters
* Calculates delta over time
* Handles counter rollover
* Filters loopback traffic
* Automatically selects active interface

### 5.4 Context Menu

Minimal menu (as implemented):

```
NetFlow     (disabled label)
────────
Quit
```

No Settings or display-mode options; tray shows app icon only, speeds on hover.

---

## 6. Non-Functional Requirements

### 6.1 Performance

* Idle CPU: < 1%
* Active CPU: < 3%
* RAM usage: < 20 MB

### 6.2 Reliability

* Handles network disconnects gracefully
* Recovers from interface changes
* No crashes on suspend / resume

### 6.3 Usability

* Zero configuration
* No learning curve
* OS-native behavior

### 6.4 Security & Privacy

* No data collection
* No external API calls
* No network transmission
* No admin privileges required

---

## 7. UI / UX Requirements

### 7.1 Design Principles

* Invisible
* Calm
* Minimal
* Professional
* OS-native

### 7.2 UI Elements

| Element   | Description                      |
| --------- | -------------------------------- |
| Tray Icon | Monochrome, static, no animation |
| Tooltip   | Single line, real-time speed     |
| Menu      | Exit only                        |

### 7.3 Idle State

When speed < 1 KB/s:

```
↓ 0 KB/s    ↑ 0 KB/s
```

---

## 8. Technical Requirements

### 8.1 Technology Stack

* Language: Go 1.21+
* Network stats: `github.com/shirou/gopsutil/v3/net`
* Tray UI: `github.com/getlantern/systray`

### 8.2 Architecture Overview

```
Network Monitor → Speed Calculator → Formatter → Tray Tooltip
```

UI layer only displays formatted output.

## 9. Platform Requirements

### 9.1 Windows

* Windows 10+
* x64 (amd64)
* Tray-only application; built with GUI subsystem (no console window)

### 9.2 macOS

* macOS 10.15+
* Intel & Apple Silicon
* Menu bar only
* No Dock icon

### 9.3 Linux

* Ubuntu 20.04+
* Debian 11+
* Fedora 34+
* GNOME, KDE, XFCE

---

## 10. Build & Distribution

### 10.1 Build Targets

* Windows: `netflow.exe` from GitHub Releases (GUI subsystem, no console)
* macOS: `netflow-macos-amd64`, `netflow-macos-arm64`
* Linux: `netflow-linux-amd64`

### 10.2 Distribution

* GitHub Releases
* Optional installers later

---

## 11. User Stories

### Core User Story

> As a user, I want to see my real-time upload and download speed at a glance so I can understand my network activity without opening any application.

---

## 12. Development Phases

### Phase 1 – Core MVP

* Network monitoring
* Speed calculation
* Tray icon
* Tooltip updates

### Phase 2 – Cross-Platform Testing

* Windows testing
* macOS testing
* Linux testing

### Phase 3 – Optimization

* Performance tuning
* Edge case handling
* Documentation

---

## 13. Risks & Mitigation

| Risk                       | Mitigation           |
| -------------------------- | -------------------- |
| Interface detection issues | Use gopsutil         |
| Tray differences           | Use systray          |
| Counter rollover           | Implement safe logic |

---

## 14. Future Enhancements (Post v1)

* Peak speed display
* Interface selection
* Optional preferences
* Statistics view (opt-in)
* Dark/light theming (OS-driven)

---

## 15. Success Metrics

* Stable tooltip updates
* No UI distractions
* Positive user feedback
* Minimal resource footprint
* Long runtime without restart

---

## 16. Final Approval

**Product Name:** NetFlow  
**Version:** 1.0  
**Status:** Approved for implementation

---

### 🏁 Final Note (Architect's Closing)

> NetFlow is not an app users interact with —  
> it is a **signal they trust**.

---

**Document Status**: Final  
**Last Updated**: 2025  
**Next Review**: Post v1.0 release

