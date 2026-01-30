#!/usr/bin/env python3
"""
Script to fix test files that have:
1. Duplicate imports of "leetcode/utils"
2. Function calls that should use uppercase names
"""

import os
import re

def fix_test_file(filepath):
    """Fix a test file with import and function call issues"""
    with open(filepath, 'r') as f:
        content = f.read()
    
    # Fix 1: Remove duplicate imports
    lines = content.split('\n')
    fixed_lines = []
    utils_import_count = 0
    
    for line in lines:
        stripped = line.strip()
        if '"leetcode/utils"' in stripped:
            utils_import_count += 1
            if utils_import_count == 1:
                fixed_lines.append(line)  # Keep first occurrence
            # Skip duplicates
        else:
            fixed_lines.append(line)
    
    content = '\n'.join(fixed_lines)
    
    # Fix 2: Replace function calls with correct case
    # These are the functions in utils package
    function_mappings = {
        r'\butils\.sortSubsets\b': 'utils.SortSubsets',  # Now uppercase
        r'\butils\.subsetsEqual\b': 'utils.SubsetsEqual',
        r'\butils\.hasDuplicateSubsets\b': 'utils.HasDuplicateSubsets',
        r'\butils\.isSubset\b': 'utils.IsSubset',
    }
    
    for pattern, replacement in function_mappings.items():
        content = re.sub(pattern, replacement, content)
    
    # Write back to file
    with open(filepath, 'w') as f:
        f.write(content)
    
    return True

def main():
    # Find all test files that might have these issues
    test_files = []
    
    for root, dirs, files in os.walk("."):
        # Skip hidden directories
        dirs[:] = [d for d in dirs if not d.startswith(".")]
        
        for file in files:
            if file.endswith("_test.go"):
                filepath = os.path.join(root, file)
                test_files.append(filepath)
    
    print(f"Checking {len(test_files)} test files for issues...")
    
    fixed_count = 0
    for filepath in test_files:
        with open(filepath, 'r') as f:
            content = f.read()
        
        # Check if this file needs fixing
        needs_fix = False
        
        # Check for duplicate imports
        if content.count('"leetcode/utils"') > 1:
            needs_fix = True
        
        # Check for lowercase function calls
        if (re.search(r'\butils\.sortSubsets\b', content) or
            re.search(r'\butils\.subsetsEqual\b', content) or
            re.search(r'\butils\.hasDuplicateSubsets\b', content) or
            re.search(r'\butils\.isSubset\b', content)):
            needs_fix = True
        
        if needs_fix:
            print(f"  Fixing {filepath}...")
            if fix_test_file(filepath):
                fixed_count += 1
    
    print(f"\n✅ Fixed {fixed_count} test files")

if __name__ == "__main__":
    main()