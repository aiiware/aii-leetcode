# LeetCode Solutions - Project Backlog

## 📋 Overview

This backlog tracks all current and future work for the LeetCode solutions repository. It includes both reorganization tasks and general project improvements.

## 🎯 Priority Legend

- **P0 (Critical)**: Blocking issues that prevent basic functionality
- **P1 (High)**: Important improvements that add significant value
- **P2 (Medium)**: Nice-to-have enhancements
- **P3 (Low)**: Future ideas and optimizations

## 🚨 Current Sprint (P0 - Critical)

### 1. Fix Cross-Package Test Dependencies
**Status**: 🔴 Blocking
**Description**: Test files in one category call functions from other categories, causing compilation failures
**Examples**:
- `arrays/0108_convert_sorted_array_to_binary_search_tree_test.go` calls tree validation functions
- Similar issues in `dp/` and `trees/` packages

**Acceptance Criteria**:
- [ ] All tests pass: `go test ./...`
- [ ] No circular dependencies between packages
- [ ] Test files only call functions from their own package or shared utilities

**Estimated Effort**: 2-3 hours

### 2. Fix Command-Line Tools
**Status**: 🔴 Blocking
**Description**: Command-line tools in `cmd/` directory import `"leetcode"` package which no longer exists
**Files affected**:
- `cmd/leetcode-cli/main.go`
- `cmd/leetcode-demo/main.go`
- Other command-line tools

**Acceptance Criteria**:
- [ ] All command-line tools compile without errors
- [ ] Tools work with new package structure
- [ ] Demo programs run successfully

**Estimated Effort**: 1-2 hours

### 3. Complete Import Fixes
**Status**: 🔴 Blocking
**Description**: Some test files still have missing or incorrect imports
**Current Status**: 45 files fixed, but need verification

**Acceptance Criteria**:
- [ ] All files have correct import statements
- [ ] No "undefined" compilation errors
- [ ] All packages can be imported correctly

**Estimated Effort**: 1 hour

## 📋 Next Sprint (P1 - High Priority)

### 4. Create Test Utilities Package
**Status**: 🟡 Important
**Description**: Test helper functions are scattered across test files
**Solution**: Create `testutils/` package for shared test helpers

**Acceptance Criteria**:
- [ ] `testutils/` package created with common test helpers
- [ ] Cross-package test dependencies moved to `testutils/`
- [ ] Test files updated to use `testutils/`

**Estimated Effort**: 2 hours

### 5. Generate Index Files
**Status**: 🟡 Important
**Description**: Planned index files (`indexes/by_category.md`, etc.) not created
**Files to create**:
- `indexes/by_category.md` - Problems grouped by category
- `indexes/by_difficulty.md` - Problems grouped by difficulty
- `indexes/by_number.md` - All problems in numerical order
- `indexes/search.json` - Searchable index (optional)

**Acceptance Criteria**:
- [ ] All index files created and accurate
- [ ] Indexes include all 189+ problems
- [ ] Indexes are properly formatted and linked

**Estimated Effort**: 1-2 hours

### 6. Update README Documentation
**Status**: 🟡 Important
**Description**: README.md still shows old structure (150 problems)
**Solution**: Update with new structure and 189+ problems

**Acceptance Criteria**:
- [ ] README.md updated with new categorized structure
- [ ] Problem count updated to 189+
- [ ] Documentation reflects current organization

**Estimated Effort**: 30 minutes

## 📚 Backlog (P2 - Medium Priority)

### 7. Implement Missing Problems
**Status**: 🟢 Medium
**Description**: Gaps in problem sequence (0186, 0187 missing)
**Solution**: Implement missing problems in appropriate categories

**Acceptance Criteria**:
- [ ] Problems 0186 and 0187 implemented
- [ ] Tests created for new problems
- [ ] Problems added to appropriate categories

**Estimated Effort**: 1 hour

### 8. Add Problem Difficulty Tags
**Status**: 🟢 Medium
**Description**: Problems don't have difficulty tags (Easy/Medium/Hard)
**Solution**: Add difficulty metadata to all problems

**Acceptance Criteria**:
- [ ] All problems tagged with difficulty
- [ ] Difficulty included in index files
- [ ] README updated with difficulty breakdown

**Estimated Effort**: 1 hour

### 9. Create Solution Explanations
**Status**: 🟢 Medium
**Description**: Add detailed explanations for complex solutions
**Solution**: Create `explanations/` directory with solution walkthroughs

**Acceptance Criteria**:
- [ ] Explanations for top 20 most complex problems
- [ ] Clear algorithm explanations with diagrams
- [ ] Time/space complexity analysis

**Estimated Effort**: 3-4 hours

## 💡 Future Ideas (P3 - Low Priority)

### 10. Create Searchable Web Interface
**Status**: 🔵 Future
**Description**: Web interface for browsing and searching problems
**Solution**: Simple web app with search, filtering, and code viewing

**Acceptance Criteria**:
- [ ] Basic web interface created
- [ ] Search functionality
- [ ] Problem filtering by category/difficulty

**Estimated Effort**: 8-10 hours

### 11. Add Performance Benchmarks
**Status**: 🔵 Future
**Description**: Add comprehensive performance benchmarks
**Solution**: Create benchmark suite for comparing solutions

**Acceptance Criteria**:
- [ ] Benchmarks for all problems
- [ ] Performance comparison charts
- [ ] Memory usage tracking

**Estimated Effort**: 4-5 hours

### 12. Create Learning Paths
**Status**: 🔵 Future
**Description**: Curated learning paths for different skill levels
**Solution**: Create guided study plans (beginner, intermediate, advanced)

**Acceptance Criteria**:
- [ ] 3 learning paths created
- [ ] Progressive difficulty progression
- [ ] Topic-based organization

**Estimated Effort**: 2-3 hours

## 📊 Progress Tracking

### Current Status (Jan 2026)
- **Total Problems**: 189 implemented
- **Categories**: 10 algorithm categories created
- **Tests Passing**: Partial (some packages work, others have issues)
- **Reorganization**: 90% complete

### Success Metrics
- [ ] All tests pass: `go test ./...`
- [ ] Demo programs work: `go run cmd/demo/main.go`
- [ ] Command-line tools work
- [ ] Index files generated and accurate
- [ ] README updated with new structure
- [ ] No compilation warnings

## 🔄 How to Use This Backlog

1. **Prioritize P0 items first** - Fix blocking issues before adding features
2. **Update status when working on items** - Mark as In Progress, Done, etc.
3. **Add new items as needed** - Use the priority system
4. **Reference specific items** - When discussing work, reference item numbers

## 📝 Notes

- **Reorganization details**: See `REORGANIZATION_STATUS.md` for detailed reorganization status
- **Project documentation**: See `AGENTS.md` for project overview and conventions
- **Working packages**: `strings/`, `math/`, `linked-lists/`, `sql/`, `sorting/`, `graphs/`, `design/`
- **Problematic packages**: `arrays/`, `dp/`, `trees/` (cross-package dependencies)

---
*Last Updated: $(date)*  
*Total Items: 12 (3 P0, 3 P1, 3 P2, 3 P3)*  
*Estimated Total Effort: 20-30 hours*