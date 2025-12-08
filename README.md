# 💻 Gotask CLI

A lightweight, modular CLI tool written in **Go** to manage tasks efficiently from your terminal.  
It follows clean System Design principles and can be installed in your system be it mac/linux/windows.

---

## Folder Structure

```
gotask-cli/
├── cmd/                    # CLI entry points using Cobra
│   └── app/                # Individual command implementations
├── internal/               # Core business logic (not exposed to external packages)
│   ├── cmd/                # Internal command helpers
│   ├── components/         # Reusable core modules
│   ├── config/             # Configuration management
│   ├── constants/          # Shared constants
│   ├── factory/            # Factory
│   ├── filestorage/        # Local storage layer
│   ├── helper/             # Utilities and common helpers
│   └── models/             # Domain models and entities
├── scripts/                # OS-specific scripts for setup/build
│   ├── linux/
│   ├── mac/
│   └── windows/
├── Makefile                # Developer automation (build, test, lint, etc.)
├── go.mod / go.sum         # Go module management
└── README.md               # Documentation
```

---

## System Design & Patterns

### Factory Pattern
All dependencies (e.g., storage, executors) are created via a centralized **Factory**.  
This decouples object creation from command logic, improving maintainability and testability.

**Example:**
```go
fc := factory.New()
fs := fc.FileStorage()
taskExecutor := fc.TasksExecutor(fs)
```

### Command Pattern (via Cobra)
Each CLI action (like `add`, `list`, `remove`) is implemented as a separate command struct using spf13/cobra. This modularizes the CLI and makes it easy to add new commands.

---

## Makefile Usage

The project includes a `Makefile` to simplify building and running the CLI.

| Command | Description |
|---------|-------------|
| `make build` | Build the binary (`gotask`) |
| `make run` | Run the CLI locally |
| `make install` | Install the CLI binary globally |
| `make clean` | Remove build artifacts |
| `make lint` | Run static analysis (if configured) |

You can check available commands anytime by running:

```bash
make help
```

---

## Usage

After building or installing, you can run:

```bash
gotask add "Finish writing README"
```

Example output:

```
Task added: Finish writing README
```

---

## Future Commands (Planned)

| Command | Description |
|---------|-------------|
| `gotask list` | Show all tasks |
| `gotask remove [id]` | Remove a specific task |
| `gotask complete [id]` | Mark task as complete |

---

## Tech Stack

- **Language:** Go
- **CLI Framework:** Cobra
- **Design Patterns:** Factory, Command
- **Build Tool:** Make
- **Storage:** File-based storage (extensible to DBs)

---

## Development

```bash
# Clone the repo
git clone https://github.com/saianand32/gotask-cli.git
cd gotask-cli

# Install dependencies
go mod tidy

# Build binary
make build

# Run directly
./bin/gotask add "Test new CLI"
```

---

## 🪶 License

MIT © 2025 Saishwar Anand