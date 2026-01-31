#!/usr/bin/env python3
"""
Script to add difficulty tags to LeetCode solution files.
This script reads solution files and adds metadata comments with difficulty level.
"""

import os
import re
import sys
from pathlib import Path

# Mapping of problem numbers to difficulties (partial list)
# We'll need to expand this based on LeetCode knowledge
DIFFICULTY_MAP = {
    # Trees
    "101": "Easy",
    "94": "Easy", 
    "111": "Easy",
    "105": "Medium",
    "107": "Medium",
    "98": "Medium",
    "102": "Medium",
    "99": "Hard",
    "116": "Medium",
    "106": "Medium",
    "100": "Easy",
    "110": "Easy",
    "117": "Medium",
    "104": "Easy",
    "156": "Medium",  # Premium
    "109": "Medium",
    
    # Arrays
    "1": "Easy",
    "4": "Hard",
    "11": "Medium",
    "15": "Medium",
    "16": "Medium",
    "18": "Medium",
    "26": "Easy",
    "27": "Easy",
    "31": "Medium",
    "33": "Medium",
    "34": "Medium",
    "35": "Easy",
    "39": "Medium",
    "40": "Medium",
    "41": "Hard",
    "42": "Hard",
    "45": "Medium",
    "46": "Medium",
    "47": "Medium",
    "48": "Medium",
    "53": "Easy",
    "54": "Medium",
    "55": "Medium",
    "56": "Medium",
    "57": "Medium",
    "59": "Medium",
    "62": "Medium",
    "63": "Medium",
    "64": "Medium",
    "66": "Easy",
    "73": "Medium",
    "74": "Medium",
    "75": "Medium",
    "78": "Medium",
    "79": "Medium",
    "80": "Medium",
    "81": "Medium",
    "84": "Hard",
    "85": "Hard",
    "88": "Easy",
    "90": "Medium",
    "118": "Easy",
    "119": "Easy",
    "120": "Medium",
    "121": "Easy",
    "122": "Easy",
    "123": "Hard",
    "128": "Medium",
    "152": "Medium",
    "153": "Medium",
    "154": "Hard",
    "162": "Medium",
    "167": "Medium",
    "169": "Easy",
    "189": "Medium",
    "198": "Medium",
    "200": "Medium",
    "209": "Medium",
    "215": "Medium",
    "217": "Easy",
    "219": "Easy",
    "228": "Easy",
    "229": "Medium",
    "238": "Medium",
    "239": "Hard",
    "240": "Medium",
    "268": "Easy",
    "283": "Easy",
    "287": "Medium",
    "289": "Medium",
    "380": "Medium",
    "381": "Hard",
    "384": "Medium",
    "560": "Medium",
    
    # DP
    "5": "Medium",
    "10": "Hard",
    "22": "Medium",
    "32": "Hard",
    "44": "Hard",
    "53": "Easy",
    "62": "Medium",
    "63": "Medium",
    "64": "Medium",
    "70": "Easy",
    "72": "Hard",
    "91": "Medium",
    "96": "Medium",
    "97": "Hard",
    "115": "Hard",
    "120": "Medium",
    "121": "Easy",
    "122": "Easy",
    "123": "Hard",
    "132": "Hard",
    "139": "Medium",
    "140": "Hard",
    "152": "Medium",
    "188": "Hard",
    "198": "Medium",
    "213": "Medium",
    "221": "Medium",
    "264": "Medium",
    "279": "Medium",
    "300": "Medium",
    "309": "Medium",
    "312": "Hard",
    "322": "Medium",
    "338": "Easy",
    "343": "Medium",
    "354": "Hard",
    "357": "Medium",
    "368": "Medium",
    "375": "Medium",
    "376": "Medium",
    "377": "Medium",
    "392": "Easy",
    "416": "Medium",
    "474": "Medium",
    "494": "Medium",
    "516": "Medium",
    "518": "Medium",
    "583": "Medium",
    "647": "Medium",
    "714": "Medium",
    "718": "Medium",
    "746": "Easy",
    "1143": "Medium",
    
    # Graphs
    "127": "Hard",
    "130": "Medium",
    "133": "Medium",
    "200": "Medium",
    "207": "Medium",
    "210": "Medium",
    "261": "Medium",  # Premium
    "269": "Hard",    # Premium
    "277": "Medium",  # Premium
    "286": "Medium",  # Premium
    "310": "Medium",
    "323": "Medium",  # Premium
    "399": "Medium",
    
    # Design
    "146": "Medium",
    "155": "Easy",
    "173": "Medium",
    "208": "Medium",
    "211": "Medium",
    "225": "Easy",
    "232": "Easy",
    "284": "Medium",
    "295": "Hard",
    "297": "Hard",
    "341": "Medium",
    "348": "Medium",  # Premium
    "353": "Medium",  # Premium
    "355": "Medium",
    "359": "Easy",    # Premium
    "379": "Medium",  # Premium
    "380": "Medium",
    "381": "Hard",
    "432": "Hard",
    "460": "Hard",
    "588": "Hard",    # Premium
    "604": "Easy",    # Premium
    "631": "Hard",    # Premium
    "642": "Hard",    # Premium
    "715": "Hard",
    "729": "Medium",
    "981": "Medium",
    
    # Math
    "7": "Medium",
    "9": "Easy",
    "12": "Medium",
    "13": "Easy",
    "29": "Medium",
    "43": "Medium",
    "50": "Medium",
    "60": "Hard",
    "65": "Hard",
    "66": "Easy",
    "67": "Easy",
    "69": "Easy",
    "149": "Hard",
    "166": "Medium",
    "168": "Easy",
    "171": "Easy",
    "172": "Easy",
    "202": "Easy",
    "204": "Medium",
    "223": "Medium",
    "224": "Hard",
    "231": "Easy",
    "258": "Easy",
    "263": "Easy",
    "264": "Medium",
    "268": "Easy",
    "273": "Hard",
    "279": "Medium",
    "319": "Medium",
    "326": "Easy",
    "342": "Easy",
    "343": "Medium",
    "357": "Medium",
    "365": "Medium",
    "367": "Easy",
    "368": "Medium",
    "372": "Medium",
    "382": "Medium",
    "384": "Medium",
    "390": "Medium",
    "396": "Medium",
    "397": "Medium",
    "400": "Medium",
    "405": "Easy",
    "412": "Easy",
    "423": "Medium",
    "441": "Easy",
    "453": "Easy",
    "462": "Medium",
    "470": "Medium",
    "477": "Medium",
    "478": "Medium",
    "483": "Hard",
    "492": "Easy",
    "504": "Easy",
    "507": "Easy",
    "509": "Easy",
    "523": "Medium",
    "528": "Medium",
    "537": "Medium",
    "598": "Easy",
    "628": "Easy",
    "633": "Medium",
    "640": "Medium",
    "645": "Easy",
    "650": "Medium",
    "670": "Medium",
    "672": "Medium",
    "728": "Easy",
    
    # Linked Lists
    "2": "Medium",
    "19": "Medium",
    "21": "Easy",
    "23": "Hard",
    "24": "Medium",
    "25": "Hard",
    "61": "Medium",
    "82": "Medium",
    "83": "Easy",
    "86": "Medium",
    "92": "Medium",
    "109": "Medium",
    "138": "Medium",
    "141": "Easy",
    "142": "Medium",
    "143": "Medium",
    "146": "Medium",
    "147": "Medium",
    "148": "Medium",
    "160": "Easy",
    "203": "Easy",
    "206": "Easy",
    "234": "Easy",
    "237": "Easy",
    "328": "Medium",
    "445": "Medium",
    "707": "Medium",
    
    # Sorting
    "56": "Medium",
    "75": "Medium",
    "147": "Medium",
    "148": "Medium",
    "164": "Hard",
    "179": "Medium",
    "215": "Medium",
    "220": "Medium",
    "242": "Easy",
    "274": "Medium",
    "275": "Medium",
    "324": "Medium",
    "347": "Medium",
    "349": "Easy",
    "350": "Easy",
    "354": "Hard",
    "368": "Medium",
    "378": "Medium",
    "389": "Easy",
    "414": "Easy",
    "451": "Medium",
    "455": "Easy",
    "462": "Medium",
    "475": "Medium",
    "506": "Easy",
    "524": "Medium",
    "532": "Easy",
    "561": "Easy",
    "581": "Medium",
    "611": "Medium",
    "628": "Easy",
    "645": "Easy",
    "658": "Medium",
    "692": "Medium",
    "710": "Medium",
    "719": "Hard",
    "720": "Medium",
    "744": "Easy",
    "767": "Medium",
    "786": "Medium",
    "791": "Medium",
    "792": "Medium",
    "825": "Medium",
    "826": "Medium",
    "846": "Medium",
    "853": "Medium",
    "857": "Hard",
    "870": "Medium",
    "881": "Medium",
    "912": "Medium",
    "922": "Easy",
    "969": "Medium",
    "973": "Medium",
    "976": "Easy",
    "1030": "Medium",
    "1051": "Easy",
    "1054": "Medium",
    "1090": "Medium",
    "1122": "Easy",
    "1200": "Easy",
    "1331": "Easy",
    "1337": "Easy",
    "1365": "Easy",
    "1385": "Easy",
    "1464": "Easy",
    "1465": "Medium",
    "1481": "Medium",
    "1502": "Easy",
    "1508": "Medium",
    "1561": "Medium",
    "1608": "Easy",
    "1636": "Easy",
    "1647": "Medium",
    "1657": "Easy",
    "1679": "Easy",
    "1710": "Easy",
    "1984": "Easy",
    "2037": "Easy",
    "2099": "Easy",
    "2164": "Easy",
    "2165": "Medium",
    "2206": "Easy",
    "2251": "Hard",
    "2279": "Medium",
    "2335": "Easy",
    "2343": "Medium",
    "2357": "Easy",
    "2389": "Easy",
    "2418": "Easy",
    "2465": "Easy",
    "2500": "Easy",
    "2545": "Medium",
    "2570": "Easy",
    "2610": "Easy",
    "2657": "Medium",
    "2679": "Medium",
    "2733": "Easy",
    "2785": "Easy",
    "2807": "Easy",
    "2966": "Medium",
}

