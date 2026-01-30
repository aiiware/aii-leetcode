# LeetCode Reorganization Backlog Specification

## Current Status Summary

**✅ COMPLETED:**
1. **File Reorganization** - All 189 LeetCode problems moved to categorized directories:
   - `arrays/` (58 problems)
   - `strings/` (28 problems)
   - `dp/` (28 problems)
   - `trees/` (21 problems)
   - `math/` (17 problems)
   - `linked-lists/` (14 problems)
   - `sql/` (9 problems)
   - `sorting/` (8 problems)
   - `graphs/` (3 problems)
   - `design/` (3 problems)

2. **Utils Package Created** - Consolidated shared types and functions into `utils/utils.go`

3. **Package Declarations Updated** - All files updated to use category-specific packages

4. **Demo Program Updated** - `cmd/demo/main.go` updated to use new package structure

**⚠️ PARTIALLY COMPLETED:**
1. **Import Fixes** - Basic import fixes applied, but some test files still have issues
2. **Test Execution** - Some packages pass tests, others have compilation errors

## 🔴 CRITICAL ISSUES (Blocking)

### 1. Cross-Package Test Dependencies
**Problem**: Test files in one category call functions from other categories
**Examples**:
- `arrays/0108_convert_sorted_array_to_binary_search_tree_test.go` calls:
  - `isValidBST` (from `trees/0098_validate_binary_search_tree.go`)
  - `isBalancedBottomUp` (from `trees/0110_balanced_binary_tree.go`)
  - `inorderTraversal` (from `trees/0094_binary_tree_inorder_traversal.go`)

**Impact**: Circular dependencies and compilation failures
**Priority**: HIGH
**Estimated Effort**: Medium (2-3 hours)

### 2. Unexported Helper Functions
**Problem**: Test files call unexported helper functions that were previously accessible when everything was in one package
**Examples**:
- `sortSubsets` function in `utils/utils.go` was lowercase (now fixed to `SortSubsets`)
- Other similar functions may exist

**Impact**: Compilation errors in test files
**Priority**: HIGH
**Estimated Effort**: Low (1 hour)

### 3. Missing Imports in Test Files
**Problem**: Test files use `utils.` functions but don't import `leetcode/utils`
**Status**: Partially fixed (45 files updated), but may need verification
**Priority**: MEDIUM
**Estimated Effort**: Low (30 minutes)

## 🟡 MEDIUM PRIORITY ISSUES

### 4. Command-Line Tools Broken
**Problem**: Files in `cmd/` directory still import `"leetcode"` package which no longer exists
**Files affected**:
- `cmd/leetcode-cli/main.go`
- `cmd/leetcode-demo/main.go`
- Other command-line tools

**Solution**: Update imports to use specific category packages or create facade
**Priority**: MEDIUM
**Estimated Effort**: Medium (1-2 hours)

### 5. Test Utilities Package
**Problem**: Test helper functions are scattered across test files
**Solution**: Create `testutils/` package for shared test helpers
**Priority**: MEDIUM
**Estimated Effort**: Medium (2 hours)

### 6. Index Files Incomplete
**Problem**: Planned index files (`indexes/by_category.md`, etc.) not created
**Solution**: Generate comprehensive index files
**Priority**: LOW
**Estimated Effort**: Medium (1-2 hours)

## 🟢 LOW PRIORITY / ENHANCEMENTS

### 7. README Update
**Problem**: README.md still shows old structure (150 problems)
**Solution**: Update with new structure and 189+ problems
**Priority**: LOW
**Estimated Effort**: Low (30 minutes)

### 8. Missing Problems
**Problem**: Gaps in problem sequence (0186, 0187 missing)
**Solution**: Implement missing problems
**Priority**: LOW
**Estimated Effort**: Low (1 hour)

### 9. Searchable Index
**Problem**: Planned `search.json` not created
**Solution**: Create JSON index for web interface
**Priority**: LOW
**Estimated Effort**: Medium (1 hour)

## 📋 RESTART PLAN

When restarting the reorganization process, follow this sequence:

### Phase 1: Foundation (1 hour)
1. **Verify current state** - Run `go test ./...` to see current failures
2. **Create test utilities** - Move cross-package test helpers to `testutils/`
3. **Fix critical imports** - Ensure all test files import needed packages

### Phase 2: Cross-Package Resolution (2-3 hours)
1. **Identify all cross-package calls** - Create mapping of dependencies
2. **Create facade/interface layer** - Design solution for cross-package access
3. **Implement solution** - Either:
   - Move commonly used functions to `utils/`
   - Create test helper package
   - Refactor tests to avoid cross-package calls

### Phase 3: Command Tools (1-2 hours)
1. **Update cmd/ imports** - Fix command-line tools
2. **Test all commands** - Verify they work with new structure

### Phase 4: Documentation (1-2 hours)
1. **Create index files** - Generate `indexes/` directory contents
2. **Update README** - Document new structure
3. **Create migration guide** - For future contributors

### Phase 5: Validation (1 hour)
1. **Run all tests** - Ensure `go test ./...` passes
2. **Test demo programs** - Verify they work
3. **Create final commit** - With comprehensive message

## 🎯 RECOMMENDED APPROACH FOR RESTART

### Option A: Incremental Fix (Recommended)
1. Start with one category (e.g., `strings/` which already passes tests)
2. Fix dependencies outward from there
3. Move category by category

### Option B: Architectural Redesign
1. Create proper package interfaces
2. Design clear boundaries between categories
3. Implement systematically

### Option C: Simplified Structure
1. Reduce number of categories (merge some)
2. Keep related problems together
3. Simplify package structure

**Recommended**: **Option A** - Incremental fix starting with working categories

## 📊 SUCCESS METRICS

- [ ] All tests pass: `go test ./...`
- [ ] Demo programs work: `go run cmd/demo/main.go`
- [ ] Command-line tools work
- [ ] Index files generated and accurate
- [ ] README updated with new structure
- [ ] No compilation warnings

## 🚨 RISK MITIGATION

1. **Git commits after each phase** - Easy rollback
2. **Start with working categories** - Build confidence
3. **Test frequently** - Catch issues early
4. **Document decisions** - For future reference

## 📝 NOTES FOR NEXT SESSION

1. **Current working packages**: `strings/`, `math/`, `linked-lists/`, `sql/`, `sorting/`, `graphs/`, `design/`
2. **Problematic packages**: `arrays/`, `dp/`, `trees/` (have cross-package dependencies)
3. **Key decision needed**: How to handle cross-package test dependencies
4. **Time estimate for completion**: 6-8 hours total

---

*Last Updated: $(date)*  
*Total Issues: 9 (3 Critical, 3 Medium, 3 Low)*  
*Estimated Completion Time: 6-8 hours*