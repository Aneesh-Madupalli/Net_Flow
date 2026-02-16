# 🎨 UI PLAN — Network Speed Indicator (v1.0)
## Final Design Specification

**Design Philosophy**: *"Invisible excellence — feels like part of the OS"*

---

## 📐 Design Principles

### Core Values
1. **Silent Confidence** — Never draws attention to itself
2. **Precision Over Polish** — Accurate data, not flashy animations
3. **OS-Native Feel** — User forgets it's a third-party app
4. **Zero Cognitive Load** — Instant comprehension, zero learning curve

### User Emotion Goal
> *"This feels like it was always there, and I can't imagine my system without it."*

---

## 🎯 UI SURFACES (Only 3 Exist)

| Surface | Exists? | Purpose | Interaction |
|---------|---------|---------|-------------|
| System Tray Icon | ✅ | Visual presence | Passive (no click needed) |
| Tooltip | ✅ | Primary information | Hover/click to reveal |
| Context Menu | ✅ | Exit only | Right-click |

**Rule**: No app window, no dialogs, no popups, no notifications.

---

## 1️⃣ SYSTEM TRAY ICON — Final Specification

### Visual Identity

**Icon Concept**: Dual Flow Arrow (Network Flow Symbol)

```
    ↑
  ────
    ↓
```

**Design Rules**:
- **Style**: Minimalist outline, no fill
- **Color**: OS-controlled (template image)
- **Size**: 16x16px (standard), scales to 22x22px (HiDPI)
- **Weight**: Thin stroke (1-1.5px)
- **Balance**: Vertically centered, horizontally balanced

### Platform-Specific Behavior

| Platform | Icon Behavior |
|----------|---------------|
| **macOS** | Template image (auto light/dark mode), menu bar native |
| **Windows** | Monochrome tray icon, respects system theme |
| **Linux** | Desktop environment handles coloring automatically |

### Icon States

| State | Visual | Notes |
|-------|--------|-------|
| **Normal** | Standard icon | Always visible |
| **Active** | Same icon | No animation, no color change |
| **Error** | Same icon | Tooltip shows error message |

**Critical Rule**: Icon never changes appearance. Consistency = trust.

### Icon Design Specifications

**SVG Structure**:
- ViewBox: `0 0 16 16`
- Stroke: `currentColor` (OS-controlled)
- Fill: `none`
- Stroke width: `1.5`
- Rounded line caps

**Visual Balance**:
- Up arrow: 4px height, centered
- Horizontal line: 6px width, centered
- Down arrow: 4px height, centered
- Total vertical spacing: 2px between elements

---

## 2️⃣ TOOLTIP — Primary Information Display

### Layout (Immutable)

```
↓ 12.4 MB/s    ↑ 1.8 MB/s
```

### Why This Layout Works

1. **Download First** — Matches user mental model (download > upload priority)
2. **Visual Hierarchy** — Arrows eliminate need for labels
3. **Single Line** — No eye travel, instant scan
4. **Balanced Spacing** — Two spaces between sections creates visual separation
5. **No Noise** — No separators, no punctuation, no units label

### Spacing Rules (Exact)

```
↓ [space] [number] [space] [unit] [space][space] ↑ [space] [number] [space] [unit]
```

- **One space** between arrow and number
- **Two spaces** between download and upload sections
- **One space** between number and unit

### Typography

- **Font**: System default (OS-controlled)
- **Size**: System default tooltip size
- **Weight**: Regular (not bold)
- **Color**: System default tooltip text color

### Update Behavior

**Timing**:
- Fixed 1-second interval
- No adaptive timing
- No smoothing

**Visual Update**:
- Numbers replace silently
- No fade, no tween, no blink
- No visual transition effects

> *"The best update is invisible — user notices the number changed, not the change itself."*

---

## 3️⃣ DATA FORMATTING — Precision Rules

### Unit Conversion Table

| Range | Format | Example | Rationale |
|-------|--------|---------|-----------|
| 0 B/s | `0 KB/s` | `0 KB/s` | Avoids "0 B/s" confusion |
| 1-999 B/s | `### B/s` | `125 B/s` | Full precision for low speeds |
| 1-999 KB/s | `### KB/s` | `512 KB/s` | No decimals needed |
| 1.0-999.9 MB/s | `#.0 MB/s` | `12.4 MB/s` | One decimal for readability |
| ≥ 1 GB/s | `#.00 GB/s` | `1.25 GB/s` | Two decimals for precision |

