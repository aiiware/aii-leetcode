#!/usr/bin/env python3
"""
Move LeetCode problem files to categorized directories based on CSV mapping.
"""

import csv
import os
import shutil
from pathlib import Path

def read_categorization():
    """Read categorization from CSV file."""
    with open('problem_categories_corrected.csv', 'r') as f:
        reader = csv.DictReader(f)
        return list(reader)

def move_files(problems):
    """Move problem files to their categorized directories."""
    moved = []
    errors = []
    
    for p in problems:
        filename = p['filename']
        category = p['category']
        
        # Source paths
        src_impl = filename
        src_test = filename.replace('.go', '_test.go')
        
        # Destination directory
        dest_dir = category
        
        # Check if source files exist
        if not os.path.exists(src_impl):
            errors.append(f"Source file not found: {src_impl}")
            continue
            
        # Create destination directory if it doesn't exist
        os.makedirs(dest_dir, exist_ok=True)
        
        # Destination paths
        dest_impl = os.path.join(dest_dir, filename)
        dest_test = os.path.join(dest_dir, src_test)
        
        try:
            # Move implementation file
            shutil.move(src_impl, dest_impl)
            moved.append((src_impl, dest_impl))
            
            # Move test file if it exists
            if os.path.exists(src_test):
                shutil.move(src_test, dest_test)
                moved.append((src_test, dest_test))
            else:
                print(f"  Note: Test file not found: {src_test}")
                
            print(f"  Moved {filename} -> {dest_dir}/")
            
        except Exception as e:
            errors.append(f"Error moving {filename}: {e}")
    
    return moved, errors

def update_package_declarations():
    """Update package declarations in moved files."""
    # This will be done in a separate step
    print("\nPackage declarations will be updated in Phase 5")
    return []

def main():
    print("MOVING FILES TO CATEGORIZED DIRECTORIES")
    print("=" * 60)
    
    # Read categorization
    problems = read_categorization()
    print(f"Found {len(problems)} problems to categorize")
    
    # Group by category for summary
    from collections import Counter
    categories = Counter(p['category'] for p in problems)
    print("\nCategory distribution:")
    for cat, count in sorted(categories.items(), key=lambda x: x[1], reverse=True):
        print(f"  {cat:15}: {count:3} files")
    
    # Ask for confirmation
    print("\nThis will move files from root directory to categorized directories.")
    response = input("Continue? (y/n): ").strip().lower()
    
    if response != 'y':
        print("Aborted.")
        return
    
    # Move files
    print("\nMoving files...")
    moved, errors = move_files(problems)
    
    # Print summary
    print(f"\nSUMMARY:")
    print(f"  Files moved: {len(moved)}")
    print(f"  Errors: {len(errors)}")
    
    if errors:
        print("\nERRORS:")
        for error in errors:
            print(f"  {error}")
    
    # Print what's left in root
    print("\nFiles remaining in root directory:")
    remaining = list(Path('.').glob('[0-9][0-9][0-9][0-9]_*.go'))
    for f in sorted(remaining):
        print(f"  {f.name}")
    
    if not remaining:
        print("  (none)")

if __name__ == '__main__':
    main()