# hg

> An interactive shell history explorer and launcher written in Go.

`hg` is a command-line tool for searching, exploring, copying, and eventually reusing commands from your shell history.

The project is primarily built as a **Go learning project**, while aiming to become a practical daily-driver utility.

## Features

### Current

* [x] Parse shell history
  * [x] Zsh
  * [x] Bash
* [x] Preserve command timestamps
* [x] Fuzzy command search
* [x] Fuzzy match scores
* [x] Matched character indexes for highlighting
* [x] Clipboard abstraction
* [x] Copy matching commands from the CLI

### Planned

* [ ] Interactive TUI
* [ ] Realtime fuzzy search
* [ ] Keyboard navigation
* [ ] Matched character highlighting
* [ ] Interactive command selection
* [ ] Copy selected command
* [ ] Execute selected command
* [ ] Command preview
* [ ] Shell integration
* [ ] Configurable keybindings
* [ ] Search and performance improvements
* [ ] Comprehensive tests and benchmarks
* [ ] Cross-platform support
* [ ] Release binaries and packaging

See [ROADMAP.md](ROADMAP.md) for the complete development roadmap.

## Motivation

Shell history is useful, but traditional history search can become cumbersome when the history contains thousands of commands.

For example:

```bash
history | grep docker
```

works well for simple searches, but it does not provide an interactive way to:

* browse matching commands,
* navigate between results,
* see when a command was executed,
* highlight fuzzy matches,
* copy a selected command,
* or eventually execute a selected command.

`hg` aims to provide these capabilities through an interactive terminal interface.

The goal is not to replace `grep`. Instead, `hg` is intended to become an **interactive history explorer and launcher**.

## Example

The current CLI can search shell history using fuzzy matching:

```bash
hg docker
```

Conceptually, this searches commands such as:

```text
docker compose up -d
docker ps
docker exec -it postgres bash
docker compose down
```

Copy mode is also available:

```bash
hg -c docker
```

The interactive copy workflow will eventually allow the user to select which matching command should be copied.

## Architecture

The project separates history parsing, searching, presentation, and external actions.

```text
                    Shell history
                         │
                         ▼
                  ┌─────────────┐
                  │   history   │
                  │   package   │
                  └──────┬──────┘
                         │
                         ▼
                  []history.Entry
                  ┌─────────────┐
                  │ Command     │
                  │ Timestamp   │
                  └──────┬──────┘
                         │
                         ▼
                  ┌─────────────┐
                  │   search    │
                  │   package   │
                  └──────┬──────┘
                         │
                         ▼
                  []search.Result
                  ┌─────────────┐
                  │ Entry       │
                  │ Score       │
                  │ Matched...  │
                  └──────┬──────┘
                         │
                         ▼
                  ┌─────────────┐
                  │     app     │
                  │    / TUI    │
                  └──────┬──────┘
                         │
                ┌────────┴────────┐
                ▼                 ▼
           Clipboard          Execute
```

### `history`

Responsible for reading and parsing shell history files.

The package converts shell-specific history formats into a common representation:

```go
type Entry struct {
    Command   string
    Timestamp time.Time
}
```

The rest of the application does not need to know whether an entry came from Bash or Zsh.

### `search`

Responsible for fuzzy matching history entries.

A search result contains information that belongs specifically to a search operation:

```go
type Result struct {
    Entry          history.Entry
    Score          int
    MatchedIndexes []int
}
```

`MatchedIndexes` is intended to be used by the TUI to highlight matching characters.

### `clipboard`

Provides a small abstraction around clipboard operations so the rest of the application does not need to depend directly on a particular clipboard implementation.

### `app`

The TUI layer will eventually own:

* query state,
* selected result,
* keyboard input,
* rendering,
* actions,
* and interaction with search and clipboard packages.

## Project Structure

```text
hg/
├── README.md
├── ROADMAP.md
├── go.mod
├── cmd/
│   └── hg/
│       └── main.go
└── internal/
    ├── tui/
    │   ├── keys.go
    │   ├── model.go
    │   ├── update.go
    │   └── view.go
    ├── clipboard/
    │   └── clipboard.go
    ├── history/
    │   ├── entry.go
    │   ├── provider.go
    │   ├── bash.go
    │   └── zsh.go
    └── search/
        └── fuzzy.go
```

Some components may change as the project evolves.

## Technology

* **Go**
* **Bubble Tea** — terminal user interface
* **sahilm/fuzzy** — fuzzy matching
* **golang.design/x/clipboard** — clipboard functionality

The exact dependencies may change during development.

## Development

Clone the repository:

```bash
git clone https://github.com/hallwack/hg.git
cd hg
```

Run the application:

```bash
go run ./cmd/hg
```

Search history:

```bash
go run ./cmd/hg docker
```

Build the binary:

```bash
go build -o hg ./cmd/hg
```

Run tests:

```bash
go test ./...
```

Run benchmarks:

```bash
go test -bench=. ./...
```

## Design Goals

`hg` is being developed with a few principles in mind:

### Keep shell-specific logic isolated

Bash and Zsh have different history formats. That complexity should stay inside the `history` package.

### Keep search independent from presentation

The search package should return structured results rather than printing terminal output.

This makes it possible to use the same search logic from both the CLI and TUI.

### Keep the TUI independent from the data source

The TUI should operate on `history.Entry` and `search.Result`, rather than knowing how history was loaded.

### Prefer small abstractions

Interfaces and abstractions should be introduced when they solve a real problem, rather than designing the entire application around hypothetical future requirements.

### Learn Go through real problems

The project is intentionally expected to evolve as new Go concepts are encountered, including:

* interfaces,
* package design,
* error handling,
* concurrency,
* process management,
* testing,
* benchmarking,
* cross-platform builds,
* and terminal applications.

## License

MIT