### Formatting Logic

```go
// Pseudo-code logic
if speed < 1024 {
    return fmt.Sprintf("%d B/s", speed)
} else if speed < 1024*1024 {
    return fmt.Sprintf("%d KB/s", speed/1024)
} else if speed < 1024*1024*1024 {
    return fmt.Sprintf("%.1f MB/s", float64(speed)/(1024*1024))
} else {
    return fmt.Sprintf("%.2f GB/s", float64(speed)/(1024*1024*1024))
}
```

### Precision Philosophy

- **Low speeds** (< 1 MB/s): Integer precision (no decimals)
- **Medium speeds** (1-999 MB/s): One decimal (balance of precision and readability)
- **High speeds** (≥ 1 GB/s): Two decimals (precision matters at scale)

**Rationale**: 
- Too much precision = visual noise
- Too little precision = loss of information
- Sweet spot = professional, readable, accurate

---

## 4️⃣ IDLE & LOW-TRAFFIC STATES

### Idle Definition

Speed < 1 KB/s for both download and upload

### Display Format

```
↓ 0 KB/s    ↑ 0 KB/s
```

**Rules**:
- Always show `0 KB/s` (never `0 B/s`)
- No special "Idle" text
- No grayed-out appearance
- No hidden tooltip
- Same format as active state

**Philosophy**: Silence = confidence. Zero is information, not absence.

### Edge Cases

| Scenario | Display | Notes |
|----------|---------|-------|
| No network interface | `↓ -- KB/s    ↑ -- KB/s` | Clear error state |
| Interface disconnected | `↓ 0 KB/s    ↑ 0 KB/s` | Treat as idle |
| Calculation error | `↓ -- KB/s    ↑ -- KB/s` | Error indicator |

---

## 5️⃣ CONTEXT MENU — Minimal v1.0

### Menu Structure

```
NetFlow        (disabled label; app name)
──────────────────────
Quit
```

As implemented: first item is "NetFlow" (disabled), then separator, then "Quit".

### Menu Rules

- **No icons** in menu items
- **No keyboard shortcuts** displayed
- **No settings** (if user needs settings, v1 failed)
- **No about dialog** (version in tooltip on hover, optional)
- **No toggles** or checkboxes

### Platform-Specific Menu Behavior

| Platform | Behavior |
|----------|----------|
| **macOS** | Native menu bar menu, separator line |
| **Windows** | Right-click context menu |
| **Linux** | Desktop environment context menu |

### Menu Item Actions

| Item | Action | Notes |
|------|--------|-------|
| **Quit** | Exit application | Graceful shutdown, no confirmation |

**Future Extension** (v2+ only):
- If users request settings, add "Preferences..." item
- Opens minimal settings window (not in v1)

---

## 6️⃣ ERROR HANDLING (UI Perspective)

### Error States

| Error Type | Tooltip Display | User Action |
|------------|-----------------|-------------|
| No network interface | `↓ -- KB/s    ↑ -- KB/s` | None (auto-recover) |
| Permission denied | `↓ -- KB/s    ↑ -- KB/s` | Check permissions |
| Interface disconnected | `↓ 0 KB/s    ↑ 0 KB/s` | None (auto-recover) |

### Error Philosophy

- **Silent recovery** — Auto-recover when possible
- **Clear indicators** — `--` shows error state
- **No popups** — Errors in tooltip only
- **No notifications** — Never interrupt user

---

## 7️⃣ ACCESSIBILITY & COMPLIANCE

### Accessibility Requirements

✅ **Screen Reader Support**
- Tooltip text is readable by screen readers
- Icon has accessible name: "NetFlow" (app name in tooltip)

✅ **Visual Accessibility**
- No color-only meaning (arrows + text)
- High contrast (OS-controlled)
- No flashing or animations

✅ **Keyboard Navigation**
- Context menu accessible via keyboard (platform-standard)

### System Compliance

✅ **macOS**
- Follows Human Interface Guidelines
- Menu bar icon (not Dock)
- No app window
- No focus stealing

✅ **Windows**
- Follows Fluent Design principles
- Tray icon only (not taskbar)
- No balloon notifications
- Respects system theme

✅ **Linux**
- Respects desktop environment conventions
- Uses standard system tray APIs
- No custom UI layers

---

## 8️⃣ PLATFORM-SPECIFIC UX GUARANTEES

### macOS Specific

