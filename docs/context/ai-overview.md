## NetFlow — AI Overview

This document gives AI assistants a fast, high-level understanding of the NetFlow project so they can make safe, aligned changes.

### 1. What NetFlow is

- **Type**: Cross-platform system tray utility.
- **Purpose**: Show **real-time download and upload speeds** in the system tray / menu bar tooltip.
- **Surfaces**: Only three UI surfaces exist in v1 — tray icon, tooltip, and a minimal context menu with “Quit”.
- **Experience goal**: Feels like a built-in OS utility; invisible, calm, and professional.

Key requirements (from `README.md` and `PRD.md`):

- Tooltip updates every **1 second**.
- Accuracy within **±2%**.
- CPU: **< 1% idle**, **< 3% active**.
- Memory: **< 20 MB**.
- Startup to visible tray icon in **< 2 seconds**.
- No telemetry, no external APIs, no admin privileges.

### 2. Current architecture (mental model)

Core flow (see `README.md` and source files):

- `main.go`
  - Handles process lifecycle, signal handling, and a 1-second ticker.
  - Wires together network monitoring, formatting, and the tray UI.
- `network/monitor.go`
  - Uses `gopsutil` to read cumulative interface counters.
  - Sums non-loopback interfaces, handles counter rollover, and computes bytes-per-second deltas.
- `formatter/formatter.go`
  - Converts raw bytes-per-second values into human-readable strings like `↓ 12.4 MB/s    ↑ 1.8 MB/s` following strict rules in `UI_PLAN.md`.
- `tray/tray.go`
  - Owns the system tray integration (`systray`), app icon, context menu (“NetFlow” label + separator + “Quit”), and updating the tooltip.
- `internal/config/config.go`
  - Reserved for future configuration, using platform-appropriate config paths.
- `internal/icon/icon.go` and `public/netflow.ico`
  - Provide icon assets and helpers; the tray UI uses embedded icon resources.

**Separation of concerns (must keep):**

- Network monitor: only collects data.
- Formatter: only formats strings.
- Tray UI: only displays formatted strings and handles simple user interaction.

### 3. UX and UI constraints

From `PRD.md`, `UI_PLAN.md`, and global UI rules:

- **No additional windows, dialogs, notifications, or popups.**
- **No settings UI in v1**; if a feature needs configuration, it likely does not belong in this version.
- Tooltip format and spacing are **strict**:
  - Always `↓ ...    ↑ ...` on a **single line**.
  - One space between arrow and number; one between number and unit; **two spaces** between download and upload segments.
- Icon:
  - Static, monochrome, OS-themed.
  - No animations, color changes, or state-dependent variations.
- Overall UI feel:
  - Calm, predictable, and premium.
  - Inspired by high-quality OS utilities but **must not copy any proprietary system UI directly**.

### 4. Scope boundaries for AI changes

When proposing or implementing changes, AI assistants should:

- **Stay in scope for v1**:
  - Do not add graphs, historical statistics, complex diagnostics, or multi-interface selection unless explicitly requested and PRD/UI docs are updated.
- **Preserve the 3-surface model**:
  - No new UI surfaces without an explicit design update (e.g., new PRD/UI_PLAN section).
- **Keep cross-platform behavior consistent**:
  - Any platform-specific change should be justified and reflected in docs if it changes behavior.
- **Respect privacy**:
  - Do not introduce telemetry, logging of user traffic patterns, or any external network calls.

### 5. Where to look before coding

Before making any non-trivial change, AI agents should:

1. **Check `README.md`** for overview, architecture diagram, and build instructions.
2. **Read `docs/PRD.md`** for product goals, scope, and non-functional requirements.
3. **Read `docs/UI_PLAN.md`** for strict tooltip and tray behavior.
4. **Skim relevant Go files**:
   - `main.go`
   - `network/monitor.go`
   - `formatter/formatter.go`
   - `tray/tray.go`

Use this document as the starting point, then follow links to deeper sources as needed.

