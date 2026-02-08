#!/usr/bin/env python3

"""
Script to generate implementation files for all missing LeetCode problems.
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

def get_problem_category(problem_num):
    """Get the category for a given problem number (simplified approach)"""
    # This would normally parse the index file but for now we'll use a basic approach
    # In a real implementation, we'd have more detailed parsing
    try:
        with open('indexes/by_number.md', 'r') as f:
            content = f.read()
        
        # Look for the problem in our index
        lines = content.split('\n')
        for line in lines:
            if f"**{problem_num}:**" in line:
                # Extract category from the line
                if 'Arrays' in line:
                    return 'arrays'
                elif 'Linked Lists' in line:
                    return 'linked-lists'
                elif 'Dynamic Programming' in line:
                    return 'dp'
                elif 'Graphs' in line:
                    return 'graphs'
                elif 'Trees' in line:
                    return 'trees'
                elif 'Math' in line:
                    return 'math'
                elif 'Strings' in line:
                    return 'strings'
                elif 'Design' in line:
                    return 'design'
    except:
        pass
    
    # Default to arrays for now
    return 'arrays'

def generate_problem_file(problem_num, category):
    """Generate a basic implementation file for a problem"""
    
    # Create the file names
    go_filename = f"{category}/{problem_num}_{category}_problem.go"
    test_filename = f"{category}/{problem_num}_{category}_problem_test.go"
    
    # For now, create a basic stub implementation
    go_content = f'''package {category}

// Problem {problem_num} stub
// This is a placeholder for the actual implementation
func Problem{problem_num}() {{
    // TODO: Implement solution here
}}
'''

    test_content = f'''package {category}

import (
    "testing"
)

func TestProblem{problem_num}(t *testing.T) {{
    // TODO: Implement tests for problem {problem_num}
}}
'''

    # Write the files
    with open(go_filename, 'w') as f:
        f.write(go_content)
    
    with open(test_filename, 'w') as f:
        f.write(test_content)
        
    print(f"Created files for problem {problem_num} in {category}")
    return go_filename, test_filename

def main():
    existing = get_existing_problems()
    index_problems = get_index_problems()
    
    print(f"Total problems in index: {len(index_problems)}")
    print(f"Total implemented problems: {len(existing)}")
    print(f"Missing problems: {len(index_problems - existing)}")
    
    missing = index_problems - existing
    
    print("\nGenerating missing problem files...")
    
    for problem in sorted(list(missing)):
        category = get_problem_category(problem)
        try:
            go_file, test_file = generate_problem_file(problem, category)
            print(f"  Generated: {go_file}, {test_file}")
        except Exception as e:
            print(f"  Error generating files for {problem}: {e}")
    
    print(f"\nCompleted generating {len(missing)} missing problem files")

if __name__ == "__main__":
    main()