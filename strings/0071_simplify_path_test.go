package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "Example 1: /home/",
			path:     "/home/",
			expected: "/home",
		},
		{
			name:     "Example 2: /../",
			path:     "/../",
			expected: "/",
		},
		{
			name:     "Example 3: /home//foo/",
			path:     "/home//foo/",
			expected: "/home/foo",
		},
		{
			name:     "Example 4: /a/./b/../../c/",
			path:     "/a/./b/../../c/",
			expected: "/c",
		},
		{
			name:     "Example 5: /a/../../b/../c//.//",
			path:     "/a/../../b/../c//.//",
			expected: "/c",
		},
		{
			name:     "Example 6: /a//b////c/d//././/..",
			path:     "/a//b////c/d//././/..",
			expected: "/a/b/c",
		},
		{
			name:     "Empty path",
			path:     "",
			expected: "/",
		},
		{
			name:     "Root path",
			path:     "/",
			expected: "/",
		},
		{
			name:     "Multiple dots treated as filename",
			path:     "/...",
			expected: "/...",
		},
		{
			name:     "Multiple dots with slashes",
			path:     "/.../...",
			expected: "/.../...",
		},
		{
			name:     "Complex path with underscores",
			path:     "/home/user/_docs/../files/./report.txt",
			expected: "/home/user/files/report.txt",
		},
		{
			name:     "Path with numbers",
			path:     "/usr/local/bin/../lib64/./python3",
			expected: "/usr/local/lib64/python3",
		},
		{
			name:     "Multiple consecutive ..",
			path:     "/a/b/c/../../../",
			expected: "/",
		},
		{
			name:     "Cannot go above root",
			path:     "/../../../../",
			expected: "/",
		},
		{
			name:     "Mixed case and special characters",
			path:     "/Home/User/Documents/../Downloads/./../Pictures/./Vacation",
			expected: "/Home/User/Pictures/Vacation",
		},
		{
			name:     "Path ending with .",
			path:     "/home/user/.",
			expected: "/home/user",
		},
		{
			name:     "Path ending with ..",
			path:     "/home/user/..",
			expected: "/home",
		},
		{
			name:     "Path with only . and ..",
			path:     "/./././../..",
			expected: "/",
		},
		{
			name:     "Long path with many components",
			path:     "/a/b/c/d/e/f/g/../../../../../../..",
			expected: "/",
		},
		{
			name:     "Path with spaces in component names",
			path:     "/my%20documents/../photos/./vacation%202024",
			expected: "/photos/vacation%202024",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SimplifyPath(tt.path)
			assert.Equal(t, tt.expected, result,
				"SimplifyPath(%q) = %q, expected %q",
				tt.path, result, tt.expected)
		})
	}
}

func TestSimplifyPath_EdgeCases(t *testing.T) {
	t.Run("Very long path", func(t *testing.T) {
		// Build a long path
		longPath := "/"
		for i := 0; i < 100; i++ {
			longPath += "dir" + string(rune('a'+i%26)) + "/"
		}
		longPath += ".."
		
		// Expected: all directories except the last ".."
		expected := "/"
		for i := 0; i < 99; i++ {
			expected += "dir" + string(rune('a'+i%26)) + "/"
		}
		expected = expected[:len(expected)-1] // Remove trailing slash
		
		result := SimplifyPath(longPath)
		assert.Equal(t, expected, result)
	})

	t.Run("Path with many consecutive slashes", func(t *testing.T) {
		result := SimplifyPath("////////home//////user////////docs////")
		assert.Equal(t, "/home/user/docs", result)
	})

	t.Run("Path starting with multiple slashes", func(t *testing.T) {
		result := SimplifyPath("//home/user")
		assert.Equal(t, "/home/user", result)
	})

	t.Run("Path with only dots and slashes", func(t *testing.T) {
		result := SimplifyPath("/./././././")
		assert.Equal(t, "/", result)
	})

	t.Run("Path with .. at beginning", func(t *testing.T) {
		result := SimplifyPath("/../home/../user/../docs")
		assert.Equal(t, "/docs", result)
	})
}

func TestSimplifyPath_Consistency(t *testing.T) {
	// Test that simplifying an already simplified path doesn't change it
	testPaths := []string{
		"/home",
		"/home/user",
		"/usr/local/bin",
		"/",
		"/a/b/c",
	}

	for _, path := range testPaths {
		t.Run("Already simplified: "+path, func(t *testing.T) {
			result := SimplifyPath(path)
			assert.Equal(t, path, result,
				"SimplifyPath should not change already simplified path")
		})
	}

	// Test idempotence: simplify(simplify(path)) == simplify(path)
	complexPaths := []string{
		"/home//user/../docs/./",
		"/a/./b/../../c/",
		"/../usr/./local/../bin/",
	}

	for _, path := range complexPaths {
		t.Run("Idempotence: "+path, func(t *testing.T) {
			once := SimplifyPath(path)
			twice := SimplifyPath(once)
			assert.Equal(t, once, twice,
				"SimplifyPath should be idempotent")
		})
	}
}

func BenchmarkSimplifyPath(b *testing.B) {
	testCases := []struct {
		name string
		path string
	}{
		{"Short path", "/home/user"},
		{"Medium path", "/usr/local/bin/../lib/./python3"},
		{"Long path", "/a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t/u/v/w/x/y/z"},
		{"Many dots", "/./././.././.././././../"},
		{"Many slashes", "////////home//////user////////docs////"},
		{"Complex path", "/home/user/../docs/./files/../../photos/./vacation/../2024"},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SimplifyPath(tc.path)
			}
		})
	}
}