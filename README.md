# Mog Lang

Mog Lang is a small interpreted programming language, built from scratch in Go. It comes with its own lexer, parser, and interpreter, and ships as a CLI tool called `amogh`.

```
let x = 42
let y = 69
print(x)
print(y)
```

```
$ amogh runs hello.mog
42
69
```

## Features

- `let` statements for variable assignment
- `print(...)` for output
- Multiple statements per file, in any order
- Simple, readable Go implementation (lexer → parser → AST → interpreter)

## Installation

Mog Lang is distributed as Go source and installed with `go install`. You'll need [Go](https://go.dev/dl/) installed first.

```bash
go install github.com/amoghhalemani/Mog_Lang/cmd/amogh@latest
```

This downloads the source, compiles it, and places an `amogh` binary in your Go bin directory.

### "amogh: command not found"?

`go install` places the binary in your Go bin folder, but that folder isn't always on your system's PATH by default. Here's how to fix it:

**Find your Go bin folder:**
```bash
go env GOPATH
```
Your binary is in `<that path>/bin`.

**macOS / Linux (zsh):**
```bash
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.zshrc
source ~/.zshrc
```

**macOS / Linux (bash):**
```bash
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.bashrc
source ~/.bashrc
```

**Windows:**
1. Press the Windows key and search for **"Edit the system environment variables"**
2. Click **Environment Variables...**
3. Under "User variables," select **Path** → **Edit** → **New**
4. Add your Go bin path (usually `C:\Users\<yourname>\go\bin`)
5. Click OK on everything, then close and reopen your terminal

**Verify it worked:**
```bash
amogh runs hello.mog
```

## Usage

Create a `.mog` file:

```
let x = 42
print(x)
```

Run it:

```bash
amogh runs hello.mog
```

## How it works

Mog Lang follows the classic interpreter pipeline:

```
source code (.mog) → Lexer → tokens → Parser → AST → Interpreter → output
```

- **Lexer** (`lexer/`) — reads raw source text character by character and groups it into tokens (`KEYWORD`, `IDENTIFIER`, `NUMBER`, `EQUAL`, `LPAREN`, `RPAREN`, `EOF`, `ILLEGAL`).
- **Parser** (`parser/`) — reads the token stream and builds an AST (Abstract Syntax Tree) out of it: structured Go values like `LetStatement` and `PrintStatement` that represent what the code actually means.
- **Interpreter** (`interpreter/`) — walks the AST and executes it: storing variables in memory and printing output.

## Project structure

```
Mog_Lang/
├── cmd/
│   └── amogh/
│       └── main.go       # CLI entry point
├── lexer/
│   └── lexer.go
├── parser/
│   ├── ast.go             # AST node definitions (LetStatement, PrintStatement)
│   └── parser.go          # token → AST logic
└── interpreter/
    └── interpreter.go     # AST → execution logic
```

## Roadmap

Mog Lang is intentionally minimal right now. Planned next steps:

- [ ] Arithmetic expressions (`let z = 1 + 2`)
- [ ] Conditionals (`if` / `else`)
- [ ] Loops (`while` / `for`)
- [ ] Lists/arrays
- [ ] String support
- [ ] A small dogfooding project — a CLI to-do app written entirely in Mog Lang
- [ ] Pre-compiled binary releases (no Go installation required)

## Why?

Built as a learning project to understand how programming languages work under the hood — writing a lexer, recursive-descent parser, and tree-walking interpreter from scratch in Go.

## License

MIT
