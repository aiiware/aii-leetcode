#!/usr/bin/env python3

import os
import re

def analyze_missing_problems():
    # Read index file
    with open('indexes/by_number.md', 'r') as f:
        content = f.read()

    # Extract all problem numbers from the index
    problem_pattern = r'- \*\*(\d{4})\*\*'
    found_problems = re.findall(problem_pattern, content)

    # Get all existing problem files
    existing_files = []
    for root, dirs, files in os.walk('.'):
        if root.startswith('./scripts') or root.startswith('./cmd') or root.startswith('./temp_') or root.startswith('./docs') or root.startswith('./data') or root.startswith('./testutils'):
            continue
        for file in files:
            if file.endswith('.go') and not file.endswith('_test.go'):
                # Extract problem number from filename
                match = re.search(r'(\d{4})_', file)
                if match:
                    existing_files.append(match.group(1))

    existing_files = list(set(existing_files))  # Remove duplicates

    # Find missing problems
    missing_problems = []
    for problem in found_problems:
        if problem not in existing_files:
            missing_problems.append(problem)

    print(f'Found {len(found_problems)} total problems in index')
    print(f'Found {len(existing_files)} existing problem files')
    print(f'Missing {len(missing_problems)} problems')
    
    # Show problem breakdown by range
    ranges = {}
    for prob in missing_problems:
        range_key = f"{(int(prob) // 100) * 100:04d}-{(int(prob) // 100 + 1) * 100 - 1:04d}"
        if range_key not in ranges:
            ranges[range_key] = []
        ranges[range_key].append(prob)
    
    print("\nMissing problems by range:")
    for range_key in sorted(ranges.keys()):
        print(f"  {range_key}: {len(ranges[range_key])} problems")
        
    print("\nFirst 30 missing problems:")
    for i, prob in enumerate(missing_problems[:30]):
        print(f'  {prob}')
        
    return missing_problems

if __name__ == "__main__":
    analyze_missing_problems()