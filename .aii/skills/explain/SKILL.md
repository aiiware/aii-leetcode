---
name: explain
description: Explain code, architecture, or concepts in depth
allowed-tools: Read Grep Glob
argument-hint: "[file, function, or concept]"
---

# Explain

Provide a clear, thorough explanation of the specified code, architecture, or concept.

## Input

`$ARGUMENTS` specifies what to explain:

- **File path** — `src/auth/middleware.ts` — Explain the file's purpose and logic
- **Function/class** — `parseConfig` — Find and explain the implementation
- **Concept** — `"how authentication works"` — Trace the flow through the codebase
- **Directory** — `src/api/` — Explain the module's architecture

## Explanation Structure

### For Code (file, function, class)

1. **Purpose** — What does this code do and why does it exist?
2. **How It Works** — Walk through the logic step by step
3. **Key Decisions** — Why was it built this way? What tradeoffs were made?
4. **Dependencies** — What does it depend on? What depends on it?
5. **Edge Cases** — What special cases does it handle (or miss)?

### For Architecture (directory, system, flow)

1. **Overview** — High-level description of the system/module
2. **Components** — Key parts and their responsibilities
3. **Data Flow** — How data moves through the system (include text diagram)
4. **Patterns** — Design patterns and architectural choices
5. **Extension Points** — How to add new functionality

### For Concepts

1. **Definition** — Clear, jargon-free explanation
2. **Context** — How this concept applies in this codebase
3. **Examples** — Concrete code examples from the project
4. **Related** — Connected concepts worth understanding

## Text Diagrams

Use ASCII/Unicode text diagrams to visualize relationships and flows. Choose the appropriate diagram type for what you're explaining:

### Component / Layer Diagram

```text
┌─────────────────────────────────┐
│          Presentation           │
├─────────────────────────────────┤
│         Business Logic          │
├─────────────────────────────────┤
│          Data Access            │
└─────────────────────────────────┘
```

### Data Flow / Sequence Diagram

```text
Client ──request──▶ Router ──dispatch──▶ Handler
                                            │
                                        validates
                                            │
                                            ▼
Client ◀──response── Router ◀──result── Service
```

### Dependency / Tree Diagram

```text
App
├── AuthModule
│   ├── LoginHandler
│   └── TokenService → JWTLibrary
├── APIModule
│   ├── Router
│   └── Middleware → AuthModule
└── Database
    ├── ConnectionPool
    └── Migrations
```

### State / Lifecycle Diagram

```text
[idle] ──start──▶ [running] ──complete──▶ [done]
                      │                      ▲
                    error                  retry
                      │                      │
                      ▼                      │
                  [failed] ────────────────────
```

### When to Use Diagrams

| Explaining | Diagram Type |
| ---------- | ------------ |
| System architecture, module boundaries | Component / Layer |
| Request handling, event processing | Data Flow / Sequence |
| File structure, dependency trees, class hierarchies | Dependency / Tree |
| Process lifecycles, state machines, workflows | State / Lifecycle |

Include at least one diagram when explaining architecture, data flow, or multi-component systems. For single functions or simple concepts, diagrams are optional.

## Guidelines

- Start with the simplest accurate explanation, then add depth
- Use code references (file:line) for every claim
- Prefer concrete examples over abstract descriptions
- Include text diagrams for architecture, flows, and multi-component systems
- Adapt detail level to the complexity of the target
- If the target is ambiguous, search the codebase to find the most likely match
- Do not suggest changes — this skill is for understanding, not modification
