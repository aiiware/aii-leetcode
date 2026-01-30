#!/usr/bin/env python3
"""
Create final corrected categorization based on manual review.
"""

import csv

# Manual corrections based on README tags and problem analysis
MANUAL_CORRECTIONS = {
    # Graph problems
    '0126': 'graphs',  # Word Ladder II - BFS/Graph
    '0127': 'graphs',  # Word Ladder - BFS/Graph
    '0133': 'graphs',  # Clone Graph - Graph
    
    # DP problems currently in strings
    '0091': 'dp',      # Decode Ways - DP
    '0097': 'dp',      # Interleaving String - DP
    '0131': 'dp',      # Palindrome Partitioning - DP/Backtracking
    '0132': 'dp',      # Palindrome Partitioning II - DP
    '0139': 'dp',      # Word Break - DP
    '0140': 'dp',      # Word Break II - DP/Backtracking
    
    # DP problems currently in arrays
    '0042': 'dp',      # Trapping Rain Water - DP/Two Pointers
    '0053': 'dp',      # Maximum Subarray - DP (Kadane's)
    '0062': 'dp',      # Unique Paths - DP (already correct)
    '0063': 'dp',      # Unique Paths II - DP (already correct)
    '0064': 'dp',      # Minimum Path Sum - DP
    '0070': 'dp',      # Climbing Stairs - DP (already correct)
    '0072': 'dp',      # Edit Distance - DP (already correct)
    '0084': 'dp',      # Largest Rectangle in Histogram - DP/Stack
    '0085': 'dp',      # Maximal Rectangle - DP
    '0112': 'dp',      # Path Sum - DP/DFS (tree but DP)
    '0113': 'dp',      # Path Sum II - DP/Backtracking
    '0120': 'dp',      # Triangle - DP (already correct)
    '0124': 'dp',      # Binary Tree Maximum Path Sum - DP/DFS
    
    # Tree problems that are actually DP
    '0095': 'dp',      # Unique Binary Search Trees II - DP/Backtracking
    '0096': 'dp',      # Unique Binary Search Trees - DP (Catalan)
    
    # Array problems that are actually math
    '0048': 'math',    # Rotate Image - Math/Matrix
    '0054': 'math',    # Spiral Matrix - Math/Simulation
    '0059': 'math',    # Spiral Matrix II - Math/Simulation
    
    # Array problems that are actually sorting
    '0075': 'sorting', # Sort Colors - Sorting/Two Pointers
    '0088': 'sorting', # Merge Sorted Array - Sorting
    
    # Design problems
    '0146': 'design',  # LRU Cache - Design (already correct)
    '0155': 'design',  # Min Stack - Design (already correct)
    '0170': 'design',  # Two Sum III - Design
    
    # SQL problems (if any)
    '0175': 'sql',     # Combine Two Tables - SQL
    '0176': 'sql',     # Second Highest Salary - SQL
    '0177': 'sql',     # Nth Highest Salary - SQL
    '0178': 'sql',     # Rank Scores - SQL
    '0181': 'sql',     # Employees Earning More - SQL
    '0182': 'sql',     # Duplicate Emails - SQL
    '0183': 'sql',     # Customers Who Never Order - SQL
    '0184': 'sql',     # Department Highest Salary - SQL
    '0185': 'sql',     # Department Top Three Salaries - SQL
}

def main():
    # Read current categorization
    with open('problem_categories.csv', 'r') as f:
        reader = csv.DictReader(f)
        problems = list(reader)
    
    # Apply corrections
    corrected = []
    for p in problems:
        num = p['number']
        if num in MANUAL_CORRECTIONS:
            p['category'] = MANUAL_CORRECTIONS[num]
        corrected.append(p)
    
    # Write corrected file
    with open('problem_categories_corrected.csv', 'w') as f:
        writer = csv.DictWriter(f, fieldnames=['number', 'name', 'filename', 'category'])
        writer.writeheader()
        writer.writerows(corrected)
    
    # Print summary
    from collections import Counter
    categories = Counter(p['category'] for p in corrected)
    
    print("CORRECTED CATEGORIZATION SUMMARY")
    print("=" * 60)
    for cat, count in sorted(categories.items(), key=lambda x: x[1], reverse=True):
        print(f"{cat:15}: {count:3} problems")
    
    print(f"\nTotal problems: {sum(categories.values())}")
    print("\nCorrections applied:")
    for num, new_cat in MANUAL_CORRECTIONS.items():
        old = next(p for p in problems if p['number'] == num)
        print(f"  {num}: {old['name'][:30]:30} {old['category']:10} -> {new_cat}")
    
    print(f"\nSaved to problem_categories_corrected.csv")

if __name__ == '__main__':
    main()