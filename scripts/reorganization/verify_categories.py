#!/usr/bin/env python3
"""
Final verification of categorization.
"""

import csv
from collections import defaultdict

def main():
    # Read corrected categorization
    with open('problem_categories_corrected.csv', 'r') as f:
        reader = csv.DictReader(f)
        problems = list(reader)
    
    # Group by category
    by_category = defaultdict(list)
    for p in problems:
        by_category[p['category']].append(p)
    
    # Print detailed breakdown
    print("FINAL CATEGORIZATION VERIFICATION")
    print("=" * 80)
    
    for cat in sorted(by_category.keys()):
        print(f"\n{cat.upper()} ({len(by_category[cat])} problems):")
        print("-" * 40)
        for p in sorted(by_category[cat], key=lambda x: x['number']):
            print(f"  {p['number']}: {p['name']}")
    
    # Check for any uncategorized problems
    all_nums = {p['number'] for p in problems}
    expected_nums = set(f"{i:04d}" for i in range(1, 191))
    missing = expected_nums - all_nums
    if missing:
        print(f"\nWARNING: Missing problems: {sorted(missing)}")
    
    # Check for duplicates
    nums = [p['number'] for p in problems]
    duplicates = {num for num in nums if nums.count(num) > 1}
    if duplicates:
        print(f"\nWARNING: Duplicate problem numbers: {sorted(duplicates)}")
    
    print(f"\n✓ Total problems categorized: {len(problems)}")
    print(f"✓ Categories: {len(by_category)}")
    print(f"✓ Average per category: {len(problems)/len(by_category):.1f}")

if __name__ == '__main__':
    main()