#!/usr/bin/env python3

"""
Identify missing LeetCode problems from 1-1000
"""

import os
import re
import sys

def identify_missing_problems():
    # Get all implemented problems with their categories
    implemented = {}
    
    # Walk through the project directory
    for root, dirs, files in os.walk('.'):
        for file in files:
            if file.endswith('.go') and re.search(r'0\d{3}_', file):
                # Extract problem number
                problem_match = re.search(r'0(\d{3})_', file)
                if problem_match:
                    problem_num = problem_match.group(1)
                    # Extract directory name (category)
                    dir_name = os.path.basename(root)
                    implemented[problem_num] = dir_name
    
    # Create set of all problem numbers 1-1000
    all_problems = set(f'{i:03d}' for i in range(1, 1001))
    
    # Find missing problems
    missing = all_problems - set(implemented.keys())
    
    print(f"Total problems: 1000")
    print(f"Implemented problems: {len(implemented)}")
    print(f"Missing problems: {len(missing)}")
    print(f"Highest implemented problem: {max(implemented.keys()) if implemented else 'None'}")
    
    if missing:
        print("\nFirst 20 missing problems:")
        for problem in sorted(missing)[:20]:
            print(f"  {problem}")
        
        print(f"\nTotal missing: {len(missing)} problems")
    else:
        print("All problems implemented!")

if __name__ == "__main__":
    identify_missing_problems()