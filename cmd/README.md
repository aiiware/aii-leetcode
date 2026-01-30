# Command Line Tools

This directory contains command-line tools and demo programs for the LeetCode project.

## 📁 Directory Structure

```
cmd/
├── README.md              # This file
├── analyze/              # Analysis tools
│   └── main.go
├── brute_force/          # Brute force algorithm demos
│   └── main.go
├── debug/                # Debugging tools
│   └── main.go
├── debug_tree/           # Tree debugging tools
│   └── main.go
├── demo/                 # Main demo program
│   └── main.go
├── nqueens_demo/         # N-Queens demo
│   └── main.go
├── nqueens_ii_demo/      # N-Queens II demo
│   └── main.go
├── quadtree_demo/        # QuadTree demo
│   └── main.go
├── test_interleave_debug/ # Interleaving string debug
│   └── main.go
└── tree_debug/           # Tree structure debug
    └── main.go
```

## 🚀 Building and Running

Each directory contains a standalone Go program. To build and run:

```bash
# Build and run a specific tool
go run cmd/demo/main.go

# Build all tools
go build ./cmd/...

# Build a specific tool
go build -o demo ./cmd/demo
```

## 🔧 Available Tools

### `demo/` - Main Demo Program
Runs examples from various algorithm categories.

### `debug/` - General Debugging
General debugging utilities for algorithm testing.

### `debug_tree/` - Tree Debugging
Tools for debugging tree algorithms and structures.

### `tree_debug/` - Tree Structure Debug
Additional tree structure debugging tools.

### `analyze/` - Analysis Tools
Algorithm analysis and visualization tools.

### `brute_force/` - Brute Force Demos
Demonstrations of brute force algorithms for comparison.

### `nqueens_demo/` - N-Queens Demo
N-Queens problem solver demonstration.

### `nqueens_ii_demo/` - N-Queens II Demo
N-Queens II problem solver demonstration.

### `quadtree_demo/` - QuadTree Demo
QuadTree data structure demonstration.

### `test_interleave_debug/` - Interleaving Debug
Debugging tools for string interleaving problems.

## 📝 Conventions

- Each directory contains exactly one `main.go` file
- All tools are standalone and don't depend on each other
- Tools should have clear, descriptive names
- Include usage instructions in comments at the top of each `main.go`

## 🔄 Adding New Tools

When adding a new command-line tool:

1. Create a new directory under `cmd/` with a descriptive name
2. Create a `main.go` file with package `main`
3. Add a brief description at the top of the file
4. Update this README with the new tool's description

---

*Last updated: January 2026*