- **Menu Bar Integration**: Feels native to macOS menu bar
- **Dark Mode**: Automatic template image support
- **No Dock Icon**: Runs as menu bar app only
- **No Window**: Never creates app window
- **Focus**: Never steals focus or shows dialogs

### Windows Specific

- **System Tray**: Native Windows notification area
- **Theme Support**: Respects light/dark theme
- **No Taskbar**: Never appears in taskbar
- **No Notifications**: No toast notifications or balloons

### Linux Specific

- **Desktop Environment**: Works with GNOME, KDE, XFCE, Cinnamon
- **System Tray**: Uses standard tray APIs (libappindicator)
- **Theme**: Respects GTK/Qt themes
- **No Custom UI**: No custom window managers or UI layers

---

## 9️⃣ UI FAILURE CONDITIONS (DO NOT SHIP IF)

### Critical Failures

❌ **Numbers jump visually** — Formatting inconsistency  
❌ **Tooltip flickers** — Update timing issues  
❌ **Icon draws attention** — Animation or color changes  
❌ **User notices CPU usage** — Performance issues  
❌ **User wants to "configure"** — Missing essential defaults

### Quality Gates

Before release, verify:
- [ ] Tooltip updates smoothly (no flicker)
- [ ] Numbers format consistently
- [ ] Icon never changes appearance
- [ ] CPU usage < 1% (idle), < 3% (active)
- [ ] Memory usage < 20MB
- [ ] Works in light and dark mode
- [ ] Screen reader compatible
- [ ] No visual glitches on any platform

---

## 🔟 UI → CODE ARCHITECTURE

### Separation of Concerns

```
┌─────────────────────┐
│  Network Monitor    │ → Emits: (download, upload) bytes
└─────────────────────┘
         ↓
┌─────────────────────┐
│  Speed Calculator   │ → Calculates: speeds in bytes/sec
└─────────────────────┘
         ↓
┌─────────────────────┐
│  Formatter          │ → Formats: "↓ 12.4 MB/s    ↑ 1.8 MB/s"
└─────────────────────┘
         ↓
┌─────────────────────┐
│  Tray UI            │ → Displays: tooltip string only
└─────────────────────┘
```

### UI Layer Rules

**UI NEVER**:
- Calculates speeds
- Decides units
- Applies business logic
- Handles network monitoring

**UI ONLY**:
- Displays formatted string
- Updates tooltip
- Shows context menu
- Handles user interactions

---

## 1️⃣1️⃣ FUTURE-SAFE EXTENSION (Without Breaking v1)

### Extension Points (v2+ Only)

If users request more features, extend without breaking v1:

| Feature | Location | Implementation |
|---------|----------|----------------|
| **Peak Speed** | Tooltip line 2 | `Peak: ↓ 25.3 MB/s` |
| **Interface Name** | Tooltip footer | `Wi-Fi (en0)` |
| **Total Data** | Tooltip line 3 | `Today: ↓ 2.5 GB ↑ 500 MB` |
| **Preferences** | Context menu | "Preferences..." item |
| **Statistics** | Context menu | "Show Statistics..." item |

### Extension Rules

1. **v1 remains untouched** — Core tooltip format never changes
2. **Additive only** — New info in new lines/sections
3. **Opt-in** — New features hidden by default
4. **Backward compatible** — v1 users see no change

---

## 1️⃣2️⃣ REFINEMENTS & ENHANCEMENTS

### Subtle Enhancements (Keeping Minimal)

#### 1. **Smart Zero Display**
- Show `0 KB/s` instead of `0 B/s` for consistency
- Makes idle state feel intentional, not broken

#### 2. **Consistent Spacing**
- Exact spacing rules prevent visual jitter
- Two spaces between sections creates clear separation

#### 3. **Error State Clarity**
- `--` clearly indicates error (not zero)
- User understands something is wrong, not just idle

#### 4. **Platform-Native Feel**
- OS-controlled colors and fonts
- Feels like built-in system utility

#### 5. **Silent Updates**
- No animations = no distraction
- Numbers just "are" — feels natural

### What We're NOT Adding (v1)

❌ Icon animations  
❌ Color changes  
❌ Sound effects  
❌ Notifications  
❌ Settings window  
❌ Statistics dashboard  
❌ History graphs  
❌ Multiple interfaces display

**Rationale**: Each addition breaks the "invisible excellence" principle.

---

## 1️⃣3️⃣ USER EXPERIENCE FLOW

### First Launch

