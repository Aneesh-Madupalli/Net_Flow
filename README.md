# NetFlow

**Enterprise-grade, cross-platform system tray utility for real-time network speed monitoring.**

[![Version](https://img.shields.io/badge/version-1.0-blue.svg)](https://github.com/Aneesh-Madupalli/Net_Flow/releases)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://golang.org/)
[![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20macOS%20%7C%20Linux-lightgrey.svg)]()
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

---

## Table of Contents

- [Download — ready-to-run binaries](#download--ready-to-run-binaries)
- [Quick start: clone, build & run](#quick-start-clone-build--run)
- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Architecture](#architecture)
- [Project Structure](#project-structure)
- [Building from Source](#building-from-source)
- [Technology Stack](#technology-stack)
- [System Requirements](#system-requirements)
- [Performance](#performance)
- [Privacy & Security](#privacy--security)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)

---

## Download — ready-to-run binaries

Pre-built binaries: **click to download** from the latest [GitHub Release](https://github.com/Aneesh-Madupalli/Net_Flow/releases/latest).

| Platform | Download | How to run |
|----------|----------|------------|
| **Windows** | [⬇ netflow.exe](https://github.com/Aneesh-Madupalli/Net_Flow/releases/latest/download/netflow.exe) | Double-click or run from Explorer. Icon in system tray (no console window). |
| **macOS (Intel)** | [⬇ netflow_macos_amd64](https://github.com/Aneesh-Madupalli/Net_Flow/releases/latest/download/netflow_macos_amd64) | In Terminal: `chmod +x netflow_macos_amd64` then `./netflow_macos_amd64`. Icon in menu bar. |
| **macOS (Apple Silicon)** | [⬇ netflow_macos_arm64](https://github.com/Aneesh-Madupalli/Net_Flow/releases/latest/download/netflow_macos_arm64) | In Terminal: `chmod +x netflow_macos_arm64` then `./netflow_macos_arm64`. Icon in menu bar. |
| **Linux** | [⬇ netflow_linux_amd64](https://github.com/Aneesh-Madupalli/Net_Flow/releases/latest/download/netflow_linux_amd64) | In terminal: `chmod +x netflow_linux_amd64` then `./netflow_linux_amd64`. Icon in system tray. |

- **Windows:** No installer; run the `.exe`. Optional: add to Startup for run at logon.
- **macOS / Linux:** Grant network access if prompted. Optional: add to Login Items (macOS) or session autostart (Linux).

Hover over the tray/menu bar icon to see speeds; right-click → **Quit** to exit.

*Binaries are built automatically by [GitHub Actions](.github/workflows/build_release.yml) on push to the `build` branch.*

---

## Quick start: clone, build & run

**Prerequisite:** [Go 1.21 or later](https://go.dev/doc/install) installed and on your `PATH`.

### 1. Clone the repository

```bash
git clone https://github.com/Aneesh-Madupalli/Net_Flow.git
cd Net_Flow
```

### 2. Download dependencies

```bash
go mod download
```

### 3. Build for your platform

**Windows (no console window; tray only):**

```bash
go build -ldflags="-s -w -H windowsgui" -o netflow.exe
```

**macOS (Intel):**

```bash
go build -ldflags="-s -w" -o netflow
```

**macOS (Apple Silicon):**

```bash
GOARCH=arm64 go build -ldflags="-s -w" -o netflow
```

**Linux:**

```bash
go build -ldflags="-s -w" -o netflow
```

### 4. Run the application

- **Windows:** Double-click `netflow.exe` or run `.\netflow.exe` in PowerShell. The icon appears in the system tray (no console window).
- **macOS / Linux:** Run `./netflow`. The icon appears in the menu bar (macOS) or system tray (Linux).

Hover over the icon to see download and upload speeds. Right-click → **Quit** to exit.

See [Building from Source](#building-from-source) for more options.

---

## Overview

NetFlow is a lightweight, cross-platform application that displays **real-time download and upload speeds** in the system tray (Windows, Linux) or menu bar (macOS). It runs without windows or dialogs, uses minimal CPU and memory, and requires no configuration to start. NetFlow is designed to feel like a built-in system utility: calm, minimal, and always available.

**Target users:** Developers, remote workers, content creators, IT professionals, and power users who need instant network visibility without opening another app.

---

## Features

| Feature | Description |
|--------|-------------|
| **Real-time speeds** | Download and upload in B/s, KB/s, MB/s, or GB/s; updates every 1 second |
| **Tray icon** | App icon in tray; hover to see app name and download/upload speeds |
| **Cross-platform** | Windows 10+, macOS 10.15+, Linux (Ubuntu 20.04+, Debian 11+, Fedora 34+, or equivalent) |
| **Zero config** | Works out of the box |
| **Lightweight** | &lt; 1% CPU idle, &lt; 3% active; &lt; 20 MB RAM |
| **Privacy-first** | No telemetry, no external calls, no network transmission; all processing is local |
| **Graceful shutdown** | Handles SIGINT/SIGTERM and Quit from the tray menu |

### Tooltip and speed format

- **Format:** `↓ 12.4 MB/s    ↑ 1.8 MB/s` (download first, then upload; two spaces between)
- **Units:** 0 B/s → shown as `0 KB/s`; &lt; 1 MB/s → integer KB/s; 1–999 MB/s → one decimal; ≥ 1 GB/s → two decimals
- **Error state:** `↓ -- KB/s    ↑ -- KB/s` when no interface or error (no popups)

---

## Installation

- **Use a pre-built binary:** See [Download — ready-to-run binaries](#download--ready-to-run-binaries) for Windows, macOS, and Linux. Download the file for your platform and run it.
- **Build from source:** See [Quick start: clone, build & run](#quick-start-clone-build--run).

---

## Configuration

NetFlow may store an optional config file under the platform config directory for future use. No configuration is required; the tray always shows the app icon with speeds on hover.

| Platform | Config path |
|----------|-------------|
| **Windows** | `%AppData%\NetFlow\config.json` |
| **macOS** | `~/Library/Application Support/NetFlow/config.json` |
| **Linux** | `~/.config/NetFlow/config.json` |

---

## Usage

1. **Launch** — Run the NetFlow executable. The tray/menu bar icon appears within a few seconds.
2. **View speeds** — Hover over the icon to see the tooltip (app name + download/upload speeds).
3. **Quit** — Right-click the icon → **Quit**. No confirmation dialog.

---

## Architecture

```
┌─────────────────────────────────────────────────────────────────────────┐
│                              main.go                                     │
│  • Context & signal handling  • 1s ticker  • Wires monitor ↔ tray        │
└───────────────────────────────┬─────────────────────────────────────────┘
                                │
        ┌───────────────────────┼───────────────────────┐
        ▼                       ▼                         ▼
┌───────────────┐      ┌─────────────────┐      ┌─────────────────┐
│   network/    │      │   formatter/     │      │     tray/        │
│   monitor.go  │─────▶│   formatter.go  │─────▶│     tray.go      │
│               │      │                 │      │                  │
│ • gopsutil    │      │ • B/s → KB/MB/GB │      │ • systray        │
│ • Byte deltas │      │ • Tooltip string │      │ • Icon / title   │
│ • Rollover    │      │                 │      │ • Menu, Quit    │
└───────────────┘      └─────────────────┘      └────────┬─────────┘
                                                         │
                        ┌────────────────────────────────┼────────────────┐
                        ▼                                ▼                ▼
               ┌─────────────────┐             ┌─────────────────┐  ┌──────────┐
               │ internal/config │             │  internal/icon   │  │  public/  │
               │  config.go      │             │  icon.go         │  │ netflow.ico│
               │ (reserved)      │             │ • TransparentIcon│  │ (embedded)│
               └─────────────────┘             │ • SVG/PNG        │  └──────────┘
                                               └─────────────────┘
```

- **Network monitor:** Reads interface counters via `gopsutil`, sums non-loopback interfaces, computes bytes-per-second from deltas with rollover handling.
- **Formatter:** Converts bytes/s to human-readable strings for the tooltip.
- **Tray:** Renders menu (NetFlow, Quit), shows embedded app icon, updates tooltip with speeds on hover.
- **Config:** Optional; config directory reserved for future use.
- **Icon:** Transparent icon and SVG/PNG helpers for internal use; app icon is embedded from `public/`.

---

## Project Structure

```
netflow/
├── main.go                 # Entry point, ticker, shutdown, embed of icon
├── go.mod / go.sum         # Module and dependencies
├── docs/                   # Project documentation and AI context
│   ├── README.md           # Documentation index
│   ├── PRD.md              # Product requirements
│   ├── UI_PLAN.md          # UI/UX specification
│   └── context/            # AI-focused context and guidelines
├── formatter/
│   └── formatter.go        # Speed formatting, tooltip
├── network/
│   └── monitor.go          # Interface stats, deltas, GetSpeeds()
├── tray/
│   └── tray.go             # Systray UI, menu, UpdateTooltip
├── internal/
│   ├── config/
│   │   └── config.go       # Config path and Load/Save (reserved)
│   └── icon/
│       └── icon.go         # TransparentIcon, SVG/PNG helpers
├── public/
│   └── netflow.ico         # Application icon (embedded)
├── .github/workflows/
│   └── build_release.yml   # CI: build and publish binaries to GitHub Releases
├── Makefile                # build, deps, clean, test, fmt, lint
├── LICENSE
└── README.md               # This file
```

---

## Building from Source

### Prerequisites

- **Go 1.21 or later** — [Install Go](https://go.dev/doc/install)
- **Git** — To clone the repository

For the shortest path to a running app, follow [Quick start: clone, build & run](#quick-start-clone-build--run) above.

### Quick build (current platform)

From the project root after cloning and `go mod download`:

```bash
# Linux or macOS
go build -ldflags="-s -w" -o netflow
./netflow

# Windows (GUI app, no console)
go build -ldflags="-s -w -H windowsgui" -o netflow.exe
.\netflow.exe
```

### Make targets

| Target | Description |
|--------|-------------|
| `make build` | Build for current OS/arch → `netflow` (or `netflow.exe` on Windows) |
| `make deps` | `go mod download` and `go mod tidy` |
| `make clean` | Remove build artifacts |
| `make test` | Run tests |
| `make fmt` | Format code |
| `make lint` | Run linter (if configured) |

### Cross-compile (manual)

```bash
# Windows (GUI app: no console window)
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H windowsgui" -o netflow.exe

# macOS (Intel / Apple Silicon)
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o netflow_macos_amd64
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o netflow_macos_arm64

# Linux
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o netflow_linux_amd64
```

The `-ldflags="-s -w"` strips debug info; on Windows `-H windowsgui` builds a GUI app with no console window.

---

## Technology Stack

| Layer | Technology |
|-------|------------|
| **Language** | Go 1.21+ |
| **Module** | `netflow` |
| **Network stats** | [gopsutil/v3](https://github.com/shirou/gopsutil) — cross-platform interface counters (native OS APIs) |
| **System tray** | [getlantern/systray](https://github.com/getlantern/systray) — icon, tooltip, menu (Windows/macOS/Linux) |
| **Icon** | Embedded `public/netflow.ico`; `internal/icon` for TransparentIcon and SVG/PNG helpers |
| **Config** | Optional; `internal/config` for path and Load/Save (reserved for future use) |

All dependencies are vendored or fetched via `go mod`; no external services or APIs are called at runtime.

---

## System Requirements

| Platform | Version / notes |
|----------|------------------|
| **Windows** | Windows 10 or later; amd64. No admin rights required. Runs as GUI app (no console window). |
| **macOS** | macOS 10.15 (Catalina) or later; Intel (amd64) or Apple Silicon (arm64). Menu bar only; no Dock icon. |
| **Linux** | Ubuntu 20.04+, Debian 11+, Fedora 34+, or equivalent; amd64. Requires a desktop environment with system tray support. |

---

## Performance

| Metric | Target |
|--------|--------|
| **CPU (idle)** | &lt; 1% |
| **CPU (active)** | &lt; 3% |
| **Memory** | &lt; 20 MB |
| **Update interval** | 1 second |
| **Startup** | &lt; 2 seconds to tray visibility |
| **Accuracy** | Within ±2% of actual line rate in typical use |

---

## Privacy & Security

- **No telemetry or analytics** — The app does not send data to any server.
- **No network usage** — Beyond reading interface counters, it does not open connections.
- **No elevated privileges** — Runs as the current user; no admin/root.
- **Local only** — Config and all processing are on the host; optional config file is user-writable only.

---

## Troubleshooting

| Issue | What to try |
|-------|-------------|
| **Icon doesn’t appear** | Windows: Check notification area visibility. macOS: Check menu bar. Linux: Ensure tray (e.g. libappindicator) is available. |
| **Speeds always 0 or --** | Confirm an active interface (Wi‑Fi or Ethernet). macOS: Grant network/firewall access if prompted. Restart NetFlow. |
| **App won’t start** | Ensure Go 1.21+ when building from source. Verify OS and architecture match the binary. On Windows, use `-H windowsgui` to avoid console. |
| **Windows: console window appears** | Rebuild with `-ldflags="-s -w -H windowsgui"` so the exe uses the GUI subsystem. |
---

## Contributing

Contributions are welcome. Please open an issue or pull request; ensure code is formatted (`go fmt`) and that tests pass (`go test ./...`).

---

## License

This project is licensed under the **MIT License**. See [LICENSE](LICENSE) for the full text.

---

## Acknowledgments

- [Go](https://golang.org/) — Language and toolchain
- [gopsutil](https://github.com/shirou/gopsutil) — Cross-platform process and network utilities
- [systray](https://github.com/getlantern/systray) — Cross-platform system tray library

---

**NetFlow** — *A signal you trust*
