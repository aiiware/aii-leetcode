---
name: code-review
description: Review code for bugs, security issues, and best practices
allowed-tools: Read Grep Glob Bash
argument-hint: "[file or directory]"
---

# Code Review

Review the specified code for quality, correctness, and security. Provide actionable feedback organized by severity.

## What to Review

Analyze the code at `$ARGUMENTS` (file path, directory, or git diff). If no argument is provided, review staged changes (`git diff --cached`) or the most recent commit.

## Review Checklist

### 1. Correctness
- Logic errors, off-by-one, null/undefined handling
- Edge cases not covered
- Race conditions or concurrency issues
- Incorrect API usage or contract violations

### 2. Security
- Injection vulnerabilities (SQL, XSS, command injection)
- Hardcoded secrets or credentials
- Improper input validation or sanitization
- Insecure defaults or missing authentication checks
- Path traversal or directory escape

### 3. Performance
- Unnecessary allocations or copies in hot paths
- N+1 query patterns
- Missing indexes or inefficient data structures
- Unbounded growth (memory leaks, growing arrays)

### 4. Design Principles (S.O.L.I.D. & Beyond)

- **Single Responsibility** — Does each class/function/module do exactly one thing?
- **Open/Closed** — Is the code open for extension but closed for modification? Can new behavior be added without changing existing code?
- **Liskov Substitution** — Can subtypes be used interchangeably with their base types without breaking behavior?
- **Interface Segregation** — Are interfaces focused and minimal? No client should depend on methods it doesn't use
- **Dependency Inversion** — Does code depend on abstractions, not concrete implementations?
- **DRY** — Is logic duplicated in 3+ places that should be extracted?
- **KISS** — Is there unnecessary complexity? Could the same result be achieved more simply?
- **Separation of Concerns** — Are distinct responsibilities mixed in one place (e.g., business logic in UI layer, I/O in pure functions)?
- **Fail Fast** — Are inputs validated early with clear error messages, rather than failing deep in the call stack?
- **Contract-First** — Do public APIs have clear, stable contracts (types, schemas, documented behavior)?

### 5. Maintainability
- Unclear naming or misleading abstractions
- Overly complex logic that could be simplified
- Dead code or unused imports
- Missing error handling or swallowed errors
- Proper type hints/annotations throughout
- Comments explain "why", not "what"

### 6. Best Practices
- Adherence to project conventions (check AGENTS.md if present)
- Consistent style with surrounding code
- Appropriate test coverage for changes (aim for 90%+)
- Proper use of language/framework idioms
- Async/await correctness (no fire-and-forget, proper error propagation)
- No over-engineering — only the complexity the current task requires

## Output Format

Organize findings by severity:

**CRITICAL** — Must fix before merging (bugs, security vulnerabilities)

**WARNING** — Should fix (performance issues, code smells)

**SUGGESTION** — Nice to have (style, minor improvements)

For each finding:
1. File and line reference
2. What the issue is
3. Why it matters
4. Suggested fix (with code snippet when helpful)

End with a summary: total findings by severity, overall assessment (approve / request changes), and any positive observations worth noting.
