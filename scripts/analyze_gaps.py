#!/usr/bin/env python3
"""
Script to analyze gaps in LeetCode problem implementations
"""

import os
import re
from collections import defaultdict

def get_implemented_problems():
    """Get all implemented problems from the codebase"""
    implemented = defaultdict(list)
    
    # Scan all directories for problem files
    for root, dirs, files in os.walk('.'):
        for file in files:
            if file.endswith('.go'):
                # Extract problem number from filename
                match = re.search(r'([0-9]{4})_', file)
                if match:
                    problem_num = match.group(1)
                    # Determine category from directory
                    category = os.path.basename(root)
                    if category == '.':
                        category = 'unknown'
                    implemented[category].append(problem_num)
    
    return implemented

def analyze_gaps():
    """Analyze current gaps in implementation"""
    implemented = get_implemented_problems()
    
    print("=== CURRENT IMPLEMENTED PROBLEMS BY CATEGORY ===")
    
    # Analyze each category
    for category, problems in implemented.items():
        if category in ['.', 'cmd', 'docs', 'scripts', 'testutils', 'data']:
            continue
            
        print(f"\n{category.upper()}: {len(problems)} problems")
        if problems:
            sorted_probs = sorted(set(problems))
            print(f"  Problem range: {sorted_probs[0]} - {sorted_probs[-1]}")
            
            # Find gaps
            gaps = []
            for i in range(len(sorted_probs) - 1):
                current = int(sorted_probs[i])
                next_prob = int(sorted_probs[i + 1])
                if next_prob - current > 1:
                    gaps.append((current, next_prob))
            
            if gaps:
                print(f"  Gaps found:")
                for start, end in gaps:
                    print(f"    {start + 1} - {end - 1}")
            else:
                print("  No gaps found")
    
    # Total analysis
    total_problems = sum(len(probs) for probs in implemented.values() if probs)
    print(f"\n=== SUMMARY ===")
    print(f"Total implemented problems: {total_problems}")
    
    # Check what would be the target count (1000 problems)
    print("Target: 1000 problems")
    print(f"Missing: {1000 - total_problems}")

if __name__ == "__main__":
    analyze_gaps()