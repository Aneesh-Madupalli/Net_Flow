## NetFlow — AI Coding Guidelines

This document gives **practical guardrails** for AI assistants modifying or extending the NetFlow codebase. It complements the detailed requirements in `docs/PRD.md` and `docs/UI_PLAN.md`.

Always prefer:

- **Small, well-scoped changes**
- **Preserving existing behavior and UX**
- **Consistency across platforms**

---

### 1. Architectural boundaries (do not break)

Keep the following separation of concerns intact:

- `network/monitor.go`
  - Only responsible for reading OS network counters and computing byte deltas.
  - May expose functions that return raw speeds in bytes per second.
  - Must not know about formatting, units, or UI.

- `formatter/formatter.go`
  - Converts bytes-per-second into display strings that follow the rules in `docs/UI_PLAN.md`.
  - Responsible for unit selection (B/s, KB/s, MB/s, GB/s) and correct spacing.
  - Must not call tray APIs or perform I/O.

- `tray/tray.go`
  - Owns systray/menu bar integration, app icon, and tooltip.
  - Accepts *already formatted strings* from the formatter.
  - Must not know about network counters, platform config paths, or business logic.

- `main.go`
  - Application entrypoint, lifecycle, and wiring between the above components.
  - Owns the 1-second ticker and orchestrates updates.

When adding new features, prefer:

- Adding **new functions** inside the appropriate package instead of mixing concerns.
- Keeping function signatures narrow and explicit.

---

### 2. UX and scope rules

Respect the v1 scope from `docs/PRD.md` and `docs/UI_PLAN.md`:

- **Do NOT add**:
  - New windows, dialogs, notification toasts, or additional UI surfaces.
  - A settings/preferences window for v1.
  - Graphs, history charts, or complex statistics surfaces.
- Tooltip must remain:
  - **Single line** in the core v1 format: `↓ ...    ↑ ...`
  - With **exact spacing rules** and update every 1 second.
- Context menu must remain:
  - `NetFlow` (disabled label), separator, `Quit`.

If a future version requires more UI:

- Update `docs/PRD.md` and `docs/UI_PLAN.md` first.
- Clearly mark new behavior as **v2+** and keep v1 behavior intact and available.

---

### 3. Performance, privacy, and safety

Any code change must preserve these guardrails:

- **Performance**
  - Idle CPU < 1%, active CPU < 3%.
  - Memory usage < 20 MB.
  - 1-second update interval remains fixed unless explicitly changed in PRD.
- **Privacy**
  - No telemetry, analytics, or external network calls.
  - No collection or transmission of user traffic patterns.
- **Safety**
  - No additional permissions (no admin/root).
  - Be cautious with panics: prefer explicit error handling and graceful shutdown.

If a proposed change risks these constraints, it must be clearly documented and justified in `docs/PRD.md` before implementation.

---

### 4. Coding style and Go best practices

Follow idiomatic Go conventions:

- Use clear, descriptive names (`downloadBytes`, `uploadBytes`, `updateTooltip`).
- Handle errors explicitly; avoid ignoring them with `_` unless safe and documented.
- Keep functions short and focused; extract helper functions when logic grows.
- Use standard `go fmt` formatting.

Testing and safety:

- Prefer adding or updating tests when touching non-trivial logic (formatting, unit conversion, network calculations).
- For formatting behavior, add tests that assert **exact strings** (including spaces and units) to prevent regressions.

---

### 5. Working with UI & design rules

The UI/UX for NetFlow is very constrained and must feel:

- Calm
- Predictable
- Native to each OS
- Legally distinct from any proprietary UI systems

For any visual or interaction changes:

- Align with `docs/UI_PLAN.md` **and** the broader UI rules you have provided (premium, calm, OS-inspired but not copied).
- Do not introduce:
  - Custom fonts or bundled font files.
  - Strong gradients, heavy shadows, or flashy animations.
  - Platform-specific gimmicks that break cross-platform consistency.

If in doubt, prefer **no visual change** and keep behavior aligned with the current specification.

---

### 6. How AI should approach new tasks

When receiving a new request related to this project, AI assistants should:

1. **Clarify internally which layer is affected**:
   - Network, formatter, tray, config, build tooling, or docs.
2. **Re-check the relevant docs**:
   - Product/UX questions → `docs/PRD.md`, `docs/UI_PLAN.md`.
   - High-level context → `docs/context/ai-overview.md`.
   - Guardrails → this file.
3. **Prefer minimal, reversible changes** for the first implementation.
4. **Update documentation** in `docs/` when behavior or public expectations change.

Use this document as a checklist before and after changes to ensure NetFlow stays simple, reliable, and aligned with its original vision.

