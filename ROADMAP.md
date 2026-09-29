# Roadmap

The development of `hg` is divided into several milestones.

The project is both a practical shell utility and a learning project, so the roadmap intentionally progresses from simple components toward more advanced terminal and system-level functionality.

---

## v0.1 — History & Search

**Status: Mostly complete**

Establish the core data model and history-search functionality.

### History

* [x] Create Go module
* [x] Define `history.Entry`
* [x] Implement shell history loading
* [x] Implement Zsh history parser
* [x] Implement Bash history parser
* [x] Support multiline history entries
* [x] Parse command timestamps
* [ ] Improve shell detection
* [ ] Improve parser error handling

### Search

* [x] Implement fuzzy search
* [x] Return structured search results
* [x] Preserve original `history.Entry`
* [x] Expose fuzzy score
* [x] Expose matched character indexes
* [ ] Add result limiting
* [ ] Improve search behavior for large histories

### Clipboard

* [x] Create clipboard abstraction
* [x] Implement copy operation
* [x] Ensure copied content remains available after `hg` exits
* [ ] Finalize clipboard behavior across platforms

---

# v0.2 — Interactive TUI

**Status: Next**

Build the first usable Bubble Tea interface.

### TUI foundation

* [x] Introduce Bubble Tea
* [x] Create application model
* [x] Create update loop
* [x] Create view
* [x] Render history results
* [x] Render query input
* [ ] Render empty state
* [ ] Render error state

### Search interaction

* [x] Search while typing
* [x] Update results in realtime
* [ ] Preserve fuzzy result ordering
* [ ] Display result count

### Navigation

* [ ] Move selection with `↑` / `↓`
* [ ] Support `j` / `k`
* [ ] Keep selection within result bounds
* [ ] Handle empty result sets

### Result rendering

* [ ] Display timestamps
* [ ] Highlight `MatchedIndexes`
* [ ] Handle long commands
* [ ] Handle multiline commands
* [ ] Add scrolling

### Exit behavior

* [ ] `Esc` quits
* [ ] `Ctrl+C` quits
* [ ] Clean up terminal state correctly

### Milestone

The following workflow should work:

```text
hg
 │
 ▼
TUI
 │
 ├── type query
 ├── fuzzy search
 ├── ↑ / ↓
 └── highlight matches
```

---

# v0.3 — Selection & Actions

**Status: Planned**

Turn the TUI from a browser into an interactive history tool.

### Selection

* [ ] Select a history entry
* [ ] Expose selected `search.Result`
* [ ] Add selected-entry preview

### Copy

* [ ] Copy selected command
* [ ] Exit after successful copy
* [ ] Display copy confirmation
* [ ] Handle clipboard failures

Expected workflow:

```text
hg -c docker
      │
      ▼
    TUI
      │
      ▼
 select command
      │
      ▼
    Enter
      │
      ▼
 copy command
      │
      ▼
   exit hg
```

### Execute

* [ ] Execute selected command
* [ ] Pass stdin/stdout/stderr correctly
* [ ] Preserve exit status
* [ ] Handle command failures
* [ ] Decide whether execution requires confirmation

Possible keybinding:

```text
Enter       Copy
Ctrl+Enter  Execute
Esc         Cancel
```

The exact keybindings may change.

---

# v0.4 — Command Preview

**Status: Planned**

Improve usability for long and multiline commands.

### Preview

* [ ] Add command preview
* [ ] Show complete selected command
* [ ] Support multiline commands
* [ ] Scroll long previews
* [ ] Display command metadata

Example:

```text
┌─ History ───────────────────────────────────┐
│ > docker compose up                         │
│   git status                                │
│   go test ./...                             │
│                                            │
├─ Preview ──────────────────────────────────┤
│ docker compose -f docker-compose.prod.yml  │
│ up -d --remove-orphans                     │
└────────────────────────────────────────────┘
```

---

# v0.5 — Shell Integration

**Status: Planned**

Allow a selected history command to be returned to the shell without immediately executing it.

Desired workflow:

```text
hg
 │
 ▼
select command
 │
 ▼
insert into shell prompt
 │
 ▼
user edits command
 │
 ▼
user decides whether to execute
```

### Zsh

* [ ] Design shell integration API
* [ ] Implement Zsh integration
* [ ] Insert selected command into prompt

### Bash

* [ ] Implement Bash integration
* [ ] Insert selected command into prompt

### Other shells

* [ ] Investigate Fish
* [ ] Investigate PowerShell
* [ ] Evaluate whether each shell should be supported

---

# v0.6 — Search & UX Improvements

**Status: Planned**

Improve the search experience after the core workflow is stable.

### Search

* [ ] Case-insensitive search
* [ ] Search result limit
* [ ] Better handling of empty queries
* [ ] Search performance improvements
* [ ] Investigate exact-match mode
* [ ] Investigate prefix/suffix matching

