#!/usr/bin/env python3

"""
Script to identify missing LeetCode problems from the index.
"""

import os
import re
import glob

def get_existing_problems():
    """Find all existing problem files in the project"""
    existing_problems = set()
    
    categories = ['arrays', 'dp', 'graphs', 'trees', 'linked-lists', 'math', 'strings', 'design']
    
    for category in categories:
        if os.path.exists(category):
            # Find .go files that are actual solutions (not test files)
            go_files = glob.glob(f'{category}/0*.go')
            for file in go_files:
                base = os.path.basename(file)
                if not base.endswith('_test.go'):
                    # Extract problem number (first 4 digits)
                    problem_num = base[:4]
                    existing_problems.add(problem_num)
    
    return existing_problems

def get_index_problems():
    """Parse the index file to get all problems that should exist"""
    problems = set()
    
    try:
        with open('indexes/by_number.md', 'r') as f:
            content = f.read()
        
        # Extract problem numbers from the index
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

def main():
    existing = get_existing_problems()
    index_problems = get_index_problems()
    
    print(f"Total problems in index: {len(index_problems)}")
    print(f"Total implemented problems: {len(existing)}")
    print(f"Missing problems: {len(index_problems - existing)}")
    
    missing = index_problems - existing
    print("\nFirst 20 missing problems:")
    for problem in sorted(list(missing))[:20]:
        print(f"  {problem}")
        
    # Save list of missing problems
    with open('temp_missing_problems/missing_problems.txt', 'w') as f:
        for problem in sorted(list(missing)):
            f.write(f"{problem}\n")
    
    print(f"\nSaved {len(missing)} missing problems to temp_missing_problems/missing_problems.txt")

if __name__ == "__main__":
    main()