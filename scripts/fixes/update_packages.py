#!/usr/bin/env python3
"""
Update package declarations and imports in categorized files.
"""

import os
import re
from pathlib import Path

def get_package_name(dir_name):
    """Convert directory name to valid Go package name."""
    # Handle special cases
    if dir_name == "linked-lists":
        return "linkedlists"
    return dir_name

def update_file_package(filepath, package_name):
    """Update package declaration in a Go file."""
    with open(filepath, 'r') as f:
        content = f.read()
    
    # Replace package declaration
    # Look for "package leetcode" at the beginning of the file
    lines = content.split('\n')
    updated = False
    
    for i, line in enumerate(lines):
        if line.strip().startswith('package leetcode'):
            lines[i] = f'package {package_name}'
            updated = True
            break
    
    if not updated:
        # Try to find package declaration anywhere
        content = re.sub(r'^package leetcode\b', f'package {package_name}', content, flags=re.MULTILINE)
        lines = content.split('\n')
    
    # Update imports from utils
    for i, line in enumerate(lines):
        if 'import' in line and 'utils' in line:
            # Check if it's importing from current directory
            if '"."' in line or '"./' in line:
                # Replace with proper utils import
                lines[i] = line.replace('"."', '"leetcode/utils"').replace('"./', '"leetcode/utils')
    
    new_content = '\n'.join(lines)
    
    # Write back if changed
    if new_content != content:
        with open(filepath, 'w') as f:
            f.write(new_content)
        return True
    
    return False

def update_imports_in_file(filepath, dir_name):
    """Update imports that reference other packages."""
    package_name = get_package_name(dir_name)
    
    with open(filepath, 'r') as f:
        content = f.read()
    
    # We'll handle imports more carefully in a separate step
    # For now, just update the package declaration
    return update_file_package(filepath, package_name)

def main():
    print("UPDATING PACKAGE DECLARATIONS AND IMPORTS")
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
        
        package_name = get_package_name(dir_name)
        print(f"\nProcessing {dir_name}/ -> package {package_name}")
        
        # Process all .go files in directory
        go_files = list(dir_path.glob('*.go'))
        dir_updated = 0
        
        for filepath in go_files:
            total_files += 1
            if update_imports_in_file(filepath, dir_name):
                dir_updated += 1
                total_updated += 1
        
        print(f"  Updated {dir_updated}/{len(go_files)} files")
    
    print(f"\nSUMMARY:")
    print(f"  Total files processed: {total_files}")
    print(f"  Files updated: {total_updated}")
    
    # Check for any remaining "package leetcode" declarations
    print("\nChecking for remaining 'package leetcode' declarations...")
    remaining = []
    for dir_name in directories:
        dir_path = Path(dir_name)
        for filepath in dir_path.glob('*.go'):
            with open(filepath, 'r') as f:
                if 'package leetcode' in f.read():
                    remaining.append(str(filepath))
    
    if remaining:
        print(f"Found {len(remaining)} files still with 'package leetcode':")
        for f in remaining[:10]:  # Show first 10
            print(f"  {f}")
        if len(remaining) > 10:
            print(f"  ... and {len(remaining) - 10} more")
    else:
        print("  All files updated successfully!")

if __name__ == '__main__':
    main()