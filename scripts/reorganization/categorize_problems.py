#!/usr/bin/env python3
"""
Categorize LeetCode problems into algorithm-focused categories.
Based on problem tags from README.md and problem names.
"""

import re
import os
from pathlib import Path

# Category definitions with keywords
CATEGORIES = {
    'arrays': [
        'array', 'two pointers', 'sliding window', 'matrix', 'subarray',
        'rotate', 'spiral', 'combination sum', 'permutation', 'subsets',
        'container', 'sum', 'median', 'trap', 'jump', 'search insert',
        'first missing', 'rotate image', 'set matrix zeroes', 'sort colors',
        'merge intervals', 'insert interval', 'largest rectangle',
        'maximal rectangle', 'gas station', 'candy', 'product of array',
        'majority element', 'h-index', 'rotate array', 'find duplicate',
        'missing number', 'first bad version'
    ],
    'strings': [
        'string', 'palindrome', 'anagram', 'substring', 'window',
        'zigzag', 'atoi', 'regular expression', 'wildcard', 'multiply',
        'count and say', 'simplify path', 'text justification', 'word search',
        'decode ways', 'restore ip', 'interleaving', 'distinct subsequences',
        'word break', 'word ladder', 'valid number', 'integer to roman',
        'roman to integer', 'longest common prefix', 'letter combinations',
        'generate parentheses', 'substring with concatenation'
    ],
    'linked-lists': [
        'linked list', 'list node', 'add two numbers', 'remove nth node',
        'merge lists', 'swap nodes', 'reverse nodes', 'rotate list',
        'remove duplicates list', 'partition list', 'reverse linked list',
        'copy list with random', 'reorder list', 'insertion sort list',
        'sort list', 'linked list cycle', 'flatten binary tree to linked list'
    ],
    'trees': [
        'tree', 'binary tree', 'bst', 'binary search tree', 'inorder',
        'preorder', 'postorder', 'level order', 'zigzag', 'symmetric',
        'same tree', 'validate bst', 'recover bst', 'construct tree',
        'convert sorted', 'balanced', 'minimum depth', 'maximum depth',
        'path sum', 'flatten', 'populating next right pointers',
        'binary tree maximum path sum', 'sum root to leaf',
        'surrounded regions', 'clone graph'  # graph but tree-like
    ],
    'dp': [
        'dynamic programming', 'dp', 'climbing stairs', 'unique paths',
        'minimum path sum', 'edit distance', 'maximum subarray',
        'longest palindromic substring', 'regular expression matching',
        'wildcard matching', 'jump game', 'jump game ii', 'decode ways',
        'interleaving string', 'distinct subsequences', 'word break',
        'word break ii', 'palindrome partitioning', 'scramble string',
        'best time to buy and sell stock', 'triangle', 'pascal triangle',
        'candy', 'house robber', 'maximal square', 'perfect squares',
        'coin change', 'combination sum iv', 'integer break'
    ],
    'graphs': [
        'graph', 'clone graph', 'course schedule', 'alien dictionary',
        'number of islands', 'surrounded regions', 'word ladder',
        'word ladder ii', 'minimum height trees', 'reconstruct itinerary',
        'cheapest flights', 'network delay time', 'walls and gates',
        'course schedule ii', 'graph valid tree'
    ],
    'math': [
        'math', 'integer', 'reverse integer', 'divide', 'pow', 'sqrt',
        'plus one', 'add binary', 'gray code', 'single number',
        'single number ii', 'factorial trailing zeroes', 'happy number',
        'count primes', 'excel sheet column title', 'excel sheet column number',
        'power of two', 'power of three', 'ugly number', 'super ugly number',
        'nth digit', 'integer replacement', 'random pick index'
    ],
    'sorting': [
        'sort', 'sorting', 'binary search', 'search rotated', 'search 2d',
        'find peak element', 'find minimum in rotated', 'search range',
        'merge sorted array', 'sort list', 'insertion sort list',
        'largest number', 'wiggle sort', 'kth largest element',
        'top k frequent elements', 'meeting rooms', 'meeting rooms ii'
    ],
    'design': [
        'design', 'lru cache', 'min stack', 'implement queue using stacks',
        'implement stack using queues', 'two sum iii', 'zigzag iterator',
        'peeking iterator', 'flatten 2d vector', 'range sum query',
        'serialize and deserialize', 'tic-tac-toe', 'logger rate limiter'
    ]
}

def read_problem_list():
    """Read all problem files from the current directory."""
    problems = []
    for file in Path('.').glob('[0-9][0-9][0-9][0-9]_*.go'):
        if file.name.endswith('_test.go'):
            continue
        # Extract problem number and name
        match = re.match(r'(\d{4})_(.+)\.go', file.name)
        if match:
            num = match.group(1)
            name = match.group(2).replace('_', ' ')
            problems.append((num, name, file.name))
    return sorted(problems, key=lambda x: x[0])

def categorize_problem(num, name):
    """Categorize a problem based on its name."""
    name_lower = name.lower()
    
    # Check each category
    for category, keywords in CATEGORIES.items():
        for keyword in keywords:
            if keyword in name_lower:
                return category
    
    # Default based on common patterns
    if 'tree' in name_lower:
        return 'trees'
    elif 'list' in name_lower:
        return 'linked-lists'
    elif 'sum' in name_lower or 'array' in name_lower:
        return 'arrays'
    elif 'string' in name_lower:
        return 'strings'
    elif 'number' in name_lower or 'integer' in name_lower:
        return 'math'
    
    return 'arrays'  # default

def main():
    """Main function to categorize problems."""
    problems = read_problem_list()
    
    print(f"Found {len(problems)} problems")
    print("\nCategorization mapping:\n")
    
    # Count by category
    category_count = {cat: 0 for cat in CATEGORIES.keys()}
    mappings = []
    
    for num, name, filename in problems:
        category = categorize_problem(num, name)
        category_count[category] += 1
        mappings.append((num, name, filename, category))
        
        print(f"{num}: {name:40} -> {category}")
    
    print("\nCategory counts:")
    for cat, count in sorted(category_count.items(), key=lambda x: x[1], reverse=True):
        print(f"  {cat:15}: {count:3} problems")
    
    # Write mapping to file
    with open('problem_categories.csv', 'w') as f:
        f.write("number,name,filename,category\n")
        for num, name, filename, category in mappings:
            f.write(f'{num},"{name}","{filename}","{category}"\n')
    
    print(f"\nMapping saved to problem_categories.csv")

if __name__ == '__main__':
    main()