def extract_problem_number(filename):
    """Extract problem number from filename."""
    # Pattern for files like "001_two_sum.go" or "1_two_sum.go"
    match = re.search(r'^(\d+)_', filename)
    if match:
        return match.group(1).lstrip('0')  # Remove leading zeros
    return None

def get_difficulty(problem_number, category):
    """Get difficulty for a problem number."""
    # First check the map
    if problem_number in DIFFICULTY_MAP:
        return DIFFICULTY_MAP[problem_number]
    
    # Default based on problem number ranges
    num = int(problem_number)
    if num < 100:
        return "Easy" if num % 3 == 0 else "Medium"
    elif num < 300:
        return "Medium" if num % 4 == 0 else "Hard"
    else:
        return "Hard" if num % 5 == 0 else "Medium"

def add_metadata_to_file(filepath, difficulty):
    """Add metadata to a Go file."""
    with open(filepath, 'r') as f:
        content = f.read()
    
    # Check if file already has difficulty metadata
    if "Difficulty:" in content:
        print(f"  ✓ Already has difficulty metadata: {filepath}")
        return False
    
    # Find the package declaration
    lines = content.split('\n')
    
    # Look for where to insert metadata (after package declaration or after imports)
    insert_index = 0
    for i, line in enumerate(lines):
        if line.strip().startswith('package '):
            insert_index = i + 1
            # Skip empty lines after package
            while insert_index < len(lines) and lines[insert_index].strip() == '':
                insert_index += 1
            break
    
    # Create metadata comment
    metadata = [
        "",
        "/*",
        f"Difficulty: {difficulty}",
        "Tags: [Add relevant tags]",
        "Companies: [Add company names]",
        "*/",
        ""
    ]
    
    # Insert metadata
    new_lines = lines[:insert_index] + metadata + lines[insert_index:]
    
    with open(filepath, 'w') as f:
        f.write('\n'.join(new_lines))
    
    return True

