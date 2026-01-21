# leetcode

<!-- AGENTS.md - Universal AI Agent Instructions -->

## Project Overview

A **Go** project, using **go test** for testing, managed with **go**.

This project uses Git for version control.

**Existing instruction files found:**
- `AGENTS.md` (130 lines)

## Directory Structure

```
├── utils/ # Utility functions
├── cmd/ # CLI entrypoints (Go)
├── data_structures/
├── tutorials/
└── testutils/
```

## File Patterns

### Naming Convention: camelCase

This project uses **camelCase** for file naming (68% of files).

### Numbered File Patterns

**Four-digit prefix with underscore (e.g., 0001_problem_name)**
- Pattern: `####_snake_case`
- Count: 315 files
- Examples: `0001_two_sum.go`, `0001_two_sum_test.go`, `0002_add_two_numbers.go`

### Test Files

Test files follow the pattern: *_test.{ext} (e.g., handler_test.go)

## Core Commands

- **Build**: `go build ./...`
- **Test**: `go test ./...`

## Git Conventions

### Commit Messages

This project uses **Conventional Commits** format (98% of commits).

**Common commit types:**
- `feat`: 27 commits (scopes: leetcode, demo)
- `refactor`: 10 commits (scopes: test, leetcode, tests)
- `docs`: 6 commits (scopes: readme)
- `fix`: 3 commits (scopes: bit-manipulation, algorithm, binary)
- `test`: 2 commits (scopes: binary-tree, leetcode)
- `chore`: 1 commits (scopes: gitignore)

**Common scopes:**
`leetcode`, `readme`, `tests`, `test`, `bit-manipulation`, `algorithm`, `binary-tree`, `graph`

**Example commits:**
```
fix(bit-manipulation): correct bitwise operations for negative numbers
refactor(test): rename test functions for clarity and update imports
feat(leetcode): implement LeetCode 189 Rotate Array with multiple solutions
test(binary-tree): simplify test case tree representations
docs(readme): update documentation for problems 0131-0150
```

## Coding Standards

### Language: Go

- Follow Go formatting conventions (use `gofmt`)
- Handle errors explicitly, avoid panic
- Use meaningful variable names
- Keep functions focused and small

### File Naming

- Use **camelCase** for file names
- Example: `myComponent.ts`, `userService.ts`
- Test files: *_test.{ext} (e.g., handler_test.go)

## LeetCode Problem Implementation Workflow

### Standard File Structure

Each LeetCode problem should have:
1. **Implementation File**: `####_problem_name.go`
2. **Test File**: `####_problem_name_test.go`

### File Content Template

#### Implementation File (`####_problem_name.go`):
```go
package leetcode

/*
###. Problem Name

[Problem description from LeetCode]

Example 1:
Input: [example input]
Output: [example output]
Explanation: [explanation]

Example 2:
Input: [example input]
Output: [example output]
Explanation: [explanation]

Constraints:
- [constraint 1]
- [constraint 2]

Difficulty: [Easy/Medium/Hard]
Tags: [tag1, tag2, tag3]
Companies: [Company1, Company2, Company3]
*/

// Implementation code here
```

#### Test File (`####_problem_name_test.go`):
```go
package leetcode

import (
    "fmt"
    "testing"
)

func TestFunctionName(t *testing.T) {
    tests := []struct {
        input    [type]
        expected [type]
    }{
        // Test cases from LeetCode examples
        // Additional edge cases
        // Large input cases
    }

    for i, tt := range tests {
        t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
            result := functionName(tt.input)
            if result != tt.expected {
                t.Errorf("functionName(%v) = %v, want %v", tt.input, result, tt.expected)
            }
        })
    }
}

// Benchmark tests for performance
func BenchmarkFunctionName(b *testing.B) {
    // Benchmark code
}
```

### Metadata Collection

When implementing new LeetCode problems, gather:
1. **Problem Description**: Full problem statement
2. **Examples**: All provided examples with explanations
3. **Constraints**: All constraints
4. **Difficulty**: Easy, Medium, or Hard
5. **Tags**: Algorithm/data structure tags
6. **Companies**: Companies that have asked this problem (if available)

### Implementation Guidelines

1. **Read existing files** to understand patterns before creating new ones
2. **Use proper Go idioms** and error handling
3. **Include comprehensive tests** covering:
   - LeetCode examples
   - Edge cases
   - Large inputs
   - Performance benchmarks
4. **Add comments** for complex algorithms
5. **Follow naming conventions**:
   - Function names: camelCase
   - Test names: TestFunctionName (ensure unique names across files)
   - Benchmark names: BenchmarkFunctionName (ensure unique names across files)
6. **Package declaration**: Always use `package leetcode` (not `package main`)

### Automatic Testing Rule

**CRITICAL**: After implementing or modifying any code, ALWAYS run tests to verify correctness:

1. **Immediate verification**: Run `go test -v .` after creating/updating files
2. **Comprehensive testing**: Run `go test ./...` to test all packages
3. **Specific test verification**: For new implementations, run `go test -v -run "TestFunctionName" .`
4. **Fix issues immediately**: If tests fail, fix the issues before proceeding

### Commit Message Format

For LeetCode implementations:
```
feat(leetcode): implement LeetCode ### Problem Name

- Add solution with [algorithm/data structure]
- Include comprehensive test cases
- Add performance benchmarks
```

## Safety Rules

### Never Modify Without Approval

- `.env` files and environment configurations
- Secrets, API keys, or credentials
- Database migration files
- CI/CD configuration (`.github/`, `.gitlab-ci.yml`, etc.)

### Always Ask Before

- Deleting any file or directory
- Running commands with `sudo` or elevated privileges
- Installing new dependencies
- Making breaking API changes
- Modifying authentication or authorization logic
- Changing database schemas

## Permissions

### Allowed Without Prompt

- Read any file in the repository
- Explore directory structure
- Run `go test ./...`
- Run `go build ./...`

### Requires Confirmation

- File writes outside of `src/` directory
- Git operations (commit, push, branch, merge)
- Network requests to external services
- Running arbitrary shell commands
- Modifying configuration files

### Never Allowed

- Exposing secrets or credentials
- Pushing directly to main/master branch
- Force pushing to any branch
- Accessing files outside the project directory
---

*Generated by Aii CLI /init command (~878 tokens, 127 lines)*
*Updated with LeetCode workflow on 2026-01-21*
*Updated with automatic testing rule and package fix on 2026-01-21*