1. User double-clicks executable (on Windows: no console window; GUI app)
2. Icon appears in system tray (2 seconds max)
3. Tooltip shows app name + speeds on hover
4. No setup, no configuration, no questions

### Daily Use

1. User hovers over icon (or clicks, platform-dependent)
2. Tooltip shows current speeds
3. User gets information instantly
4. User continues work (no interruption)

### Exit

1. User right-clicks icon
2. Clicks "Quit"
3. Application exits gracefully
4. No confirmation, no delay

### Error Recovery

1. Network disconnects → Tooltip shows `0 KB/s`
2. Network reconnects → Tooltip updates automatically
3. No user action required
4. Silent, automatic recovery

---

## 1️⃣4️⃣ VISUAL DESIGN SPECIFICATIONS

### Icon Design (Detailed)

**SVG Code Structure**:

```svg
<svg viewBox="0 0 16 16" xmlns="http://www.w3.org/2000/svg">
  <path d="M8 2 L8 6 M8 10 L8 14 M4 8 L12 8" 
        stroke="currentColor" 
        stroke-width="1.5" 
        stroke-linecap="round" 
        fill="none"/>
</svg>
```

**Measurements**:
- Up arrow: From y=2 to y=6 (4px height)
- Horizontal line: From x=4 to x=12 (8px width)
- Down arrow: From y=10 to y=14 (4px height)
- Center point: (8, 8)
- Total icon: 16x16px

### Tooltip Design

**Layout**:

```
↓ 12.4 MB/s    ↑ 1.8 MB/s
```

**Measurements**:
- Arrow: Unicode `↓` and `↑`
- Spacing: 1 space (arrow-number), 2 spaces (sections)
- Alignment: Left-aligned text
- Padding: System default tooltip padding

---

## 1️⃣5️⃣ QUALITY ASSURANCE CHECKLIST

### Visual QA

- [ ] Icon renders correctly at 16x16px
- [ ] Icon renders correctly at 32x32px (HiDPI)
- [ ] Icon looks balanced and centered
- [ ] Tooltip text is readable
- [ ] Spacing is consistent
- [ ] Numbers format correctly
- [ ] Units convert properly
- [ ] Zero state displays correctly
- [ ] Error state displays correctly

### Functional QA

- [ ] Tooltip updates every 1 second
- [ ] No flickering during updates
- [ ] Context menu appears on right-click
- [ ] Quit works correctly
- [ ] Application exits gracefully
- [ ] No memory leaks
- [ ] CPU usage within limits

### Platform QA

- [ ] Works on Windows 10+
- [ ] Works on macOS 10.15+
- [ ] Works on Ubuntu 20.04+
- [ ] Works in light mode
- [ ] Works in dark mode
- [ ] Screen reader compatible
- [ ] No visual glitches

---

## 🏁 FINAL DESIGN SIGN-OFF

### Design Principles Achieved

✅ **Minimal** — Only essential UI surfaces  
✅ **Professional** — Feels like system utility  
✅ **OS-Native** — Respects platform conventions  
✅ **Timeless** — Won't feel dated in 5 years  
✅ **Maintenance-Free** — No UI debt for future

### Success Criteria

> *"If your app disappears into the OS and quietly earns trust — you have built a world-class utility."*

### Next Steps

1. ✅ **UI Plan Complete** — This document  
2. ⏭️ **Icon Design** — Create SVG icon files  
3. ⏭️ **Formatter Implementation** — Code the formatting logic  
4. ⏭️ **Tray Integration** — Implement system tray UI  
5. ⏭️ **QA Testing** — Visual and functional testing

---

**Document Status**: Final  
**Version**: 1.0  
**Last Updated**: 2025  
**Designer**: UI Architecture Team

---

## 📝 APPENDIX: Design Rationale

### Why No Settings in v1?

Settings imply the defaults are wrong. If we need settings, we failed at choosing good defaults. v1 should work perfectly with zero configuration.

### Why No Animations?

Animations draw attention. This app should be invisible. The best animation is no animation.

### Why Single-Line Tooltip?

Multi-line tooltips require eye travel. Single-line is instant comprehension. If you need more info, you're using the wrong tool.

### Why OS-Controlled Colors?

Native feel. Users trust system utilities. Third-party apps that look custom feel untrustworthy.

### Why Fixed 1-Second Updates?

Adaptive timing adds complexity. Fixed timing is predictable. Users learn the rhythm. 1 second is the sweet spot between accuracy and performance.

---

**End of UI Plan**