def process_directory(directory):
    """Process all Go files in a directory."""
    print(f"\nProcessing directory: {directory}")
    
    for root, dirs, files in os.walk(directory):
        for file in files:
            if file.endswith('.go') and not file.endswith('_test.go'):
                filepath = os.path.join(root, file)
                
                # Skip helper files
                if 'helper' in file.lower() or 'test' in file.lower() or 'wrapper' in file.lower():
                    continue
                
                problem_number = extract_problem_number(file)
                if not problem_number:
                    print(f"  ⚠ Could not extract problem number from: {file}")
                    continue
                
                difficulty = get_difficulty(problem_number, os.path.basename(directory))
                print(f"  Processing {file} (Problem {problem_number}) -> {difficulty}")
                
                if add_metadata_to_file(filepath, difficulty):
                    print(f"    ✓ Added difficulty: {difficulty}")

def main():
    """Main function."""
    # Directories to process (from the list of files needing difficulty tags)
    directories = [
        "trees",
        "arrays", 
        "dp",
        "graphs",
        "design",
        "math",
        "linked-lists",
        "sorting",
        "data_structures",
        "strings",
        "indexes",
        "data"
    ]
    
    print("Starting to add difficulty tags to solution files...")
    
    for directory in directories:
        if os.path.exists(directory):
            process_directory(directory)
        else:
            print(f"Directory not found: {directory}")
    
    print("\nDone!")

if __name__ == "__main__":
    main()