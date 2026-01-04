package leetcode

// SimplifyPath solves LeetCode problem 0071: Simplify Path
// Difficulty: Medium
// Tags: String, Stack
//
// Given a string path, which is an absolute path (starting with a slash '/') to a file or directory
// in a Unix-style file system, convert it to the simplified canonical path.
//
// In a Unix-style file system, a period '.' refers to the current directory, a double period '..'
// refers to the directory up a level, and any multiple consecutive slashes (i.e., '//') are treated
// as a single slash '/'. For this problem, any other format of periods such as '...' are treated as
// file/directory names.
//
// The canonical path should have the following format:
// - The path starts with a single slash '/'.
// - Any two directories are separated by a single slash '/'.
// - The path does not end with a trailing '/'.
// - The path only contains the directories on the path from the root directory to the target file
//   or directory (i.e., no period '.' or double period '..').
//
// Return the simplified canonical path.
//
// Example 1:
// Input: path = "/home/"
// Output: "/home"
// Explanation: Note that there is no trailing slash after the last directory name.
//
// Example 2:
// Input: path = "/../"
// Output: "/"
// Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.
//
// Example 3:
// Input: path = "/home//foo/"
// Output: "/home/foo"
// Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.
//
// Example 4:
// Input: path = "/a/./b/../../c/"
// Output: "/c"
//
// Example 5:
// Input: path = "/a/../../b/../c//.//"
// Output: "/c"
//
// Example 6:
// Input: path = "/a//b////c/d//././/.."
// Output: "/a/b/c"
//
// Constraints:
// 1 <= path.length <= 3000
// path consists of English letters, digits, period '.', slash '/' or '_'.
// path is a valid absolute Unix path.
//
// Time complexity: O(n), Space complexity: O(n)
func SimplifyPath(path string) string {
	if len(path) == 0 {
		return "/"
	}

	// Split the path by '/'
	parts := make([]string, 0)
	start := 0
	n := len(path)

	for i := 0; i < n; i++ {
		// Skip consecutive slashes
		if path[i] == '/' {
			continue
		}

		// Find the end of the current component
		start = i
		for i < n && path[i] != '/' {
			i++
		}

		// Extract the component
		component := path[start:i]

		switch component {
		case ".":
			// Current directory, do nothing
			continue
		case "..":
			// Go up one directory if possible
			if len(parts) > 0 {
				parts = parts[:len(parts)-1]
			}
		default:
			// Regular directory/file name
			parts = append(parts, component)
		}
	}

	// Build the canonical path
	if len(parts) == 0 {
		return "/"
	}

	result := ""
	for _, part := range parts {
		result += "/" + part
	}

	return result
}