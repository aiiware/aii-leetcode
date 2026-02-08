#!/usr/bin/env python3
"""
Generate standard problem templates for missing LeetCode problems
"""

import os
import re
import sys

def generate_problem_template(problem_number, problem_name, category, difficulty, tags):
    """
    Generate a standard Go problem template
    """
    # Format problem number with leading zeros
    formatted_number = f"{int(problem_number):04d}"
    
    # Clean problem name
    clean_name = re.sub(r'[^a-zA-Z0-9_]', '_', problem_name)
    
    # Create function name (remove underscores and capitalize first letter of each word)
    function_name = ''.join(word.capitalize() for word in clean_name.split('_'))
    if function_name[0].isdigit():
        function_name = 'Problem' + function_name
    
    # Create file names
    go_filename = f"{formatted_number}_{clean_name}.go"
    test_filename = f"{formatted_number}_{clean_name}_test.go"
    
    # Define package
    package = category.lower()
    
    # Generate Go file content
    go_content = f'''package {package}

// {function_name} solves LeetCode problem {formatted_number}: {problem_name}
// Difficulty: {difficulty}
// Tags: {', '.join(tags)}
//
// Description
//
// Time complexity: O(n), Space complexity: O(n)
func {function_name}() int {{
    // TODO: Implement solution
    return 0
}}
'''
    
    # Generate test file content
    test_content = f'''package {package}

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func Test{function_name}(t *testing.T) {{
    tests := []struct {{
        name string
        // Add your test parameters here
        want int
    }}{{
        {{
            name: "Example 1",
            // Add test parameters here
            want: 0,
        }},
    }}

    for _, tt := range tests {{
        t.Run(tt.name, func(t *testing.T) {{
            // result := {function_name}()
            // assert.Equal(t, tt.want, result)
            t.Skip("Not implemented")
        }})
    }}
}}

func Benchmark{function_name}(b *testing.B) {{
    // TODO: Implement benchmark
    b.Skip("Not implemented")
}}
'''
    
    return {
        'go': go_content,
        'test': test_content,
        'go_filename': go_filename,
        'test_filename': test_filename
    }

def main():
    print("Problem Template Generator")
    print("This script generates standard problem templates for missing LeetCode problems.")
    
    # For now, let's generate a simple example to test the approach
    problem_number = "0006"
    problem_name = "Longest Palindromic Substring"
    category = "arrays"
    difficulty = "Medium"
    tags = ["String", "Dynamic Programming"]
    
    template = generate_problem_template(problem_number, problem_name, category, difficulty, tags)
    
    print("\nGenerated template for problem 0006:")
    print("=" * 50)
    print("Go file content:")
    print(template['go'])
    print("\nTest file content:")
    print(template['test'])
    
    # Write files to disk for testing
    os.makedirs(category, exist_ok=True)
    
    with open(os.path.join(category, template['go_filename']), 'w') as f:
        f.write(template['go'])
    
    with open(os.path.join(category, template['test_filename']), 'w') as f:
        f.write(template['test'])
    
    print(f"\nFiles written to {category}/:")
    print(f"  {template['go_filename']}")
    print(f"  {template['test_filename']}")

if __name__ == "__main__":
    main()