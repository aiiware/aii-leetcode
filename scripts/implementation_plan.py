#!/usr/bin/env python3

import os
import re

def get_existing_problems_in_dir(directory):
    """Get list of existing problem files in a directory"""
    problems = []
    if os.path.exists(directory):
        for filename in os.listdir(directory):
            if filename.endswith('.go') and '_' in filename:
                # Extract problem number from filename like "0001_two_sum.go"
                match = re.match(r'^(\d+)_', filename)
                if match:
                    problem_number = match.group(1)
                    problems.append(problem_number)
    return set(problems)

def main():
    print("=== Current Project Status Analysis ===")
    
    # Get all existing problems by category
    categories = ['arrays', 'dp', 'graphs', 'trees', 'linked-lists', 
                   'math', 'strings', 'design', 'sorting']
    
    existing_problems = {}
    for cat in categories:
        existing_problems[cat] = get_existing_problems_in_dir(cat)
        
    # Show current status
    print("Current Implementation Status:")
    total_existing = set()
    for cat, problems in existing_problems.items():
        print(f"  {cat.capitalize()}: {len(problems)} problems")
        total_existing.update(problems)
    
    print(f"Total implemented: {len(total_existing)} problems")
    
    # Show what we should implement next
    print("\n=== Next Steps ===")
    print("1. Implement missing problems following the established patterns")
    print("2. Create comprehensive test cases")
    print("3. Add benchmarks for performance testing")
    print("4. Update documentation")
    
    print("\n=== Recommended Next Problems to Implement ===")
    print("These are high-value problems that are missing from the current implementation:")
    
    # Show a few missing problems from each category
    missing_recommendations = [
        ('arrays', ['0001', '0002', '0003', '0004', '0008', '0009', '0010']),
        ('dp', ['0001', '0002', '0004', '0005', '0006', '0008', '0009']),
        ('graphs', ['0001', '0002', '0003', '0004', '0005', '0006', '0007']),
        ('trees', ['0001', '0002', '0003', '0004', '0005', '0006', '0007']),
        ('linked-lists', ['0001', '0002', '0003', '0004', '0005', '0006', '0007']),
        ('math', ['0001', '0002', '0003', '0004', '0005', '0006', '0008']),
        ('strings', ['0001', '0002', '0003', '0004', '0005', '0006', '0007']),
        ('design', ['0001', '0002', '0003', '0004', '0005', '0006', '0007']),
        ('sorting', ['0001', '0002', '0003', '0004', '0005', '0006', '0007']),
    ]
    
    for cat, problems in missing_recommendations:
        existing = existing_problems[cat]
        print(f"  {cat.capitalize()}: {problems[:5]} (missing {len(set(problems) - existing))}")

if __name__ == "__main__":
    main()