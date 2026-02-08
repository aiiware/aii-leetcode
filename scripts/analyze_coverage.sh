#!/bin/bash

# Script to analyze problem coverage

echo "Analyzing problem coverage..."

# Count problems in arrays
arrays_count=$(find arrays -name "*.go" -exec basename {} \; | grep -E "^[0-9]{4}_" | wc -l)
echo "Arrays problems: $arrays_count"

# Count problems in dp
dp_count=$(find dp -name "*.go" -exec basename {} \; | grep -E "^[0-9]{4}_" | wc -l)
echo "DP problems: $dp_count"

# Count problems in trees
trees_count=$(find trees -name "*.go" -exec basename {} \; | grep -E "^[0-9]{4}_" | wc -l)
echo "Trees problems: $trees_count"

# Count problems in graphs
graphs_count=$(find graphs -name "*.go" -exec basename {} \; | grep -E "^[0-9]{4}_" | wc -l)
echo "Graphs problems: $graphs_count"

# Count problems in linked-lists
linked_lists_count=$(find linked-lists -name "*.go" -exec basename {} \; | grep -E "^[0-9]{4}_" | wc -l)
echo "Linked Lists problems: $linked_lists_count"

# Count problems in math
math_count=$(find math -name "*.go" -exec basename {} \; | grep -E "^[0-9]{4}_" | wc -l)
echo "Math problems: $math_count"

# Count problems in strings
strings_count=$(find strings -name "*.go" -exec basename {} \; | grep -E "^[0-9]{4}_" | wc -l)
echo "Strings problems: $strings_count"

# Count problems in design
design_count=$(find design -name "*.go" -exec basename {} \; | grep -E "^[0-9]{4}_" | wc -l)
echo "Design problems: $design_count"

echo "Total implemented problems: $((arrays_count + dp_count + trees_count + graphs_count + linked_lists_count + math_count + strings_count + design_count))"