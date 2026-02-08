#!/usr/bin/env python3

"""
Complete script to identify ALL missing LeetCode problems from 1-1000.
"""

import os
import re
import glob

def get_existing_problems():
    """Find all existing problem files in the project"""
    existing_problems = set()
    
    # Walk through all directories
    for root, dirs, files in os.walk('.'):
        for file in files:
            if file.endswith('.go'):
                # Extract problem number (first 4 digits)
                problem_match = re.search(r'0(\d{3})_', file)
                if problem_match:
                    problem_num = problem_match.group(1)
                    existing_problems.add(problem_num)
    
    return existing_problems

def get_index_problems():
    """Parse the index file to get all problems that should exist"""
    problems = set()
    
    try:
        with open('indexes/by_number.md', 'r') as f:
            content = f.read()
        
        # Extract problem numbers from the index (pattern: **0001**: Two Sum)
        lines = content.split('\n')
        for line in lines:
            # Look for problem lines like "- **0001**: Two Sum (Easy) — *Arrays*"
            match = re.search(r'\*\*(\d{4})\*\*:', line)
            if match:
                problems.add(match.group(1))
    
    except FileNotFoundError:
        print("Index file not found")
        return set()
    
    return problems

def get_all_problems_in_range():
    """Get all problems from 0001 to 1000 (1000 problems total)"""
    return set(f'{i:04d}' for i in range(1, 1001))

def main():
    print("Analyzing problem coverage...")
    
    existing = get_existing_problems()
    index_problems = get_index_problems()
    all_problems = get_all_problems_in_range()
    
    print(f"Total problems in index: {len(index_problems)}")
    print(f"Total implemented problems: {len(existing)}")
    
    # The index has only 198 problems, but we're told 1000 problems total
    # Let's check what range of problems exist in the repository
    print(f"Range of existing problems: {min(existing) if existing else 'None'} to {max(existing) if existing else 'None'}")
    
    # Check for problems that exist but aren't in the index
    existing_not_in_index = existing - index_problems
    print(f"Problems implemented but not in index: {len(existing_not_in_index)}")
    
    # Find missing problems according to index 
    missing_in_index = index_problems - existing
    print(f"Missing problems according to index: {len(missing_in_index)}")
    
    if missing_in_index:
        print("\nFirst 10 missing problems:")
        for problem in sorted(list(missing_in_index))[:10]:
            print(f"  {problem}")
        
        # Save list of missing problems
        with open('temp_missing_problems/missing_problems.txt', 'w') as f:
            for problem in sorted(list(missing_in_index)):
                f.write(f"{problem}\n")
        
        print(f"\nSaved {len(missing_in_index)} missing problems to temp_missing_problems/missing_problems.txt")
    else:
        print("No missing problems according to index")
        
    # For the full 1000 problem scope, let's also see if there are additional problems we should consider
    print("\nAnalyzing full 1000 problem range...")
    full_range = set(f'{i:04d}' for i in range(1, 1001))
    missing_full = full_range - existing
    print(f"Missing from full 1000 range: {len(missing_full)}")
    
    if missing_full:
        print("First 10 missing from full range:")
        for problem in sorted(list(missing_full))[:10]:
            print(f"  {problem}")

if __name__ == "__main__":
    main()