# LeetCode Solution Explanation Template

Use this template for all solution explanations to maintain consistency across the repository.

## File Naming Convention
- Format: `{problem_number}_{problem_name_slug}.md`
- Example: `0070_climbing_stairs.md`, `0200_number_of_islands.md`
- Use leading zeros for problem numbers (e.g., 0070 not 70)

## Template Structure

```markdown
# {Problem Number}. {Problem Name} - Solution Explanation

## Problem Statement
[Copy the problem statement from LeetCode]

## Difficulty: {Easy/Medium/Hard}

## Key Insights
1. **Insight 1**: [Core concept or observation]
2. **Insight 2**: [Algorithm choice rationale]
3. **Insight 3**: [Complexity considerations]

## Solution Approaches

### Approach 1: {Approach Name}
**Time Complexity**: O(...)
**Space Complexity**: O(...)

```go
// Go implementation
func solutionName(params) returnType {
    // Implementation
}
```

### Approach 2: {Alternative Approach Name} (Optional)
**Time Complexity**: O(...)
**Space Complexity**: O(...)

```go
// Alternative implementation
func alternativeSolution(params) returnType {
    // Implementation
}
```

## Step-by-Step Walkthrough

### Example:
```
[Provide a concrete example]
```

**Step 1**: [First step description]
- [Detail 1]
- [Detail 2]

**Step 2**: [Second step description]
- [Detail 1]
- [Detail 2]

**Step 3**: [Third step description]
- [Detail 1]
- [Detail 2]

**Final Result**: [Expected output for the example]

## Complexity Analysis
### Time Complexity
- **Best Case**: O(...) - [When]
- **Average Case**: O(...) - [Why]
- **Worst Case**: O(...) - [When]

### Space Complexity
- **Auxiliary Space**: O(...) - [Extra space used]
- **Total Space**: O(...) - [Including input]

## Common Pitfalls
1. **Pitfall 1**: [Common mistake and how to avoid it]
2. **Pitfall 2**: [Edge case handling]
3. **Pitfall 3**: [Performance optimization]

## Optimization Tips
1. **Tip 1**: [How to optimize time]
2. **Tip 2**: [How to optimize space]
3. **Tip 3**: [Alternative data structures]

## Edge Cases
1. [Edge case 1 description]
2. [Edge case 2 description]
3. [Edge case 3 description]

## Related Problems
- [Related problem 1 with link]
- [Related problem 2 with link]
- [Related problem 3 with link]

## Practice Exercises
1. [Exercise 1: Modify to solve a variant]
2. [Exercise 2: Implement alternative approach]
3. [Exercise 3: Analyze time/space tradeoffs]

## Additional Notes
- [Any additional context, historical notes, or advanced topics]
```

## Category-Specific Sections

### For Dynamic Programming Problems
Add these sections:
- **State Definition**: `dp[i]` or `dp[i][j]` represents...
- **Transition Equation**: How to compute current state from previous states
- **Base Cases**: Initial values for the DP table
- **Optimization**: Space optimization from 2D to 1D if applicable

### For Graph Problems
Add these sections:
- **Graph Representation**: Adjacency list vs matrix choice
- **Traversal Strategy**: BFS vs DFS choice and why
- **Cycle Detection**: How cycles are handled
- **Connected Components**: Approach for finding components

### For Array/String Problems
Add these sections:
- **Two Pointers Strategy**: When and how to use
- **Sliding Window**: Window management strategy
- **Sorting Preprocessing**: When sorting helps
- **Hash Map Usage**: When to use for O(1) lookups

## Code Style Guidelines

### Go Code Examples
- Use proper Go naming conventions (camelCase for functions)
- Include necessary imports
- Add comments for complex logic
- Handle edge cases explicitly
- Include helper functions if needed

### Complexity Notation
- Use Big O notation consistently
- Specify what 'n' represents (array length, nodes count, etc.)
- Mention both time and space complexity
- Note any amortized complexity

## Quality Checklist
Before marking an explanation as complete:

- [ ] Problem statement is accurate
- [ ] Code compiles and passes LeetCode tests
- [ ] Time and space complexity are correct
- [ ] Step-by-step walkthrough is clear
- [ ] Edge cases are covered
- [ ] Related problems are listed
- [ ] Formatting follows this template

## Version Control
- Update `CATEGORIES.md` when adding new explanations
- Include problem number in commit message
- Reference template version if updated

---

*Template Version: 1.0.0*  
*Last Updated: 2026-01-31*