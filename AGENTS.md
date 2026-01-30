# leetcode

<!-- AGENTS.md - Universal AI Agent Instructions -->

## Project Overview
This is a LeetCode problem solutions repository written in Go. It contains implementations of various algorithms, data structures, and SQL problems, likely for practice and reference. The project uses the standard Go toolchain for building and testing.

## Directory Structure (Reorganized - Jan 2026)
```
leetcode/
├── arrays/                    # Array manipulation, two pointers, sliding window (58 problems)
├── strings/                   # String manipulation, palindrome, matching (28 problems)
├── dp/                        # Dynamic programming problems (28 problems)
├── trees/                     # Binary trees, BST, traversal (21 problems)
├── math/                      # Math, bit manipulation, number theory (17 problems)
├── linked-lists/              # Linked list algorithms (14 problems)
├── sql/                       # SQL problems (9 problems)
├── sorting/                   # Sorting, searching, binary search (8 problems)
├── graphs/                    # Graph traversal, BFS/DFS (3 problems)
├── design/                    # System design, data structures (3 problems)
├── indexes/                   # Metadata and indexes (planned)
├── cmd/                       # Command entrypoints (Go)
├── utils/                     # Utility functions and shared types
├── data_structures/           # Data structure implementations
├── testutils/                 # Testing utilities
├── tutorials/                 # Tutorial content
└── temp_debug/                # Temporary debugging files
```

**Note**: The project has been reorganized (Jan 2026) from a flat structure to categorized directories for better learning organization. Each category is a separate Go package.

## File Patterns
*   **Naming Convention:** Source files use `camelCase` with LeetCode problem numbers (e.g., `0001_two_sum.go`).
*   **Test Files:** Follow the standard Go pattern: `*_test.go` (e.g., `0001_two_sum_test.go`).
*   **Organization:** Code is organized by algorithm category:
    *   Category directories (`arrays/`, `strings/`, etc.) for LeetCode problems
    *   `cmd/` for command-line entry points
    *   `utils/` for shared utility functions and types (ListNode, TreeNode)
    *   `data_structures/` for data structure implementations
    *   `tutorials/` for tutorial content
    *   `testutils/` for testing utilities
    *   `temp_debug/` for temporary debugging files

## Core Commands
*   **Build:** `go build ./...`
*   **Test:** `go test ./...`
*   **Test specific category:** `go test ./arrays/...`
*   **Run demo:** `go run cmd/demo/main.go`

## Git Conventions
*   **Convention:** Conventional Commits.
*   **Default Branch:** `main`.
*   **Common Commit Types & Scopes:**
    *   `feat(leetcode)`: Implement new LeetCode problems.
    *   `feat(aii)` / `feat(demo)`: Add features for AI or demos.
    *   `refactor(test)` / `refactor(leetcode)` / `refactor(graph)`: Refactor test code, problem solutions, or graph-related code.
    *   `docs(readme)`: Update project documentation.
    *   `fix(bit-manipulation)` / `fix(algorithm)` / `fix(binary)`: Correct bugs in specific algorithm categories.
    *   `test(binary-tree)` / `test(leetcode)`: Add or modify tests.
    *   `chore(gitignore)`: Update project maintenance files.
    *   `refactor(organization)`: Reorganize file structure (used for recent reorganization)

**Example Commit Messages:**
*   `feat(leetcode): implement LeetCode 0176-0178 SQL problems`
*   `fix(bit-manipulation): correct bitwise operations for negative numbers`
*   `refactor(test): rename test functions for clarity and update imports`
*   `test(binary-tree): simplify test case tree representations`
*   `docs(readme): update documentation for problems 0131-0150`
*   `refactor(organization): reorganize problems into algorithm categories`

## Coding Standards
*   **Language:** Go.
*   **File Naming:** Use `camelCase` for source files with LeetCode problem numbers.
*   **Testing:** All tests must be in `*_test.go` files and runnable via `go test`.
*   **Structure:** Follow the categorized directory pattern for organizing LeetCode problems by algorithm type.
*   **Package Naming:** Each category directory is its own Go package (e.g., `package arrays`, `package strings`).
*   **Shared Types:** Common types (ListNode, TreeNode) are in `utils/` package.

## Safety Rules
*   **Do not modify** the root `go.mod` file without explicit approval, as it manages all dependencies.
*   **Do not delete or significantly alter** the structure of the `data_structures/` or `tutorials/` directories without confirmation, as they contain core learning content.
*   **Require confirmation** before creating or modifying files in `temp_debug/`, as this directory is for transient files.
*   **Require confirmation** before altering the `cmd/` directory, as it contains application entry points.
*   **Be cautious with cross-package dependencies** - The recent reorganization created separate packages; avoid creating circular dependencies between categories.
*   **Preserve the categorization scheme** - When adding new problems, place them in the appropriate category directory.

## Permissions
*   **Allowed Without Prompt:**
    *   Create new `.go` solution files in appropriate category directories (e.g., `arrays/`, `strings/`).
    *   Create or update `*_test.go` files within category directories.
    *   Update utility functions in `utils/`.
    *   Update code within `testutils/`.
    *   Fix import statements and package declarations to maintain the new structure.
*   **Requires Confirmation:**
    *   Modifying the `go.mod` file.
    *   Changing the structure or core files in `cmd/`, `data_structures/`, or `tutorials/`.
    *   Operations in the `temp_debug/` directory.
    *   Changing established file naming conventions.
    *   Moving files between category directories.
    *   Creating new category directories.
*   **Never Allowed:**
    *   Deleting the `go.mod` file.
    *   Changing the project's fundamental organization (e.g., reverting to flat structure).
    *   Modifying git history or force-pushing.
    *   Creating circular dependencies between category packages.

## Current Reorganization Status (Jan 2026)
The project has recently been reorganized from a flat structure to categorized directories. Key changes:

1. **189 LeetCode problems** moved to 10 algorithm categories
2. **Shared types consolidated** into `utils/` package (ListNode, TreeNode, helper functions)
3. **Package structure updated** - Each category is now a separate Go package
4. **Demo program updated** - `cmd/demo/main.go` uses new package imports

**Known Issues** (see `BACKLOG.md` for details):
- Some test files have cross-package dependencies
- Some command-line tools need import updates
- Index files need to be generated

**Working packages**: `strings/`, `math/`, `linked-lists/`, `sql/`, `sorting/`, `graphs/`, `design/`
**Needs attention**: `arrays/`, `dp/`, `trees/` (cross-package test dependencies)

---

*Generated by Aii CLI /init command with LLM enhancement, updated Jan 2026 to reflect reorganization (~1200 tokens, 100 lines)*