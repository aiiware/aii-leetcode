#!/usr/bin/env python3
"""
Update imports in categorized files to use proper package paths.
"""

import os
import re
from pathlib import Path

def update_imports_in_file(filepath, dir_name):
    """Update imports in a Go file."""
    with open(filepath, 'r') as f:
        content = f.read()
    
    original = content
    
    # Pattern to find import blocks
    # Handle single import: import "package"
    # Handle multiple imports: import ( ... )
    
    # First, handle imports of current directory (.)
    # These should now import from utils
    content = re.sub(
        r'import\s+"\."',
        'import "leetcode/utils"',
        content
    )
    
    content = re.sub(
        r'import\s+\(\s*"\."\s*\)',
        'import (\n    "leetcode/utils"\n)',
        content
    )
    
    # In multi-import blocks, replace "." with "leetcode/utils"
    lines = content.split('\n')
    in_import_block = False
    import_block_start = -1
    
    for i, line in enumerate(lines):
        stripped = line.strip()
        
        # Check for import block start
        if stripped.startswith('import ('):
            in_import_block = True
            import_block_start = i
        elif in_import_block and stripped == ')':
            # End of import block
            in_import_block = False
        elif in_import_block and '"."' in line:
            # Replace "." import within block
            lines[i] = line.replace('"."', '"leetcode/utils"')
    
    content = '\n'.join(lines)
    
    # Write back if changed
    if content != original:
        with open(filepath, 'w') as f:
            f.write(content)
        return True
    
    return False

def check_test_imports():
    """Check that test files import the correct package."""
    print("\nCHECKING TEST FILE IMPORTS")
    print("=" * 60)
    
    issues = []
    
    for dir_name in Path('.').iterdir():
        if not dir_name.is_dir() or dir_name.name.startswith('.'):
            continue
        
        # Skip non-category directories
        if dir_name.name not in ['arrays', 'strings', 'dp', 'trees', 'math', 
                                'linked-lists', 'sql', 'sorting', 'graphs', 'design',
                                'utils', 'indexes', 'cmd', 'data_structures', 
                                'testutils', 'tutorials', 'temp_debug']:
            continue
        
        package_name = dir_name.name
        if package_name == 'linked-lists':
            package_name = 'linkedlists'
        
        # Check test files in this directory
        for test_file in dir_name.glob('*_test.go'):
            with open(test_file, 'r') as f:
                test_content = f.read()
            
            # Test file should import its own package
            expected_import = f'"{package_name}"'
            if expected_import not in test_content and f'package {package_name}' not in test_content:
                # Check what it's importing
                import_match = re.search(r'import\s+[^(]*?"([^"]+)"', test_content)
                if import_match:
                    actual_import = import_match.group(1)
                    issues.append(f"{test_file}: imports {actual_import}, should import {expected_import}")
                else:
                    issues.append(f"{test_file}: no clear import found")
    
    if issues:
        print(f"Found {len(issues)} import issues:")
        for issue in issues[:20]:  # Show first 20
            print(f"  {issue}")
        if len(issues) > 20:
            print(f"  ... and {len(issues) - 20} more")
    else:
        print("All test files have correct imports!")
    
    return len(issues)

def main():
    print("UPDATING IMPORTS IN CATEGORIZED FILES")
    print("=" * 60)
    
    # Directories to process
    directories = [
        'arrays', 'strings', 'dp', 'trees', 'math',
        'linked-lists', 'sql', 'sorting', 'graphs', 'design'
    ]
    
    total_updated = 0
    total_files = 0
    
    for dir_name in directories:
        dir_path = Path(dir_name)
        if not dir_path.exists():
            print(f"Directory not found: {dir_name}")
            continue
        
        print(f"\nProcessing {dir_name}/")
        
        # Process all .go files in directory
        go_files = list(dir_path.glob('*.go'))
        dir_updated = 0
        
        for filepath in go_files:
            total_files += 1
            if update_imports_in_file(filepath, dir_name):
                dir_updated += 1
                total_updated += 1
        
        print(f"  Updated imports in {dir_updated}/{len(go_files)} files")
    
    print(f"\nSUMMARY:")
    print(f"  Total files processed: {total_files}")
    print(f"  Files updated: {total_updated}")
    
    # Check test imports
    issue_count = check_test_imports()
    
    if issue_count > 0:
        print(f"\nNOTE: {issue_count} test files need manual import fixes.")
        print("Test files should import their own package, not 'leetcode' or '.'")

if __name__ == '__main__':
    main()