#!/usr/bin/env python3
"""
Generate a quick overview of problem categorization for manual review.
"""

import csv

def main():
    with open('problem_categories.csv', 'r') as f:
        reader = csv.DictReader(f)
        problems = list(reader)
    
    # Group by category
    by_category = {}
    for p in problems:
        cat = p['category']
        if cat not in by_category:
            by_category[cat] = []
        by_category[cat].append(p)
    
    # Print summary
    print("PROBLEM CATEGORIZATION SUMMARY")
    print("=" * 60)
    
    for cat in sorted(by_category.keys()):
        print(f"\n{cat.upper()} ({len(by_category[cat])} problems):")
        print("-" * 40)
        for p in sorted(by_category[cat], key=lambda x: x['number']):
            print(f"  {p['number']}: {p['name']}")
    
    # Check for potential mis-categorizations
    print("\n\nPOTENTIAL ISSUES TO REVIEW:")
    print("=" * 60)
    
    # Check for graph problems in wrong categories
    graph_keywords = ['graph', 'ladder', 'course', 'island', 'clone']
    for p in problems:
        name_lower = p['name'].lower()
        if any(kw in name_lower for kw in graph_keywords) and p['category'] != 'graphs':
            print(f"  {p['number']}: {p['name']} -> currently {p['category']}, might be graphs")
    
    # Check for DP problems in wrong categories
    dp_keywords = ['stock', 'palindrome partitioning', 'word break', 'decode', 'interleaving']
    for p in problems:
        name_lower = p['name'].lower()
        if any(kw in name_lower for kw in dp_keywords) and p['category'] != 'dp':
            print(f"  {p['number']}: {p['name']} -> currently {p['category']}, might be dp")

if __name__ == '__main__':
    main()