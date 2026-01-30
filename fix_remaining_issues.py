#!/usr/bin/env python3
"""
Fix remaining import and function call issues.
"""

import os
import re
from pathlib import Path

def fix_max_int(filepath):
    """Fix maxInt -> Max and minInt -> Min."""
    with open(filepath, 'r') as f:
        content = f.read()
    
    original = content
    
    # Fix maxInt -> Max
    content = re.sub(r'\bmaxInt\b', 'Max', content)
    # Fix minInt -> Min  
    content = re.sub(r'\bminInt\b', 'Min', content)
    
    if content != original:
        with open(filepath, 'w') as f:
            f.write(content)
        return True
    return False

def fix_unused_imports(filepath):
    """Remove unused utils imports."""
    with open(filepath, 'r') as f:
        content = f.read()
    
    original = content
    
    # Check if utils is imported but not used
    if '"leetcode/utils"' in content:
        # Simple check: if "utils." not in content, remove import
        if 'utils.' not in content:
            # Remove the import line
            lines = content.split('\n')
            new_lines = []
            i = 0
            while i < len(lines):
                line = lines[i]
                if '"leetcode/utils"' in line:
                    # Skip this line
                    i += 1
                    # If this was a single-line import, we're done
                    if 'import' in line and '(' not in line:
                        continue
                    # If it's in a multi-line import block, we need to check structure
                else:
                    new_lines.append(line)
                    i += 1
            
            new_content = '\n'.join(new_lines)
            if new_content != original:
                with open(filepath, 'w') as f:
                    f.write(new_content)
                return True
    
    return False

def fix_missing_utils_prefix(filepath):
    """Add utils. prefix to functions that need it."""
    with open(filepath, 'r') as f:
        content = f.read()
    
    original = content
    
    # Functions that should have utils. prefix
    functions = ['NewListFromSlice', 'NewTreeFromSlice', 'IntPtr']
    
    for func in functions:
        # Pattern: function call not preceded by utils.
        pattern = r'(?<!utils\.)\b' + re.escape(func) + r'\b(?=\()'
        content = re.sub(pattern, f'utils.{func}', content)
    
    if content != original:
        with open(filepath, 'w') as f:
            f.write(content)
        return True
    
    return False

def process_file(filepath):
    """Process a single file."""
    updated = False
    
    if fix_max_int(filepath):
        updated = True
    
    if fix_unused_imports(filepath):
        updated = True
    
    if fix_missing_utils_prefix(filepath):
        updated = True
    
    return updated

def main():
    print("FIXING REMAINING ISSUES")
    print("=" * 60)
    
    # Directories to process
    directories = [
        'arrays', 'strings', 'dp', 'trees', 'math',
        'linked-lists', 'sql', 'sorting', 'graphs', 'design'
    ]
    
    total_files = 0
    total_updated = 0
    
    for dir_name in directories:
        dir_path = Path(dir_name)
        if not dir_path.exists():
            print(f"Directory not found: {dir_name}")
            continue
        
        print(f"\nProcessing {dir_name}/")
        dir_updated = 0
        dir_total = 0
        
        for filepath in dir_path.glob('*.go'):
            dir_total += 1
            if process_file(filepath):
                dir_updated += 1
        
        total_files += dir_total
        total_updated += dir_updated
        
        print(f"  Updated {dir_updated}/{dir_total} files")
    
    print(f"\nSUMMARY:")
    print(f"  Total files processed: {total_files}")
    print(f"  Files updated: {total_updated}")

if __name__ == '__main__':
    main()