#!/usr/bin/env python3
"""
Fix invalid 'func utils.' syntax in Go files.
"""

import os
import re
from pathlib import Path

def fix_func_utils_syntax(filepath):
    """Fix 'func utils.FunctionName' to 'func FunctionName'."""
    with open(filepath, 'r') as f:
        content = f.read()
    
    original = content
    
    # Fix: func utils.FunctionName -> func FunctionName
    # This regex looks for 'func' followed by whitespace, then 'utils.', then function name
    content = re.sub(r'\bfunc\s+utils\.(\w+)', r'func \1', content)
    
    if content != original:
        with open(filepath, 'w') as f:
            f.write(content)
        return True
    return False

def main():
    print("FIXING INVALID 'func utils.' SYNTAX")
    print("=" * 60)
    
    # Directories to process
    directories = [
        'arrays', 'strings', 'dp', 'trees', 'math',
        'linked-lists', 'sql', 'sorting', 'graphs', 'design'
    ]
    
    total_fixed = 0
    total_checked = 0
    
    for dir_name in directories:
        dir_path = Path(dir_name)
        if not dir_path.exists():
            print(f"Directory not found: {dir_name}")
            continue
        
        print(f"\nChecking {dir_name}/")
        dir_fixed = 0
        dir_checked = 0
        
        for filepath in dir_path.glob('*.go'):
            dir_checked += 1
            if fix_func_utils_syntax(filepath):
                dir_fixed += 1
                print(f"  Fixed: {filepath.name}")
        
        total_checked += dir_checked
        total_fixed += dir_fixed
        
        print(f"  Fixed {dir_fixed}/{dir_checked} files")
    
    print(f"\nSUMMARY:")
    print(f"  Total files checked: {total_checked}")
    print(f"  Files fixed: {total_fixed}")

if __name__ == '__main__':
    main()