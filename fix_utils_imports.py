#!/usr/bin/env python3
"""
Fix imports in all Go files to use utils package.
"""

import os
import re
from pathlib import Path

# Common functions that moved to utils
UTILS_FUNCTIONS = {
    'SlicesEqual', 'IntsEqual', 'StringsEqual', 'StringSlicesEqual',
    'BoolsEqual', 'MatrixEqual', 'MakeRange', 'Repeat', 'IntPtr',
    'IntsToList', 'CopyList', 'PrintList', 'SubsetsEqual',
    'HasDuplicateSubsets', 'IsSubset', 'sortSubsets', 'subsetsEqual',
    'hasDuplicateSubsets', 'isSubset',
    'CountBits', 'IsValidGrayCode', 'IsPermutation',
    'Min', 'Max', 'Abs', 'MinFloat64', 'MaxFloat64', 'MinOf', 'MaxOf',
    'maxInt', 'minInt'  # Common helper functions
}

# Types that moved to utils
UTILS_TYPES = {
    'ListNode', 'TreeNode'
}

# Helper functions that might be used
UTILS_HELPERS = {
    'CloneTree', 'CreateCompleteTree', 'CreateRightSkewedTree',
    'CreateLeftSkewedTree', 'CreateSymmetricTree'
}

def needs_utils_import(content):
    """Check if file needs utils import."""
    for func in UTILS_FUNCTIONS:
        if f' {func}(' in content or f' {func} ' in content or f'\t{func}(' in content:
            return True
    
    for typ in UTILS_TYPES:
        if f' {typ}' in content or f'*{typ}' in content or f'[]{typ}' in content:
            return True
    
    for helper in UTILS_HELPERS:
        if f' {helper}(' in content or f' {helper} ' in content:
            return True
    
    return False

def add_utils_import(filepath):
    """Add utils import to Go file if needed."""
    with open(filepath, 'r') as f:
        content = f.read()
    
    original = content
    
    # Check if already imports utils
    if '"leetcode/utils"' in content or 'utils.' in content:
        return False
    
    # Check if needs utils
    if not needs_utils_import(content):
        return False
    
    # Add import
    lines = content.split('\n')
    
    # Find import section
    import_found = False
    import_block_start = -1
    import_block_end = -1
    
    for i, line in enumerate(lines):
        stripped = line.strip()
        
        if stripped.startswith('import'):
            import_found = True
            if '(' in stripped:
                # Multi-line import block
                import_block_start = i
                # Find closing )
                for j in range(i, len(lines)):
                    if ')' in lines[j]:
                        import_block_end = j
                        break
            else:
                # Single import
                # Convert to multi-line import block
                lines[i] = 'import ('
                lines.insert(i + 1, '    ' + line.replace('import', '').strip())
                lines.insert(i + 2, ')')
                # Now add utils import
                lines.insert(i + 1, '    "leetcode/utils"')
                new_content = '\n'.join(lines)
                with open(filepath, 'w') as f:
                    f.write(new_content)
                return True
    
    if not import_found:
        # No import found, add after package declaration
        for i, line in enumerate(lines):
            if line.strip().startswith('package'):
                lines.insert(i + 1, '')
                lines.insert(i + 2, 'import "leetcode/utils"')
                new_content = '\n'.join(lines)
                with open(filepath, 'w') as f:
                    f.write(new_content)
                return True
    
    # In multi-line import block, add utils import
    if import_block_start != -1 and import_block_end != -1:
        # Check if utils already imported
        for i in range(import_block_start, import_block_end + 1):
            if '"leetcode/utils"' in lines[i]:
                return False
        
        # Add utils import before closing )
        lines.insert(import_block_end, '    "leetcode/utils"')
        new_content = '\n'.join(lines)
        with open(filepath, 'w') as f:
            f.write(new_content)
        return True
    
    return False

def update_function_calls(filepath):
    """Update function calls to use utils package if needed."""
    with open(filepath, 'r') as f:
        content = f.read()
    
    original = content
    
    # Replace standalone function calls
    for func in UTILS_FUNCTIONS:
        # Pattern: function call not preceded by a dot (package)
        pattern = r'(?<![a-zA-Z0-9_.])\b' + re.escape(func) + r'\b(?=\()'
        content = re.sub(pattern, f'utils.{func}', content)
    
    # Replace type references
    for typ in UTILS_TYPES:
        # Pattern: type not preceded by a dot
        pattern = r'(?<![a-zA-Z0-9_.])\b' + re.escape(typ) + r'\b'
        content = re.sub(pattern, f'utils.{typ}', content)
    
    # Replace helper function calls
    for helper in UTILS_HELPERS:
        pattern = r'(?<![a-zA-Z0-9_.])\b' + re.escape(helper) + r'\b(?=\()'
        content = re.sub(pattern, f'utils.{helper}', content)
    
    if content != original:
        with open(filepath, 'w') as f:
            f.write(content)
        return True
    
    return False

def process_directory(dir_path):
    """Process all Go files in directory."""
    updated = 0
    total = 0
    
    for filepath in dir_path.glob('*.go'):
        total += 1
        file_updated = False
        
        if add_utils_import(filepath):
            file_updated = True
        
        if update_function_calls(filepath):
            file_updated = True
        
        if file_updated:
            updated += 1
    
    return total, updated

def main():
    print("FIXING IMPORTS TO USE UTILS PACKAGE")
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
        total, updated = process_directory(dir_path)
        total_files += total
        total_updated += updated
        
        print(f"  Updated {updated}/{total} files")
    
    print(f"\nSUMMARY:")
    print(f"  Total files processed: {total_files}")
    print(f"  Files updated: {total_updated}")
    
    # Also process cmd/ directory
    print("\nProcessing cmd/ directory...")
    cmd_updated = 0
    cmd_total = 0
    
    for cmd_dir in Path('cmd').iterdir():
        if cmd_dir.is_dir():
            total, updated = process_directory(cmd_dir)
            cmd_total += total
            cmd_updated += updated
    
    print(f"  Updated {cmd_updated}/{cmd_total} files in cmd/")

if __name__ == '__main__':
    main()