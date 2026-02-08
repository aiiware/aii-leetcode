#!/usr/bin/env python3

import os
import re
import json
from collections import defaultdict

def analyze_current_repository():
    """
    Analyze the current repository state to understand what's implemented vs what's missing
    """
    print("Analyzing current repository state...")
    
    # Dictionary to track problems by category and number
    problems_by_category = defaultdict(list)
    total_files = 0
    
    # Walk through all directories except scripts and system directories
    for root, dirs, files in os.walk('.'):
        if any(root.startswith(prefix) for prefix in ['./scripts', './cmd', './temp_', './docs', './data', './testutils', './advanced-algorithms', './data_structures']):
            continue
            
        for file in files:
            if file.endswith('.go') and not file.endswith('_test.go'):
                total_files += 1
                
                # Extract problem number and category
                parts = root.split('/')
                if len(parts) > 1:
                    category = parts[-1] 
                    match = re.search(r'(\d{4})_', file)
                    if match:
                        problem_number = match.group(1)
                        problems_by_category[category].append(problem_number)
    
    # Print summary
    print(f"Total problem files found: {total_files}")
    print("\nProblems by category:")
    for category in sorted(problems_by_category.keys()):
        count = len(problems_by_category[category])
        print(f"  {category}: {count} problems")
    
    # Identify missing problems by looking at the index
    try:
        with open('indexes/by_number.md', 'r') as f:
            content = f.read()
            
        # Extract problem numbers from the index
        problem_pattern = r'- \*\*(\d{4})\*\*'
        found_problems = re.findall(problem_pattern, content)
        found_problems = set(found_problems)  # Make unique
        
        # Get all implemented problems
        implemented_problems = set()
        for category, problems in problems_by_category.items():
            implemented_problems.update(problems)
            
        missing_problems = found_problems - implemented_problems
        print(f"\nFrom index: {len(found_problems)} total problems")
        print(f"Implemented: {len(implemented_problems)} problems")
        print(f"Missing: {len(missing_problems)} problems")
        
        if missing_problems:
            print("\nFirst 20 missing problems:")
            for i, prob in enumerate(sorted(missing_problems)[:20]):
                print(f"  {prob}")
        
        # Create a detailed breakdown
        breakdown = {}
        for problem in sorted(missing_problems):
            range_key = f"{(int(problem) // 100) * 100:04d}-{(int(problem) // 100 + 1) * 100 - 1:04d}"
            if range_key not in breakdown:
                breakdown[range_key] = []
            breakdown[range_key].append(problem)
        
        print("\nMissing problems by range:")
        for range_key in sorted(breakdown.keys()):
            print(f"  {range_key}: {len(breakdown[range_key])} problems")
            
    except Exception as e:
        print(f"Could not read index file: {e}")
        print("Creating basic problem count analysis...")
        
        # Simple analysis based on number of files
        all_numbers = set()
        for category, problems in problems_by_category.items():
            all_numbers.update(problems)
        
        print(f"Total implemented problems: {len(all_numbers)}")

if __name__ == "__main__":
    analyze_current_repository()