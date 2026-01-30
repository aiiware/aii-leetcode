#!/usr/bin/env python3
"""
Script to fix import statements in test files that use utils. functions
but don't import leetcode/utils.
"""

import os
import re
from pathlib import Path

def find_test_files_needing_imports():
    """Find all test files that use utils. but don't import leetcode/utils"""
    test_files = []
    
    # Find all test files
    for root, dirs, files in os.walk("."):
        # Skip hidden directories
        dirs[:] = [d for d in dirs if not d.startswith(".")]
        
        for file in files:
            if file.endswith("_test.go"):
                filepath = os.path.join(root, file)
                
                # Read file content
                with open(filepath, 'r') as f:
                    content = f.read()
                
                # Check if file uses utils. functions
                if re.search(r'\butils\.\w+', content):
                    # Check if file imports leetcode/utils
                    if 'import "leetcode/utils"' not in content:
                        test_files.append(filepath)
    
    return test_files

def add_import_to_file(filepath):
    """Add import "leetcode/utils" to a test file"""
    with open(filepath, 'r') as f:
        lines = f.readlines()
    
    # Find the package line
    package_line_index = -1
    for i, line in enumerate(lines):
        if line.strip().startswith("package "):
            package_line_index = i
            break
    
    if package_line_index == -1:
        print(f"  Warning: Could not find package declaration in {filepath}")
        return False
    
    # Check if there's already an import block
    import_start = -1
    import_end = -1
    in_import_block = False
    
    for i in range(package_line_index + 1, len(lines)):
        line = lines[i].strip()
        
        if line.startswith("import"):
            if line == "import" or line.startswith("import ("):
                in_import_block = True
                import_start = i
            else:
                # Single line import
                import_start = i
                import_end = i
                break
        elif in_import_block and line == ")":
            import_end = i
            break
        elif in_import_block and line and not line.startswith("//"):
            # Still in import block
            continue
        elif in_import_block and not line:
            # Empty line in import block
            continue
        elif in_import_block:
            # Unexpected content, end of import block
            import_end = i - 1
            break
    
    # Add the import
    if import_start == -1:
        # No imports yet, add after package line
        new_lines = lines[:package_line_index + 1]
        new_lines.append('\n')
        new_lines.append('import "leetcode/utils"\n')
        new_lines.extend(lines[package_line_index + 1:])
    elif import_end == import_start:
        # Single line import, convert to import block
        import_line = lines[import_start].strip()
        new_lines = lines[:import_start]
        new_lines.append('import (\n')
        new_lines.append(f'    {import_line[7:]}\n')  # Remove "import " prefix
        new_lines.append('    "leetcode/utils"\n')
        new_lines.append(')\n')
        new_lines.extend(lines[import_start + 1:])
    else:
        # Already have import block, add to it
        new_lines = lines[:import_end]
        new_lines.append('    "leetcode/utils"\n')
        new_lines.append(')\n')
        new_lines.extend(lines[import_end + 1:])
    
    # Write back to file
    with open(filepath, 'w') as f:
        f.writelines(new_lines)
    
    return True

def main():
    print("Finding test files that need import fixes...")
    test_files = find_test_files_needing_imports()
    
    print(f"\nFound {len(test_files)} test files needing import fixes:")
    for filepath in test_files:
        print(f"  - {filepath}")
    
    if not test_files:
        print("\nNo files need fixing!")
        return
    
    print(f"\nFixing imports in {len(test_files)} files...")
    fixed_count = 0
    for filepath in test_files:
        print(f"  Fixing {filepath}...")
        if add_import_to_file(filepath):
            fixed_count += 1
    
    print(f"\n✅ Fixed imports in {fixed_count} out of {len(test_files)} files")

if __name__ == "__main__":
    main()