### UX

* [ ] Status bar
* [ ] Result counter
* [ ] Better error messages
* [ ] Configurable timestamp format
* [ ] Configurable UI layout
* [ ] Configurable keybindings

---

# v0.7 — Configuration

**Status: Planned**

Introduce persistent configuration only after the default behavior has stabilized.

Possible configuration location:

```text
~/.config/hg/config.toml
```

Possible configuration:

```toml
[ui]
show_timestamp = true
show_duration = false
max_results = 100

[keys]
copy = "enter"
execute = "ctrl-e"
quit = "esc"
```

### Tasks

* [ ] Define configuration format
* [ ] Implement config loading
* [ ] Implement defaults
* [ ] Validate configuration
* [ ] Document configuration options

---

# v0.8 — Testing & Performance

**Status: Planned**

Strengthen the internal implementation before a stable release.

### Unit tests

* [ ] Zsh parser tests
* [ ] Bash parser tests
* [ ] Multiline parser tests
* [ ] Timestamp parsing tests
* [ ] Fuzzy search tests
* [ ] Result mapping tests
* [ ] Clipboard tests where practical

### Integration tests

* [ ] History loading integration tests
* [ ] CLI behavior tests
* [ ] Shell integration tests

### Benchmarks

* [ ] Benchmark history parsing
* [ ] Benchmark fuzzy search
* [ ] Test with 1,000 entries
* [ ] Test with 10,000 entries
* [ ] Test with 100,000 entries

Example:

```bash
go test -bench=. ./...
```

---

# v0.9 — Cross-Platform Support

**Status: Planned**

Expand support beyond the initial development environment.

### Linux

* [ ] X11
* [ ] Wayland
* [ ] Bash
* [ ] Zsh

### macOS

* [ ] History loading
* [ ] Clipboard
* [ ] Terminal behavior

### Windows

* [ ] PowerShell history
* [ ] CMD history investigation
* [ ] Clipboard
* [ ] Terminal behavior

### Other platforms

Evaluate additional platforms based on demand and implementation complexity.

---

# v1.0 — First Stable Release

**Status: Future**

A stable release should provide a coherent end-to-end workflow rather than simply a large number of features.

### Core

* [ ] Stable history parsing
* [ ] Stable fuzzy search
* [ ] Interactive TUI
* [ ] Copy selected command
* [ ] Execute selected command
* [ ] Command preview

### Shell integration

* [ ] Zsh
* [ ] Bash

### Quality

* [ ] Unit tests
* [ ] Integration tests
* [ ] Benchmarks
* [ ] Error handling review
* [ ] Documentation review

### Distribution

* [ ] Linux binaries
* [ ] macOS binaries
* [ ] Windows binaries
* [ ] GitHub Releases
* [ ] GoReleaser
* [ ] Nix package

---

# Future Ideas

These ideas are intentionally outside the core roadmap.

They should only be implemented if they provide a clear benefit.

* [ ] Command frequency
* [ ] Recently used commands
* [ ] Command categories
* [ ] History statistics
* [ ] Delete history entry
* [ ] Pin/favorite commands
* [ ] Command aliases
* [ ] Multiple history sources
* [ ] Git-aware command context
* [ ] Working-directory metadata
* [ ] Exit-code metadata
* [ ] Command duration
* [ ] History import/export
* [ ] Plugin system

---

# Development Principles

## 1. Build the smallest useful version

Avoid implementing advanced features before the core workflow is usable.

```text
History
   ↓
Search
   ↓
TUI
   ↓
Select
   ↓
Copy
   ↓
Execute
```

## 2. Keep packages focused

The intended boundaries are:

```text
history
  └── Read and parse history

search
  └── Find matching entries

clipboard
  └── Copy text

app
  └── TUI state and interaction

cmd/hg
  └── Application entry point
```

## 3. Avoid leaking implementation details

For example, the TUI should not need to know whether history came from:

```text
.zsh_history
.bash_history
PowerShell
```

It should receive a common `history.Entry`.

## 4. Prefer simple interfaces

Introduce an interface when it provides a meaningful boundary or makes testing easier.

Avoid creating abstractions only because a future feature might need them.

## 5. Optimize after measuring

Performance improvements should be guided by benchmarks rather than assumptions.

---

# Current Focus

The immediate milestone is:

## `v0.2 — Interactive TUI`

The next implementation steps are:

1. Add Bubble Tea.
2. Create the application `Model`.
3. Render `[]search.Result`.
4. Add query input.
5. Re-run fuzzy search when the query changes.
6. Add cursor navigation.
7. Render timestamps.
8. Highlight `MatchedIndexes`.
9. Add scrolling.
10. Implement clean exit behavior.

Once this works, move to `v0.3 — Selection & Actions`.
