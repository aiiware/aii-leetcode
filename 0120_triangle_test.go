package leetcode

import (
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		name     string
		triangle [][]int
		expected int
	}{
		{
			name:     "empty triangle",
			triangle: [][]int{},
			expected: 0,
		},
		{
			name: "single row",
			triangle: [][]int{
				{2},
			},
			expected: 2,
		},
		{
			name: "two rows",
			triangle: [][]int{
				{2},
				{3, 4},
			},
			expected: 5, // 2 + 3
		},
		{
			name: "three rows",
			triangle: [][]int{
				{2},
				{3, 4},
				{6, 5, 7},
			},
			expected: 10, // 2 + 3 + 5
		},
		{
			name: "example from problem",
			triangle: [][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			},
			expected: 11, // 2 + 3 + 5 + 1
		},
		{
			name: "all same numbers",
			triangle: [][]int{
				{1},
				{1, 1},
				{1, 1, 1},
			},
			expected: 3,
		},
		{
			name: "negative numbers",
			triangle: [][]int{
				{-1},
				{2, 3},
				{1, -1, -3},
			},
			expected: -1, // -1 + 3 + -3 = -1 (corrected from -2)
		},
		{
			name: "large triangle",
			triangle: [][]int{
				{1},
				{2, 3},
				{4, 5, 6},
				{7, 8, 9, 10},
				{11, 12, 13, 14, 15},
			},
			expected: 25, // 1 + 2 + 4 + 7 + 11 = 25 (corrected from 20)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy since minimumTotal modifies the triangle
			triangleCopy := createTriangle(tt.triangle)
			result := minimumTotal(triangleCopy)
			if result != tt.expected {
				t.Errorf("minimumTotal() = %d, expected %d", result, tt.expected)
			}
		})
	}
}

func TestMinimumTotalDP(t *testing.T) {
	tests := []struct {
		name     string
		triangle [][]int
		expected int
	}{
		{
			name:     "empty triangle",
			triangle: [][]int{},
			expected: 0,
		},
		{
			name: "example from problem",
			triangle: [][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			},
			expected: 11,
		},
		{
			name: "single row",
			triangle: [][]int{
				{5},
			},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumTotalDP(tt.triangle)
			if result != tt.expected {
				t.Errorf("minimumTotalDP() = %d, expected %d", result, tt.expected)
			}
		})
	}
}

func TestMinimumTotalDFS(t *testing.T) {
	tests := []struct {
		name     string
		triangle [][]int
		expected int
	}{
		{
			name:     "empty triangle",
			triangle: [][]int{},
			expected: 0,
		},
		{
			name: "example from problem",
			triangle: [][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			},
			expected: 11,
		},
		{
			name: "small triangle",
			triangle: [][]int{
				{1},
				{2, 3},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumTotalDFS(tt.triangle)
			if result != tt.expected {
				t.Errorf("minimumTotalDFS() = %d, expected %d", result, tt.expected)
			}
		})
	}
}

func TestMinimumTotalBruteForce(t *testing.T) {
	tests := []struct {
		name     string
		triangle [][]int
		expected int
	}{
		{
			name: "small triangle",
			triangle: [][]int{
				{1},
				{2, 3},
			},
			expected: 3,
		},
		{
			name: "three rows",
			triangle: [][]int{
				{1},
				{2, 3},
				{4, 5, 6},
			},
			expected: 7, // 1 + 2 + 4
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumTotalBruteForce(tt.triangle)
			if result != tt.expected {
				t.Errorf("minimumTotalBruteForce() = %d, expected %d", result, tt.expected)
			}
		})
	}
}

func TestCreateTriangle(t *testing.T) {
	tests := []struct {
		name string
		nums [][]int
	}{
		{
			name: "empty",
			nums: [][]int{},
		},
		{
			name: "simple triangle",
			nums: [][]int{
				{1},
				{2, 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			triangle := createTriangle(tt.nums)
			
			// Check dimensions
			if len(triangle) != len(tt.nums) {
				t.Errorf("createTriangle() returned %d rows, expected %d", len(triangle), len(tt.nums))
				return
			}
			
			// Check each row
			for i := range triangle {
				if len(triangle[i]) != len(tt.nums[i]) {
					t.Errorf("row %d: got length %d, expected %d", i, len(triangle[i]), len(tt.nums[i]))
					continue
				}
				for j := range triangle[i] {
					if triangle[i][j] != tt.nums[i][j] {
						t.Errorf("triangle[%d][%d] = %d, expected %d", i, j, triangle[i][j], tt.nums[i][j])
					}
				}
			}
			
			// Verify it's a deep copy (modifying original shouldn't affect copy)
			if len(tt.nums) > 0 && len(tt.nums[0]) > 0 {
				originalValue := tt.nums[0][0]
				tt.nums[0][0] = originalValue + 100
				if triangle[0][0] == tt.nums[0][0] {
					t.Error("createTriangle() didn't create a deep copy")
				}
				// Restore original
				tt.nums[0][0] = originalValue
			}
		})
	}
}

func BenchmarkMinimumTotal(b *testing.B) {
	triangle := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
		{7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20, 21},
		{22, 23, 24, 25, 26, 27, 28},
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		triangleCopy := createTriangle(triangle)
		minimumTotal(triangleCopy)
	}
}

func BenchmarkMinimumTotalDP(b *testing.B) {
	triangle := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
		{7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20, 21},
		{22, 23, 24, 25, 26, 27, 28},
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		minimumTotalDP(triangle)
	}
}

func BenchmarkMinimumTotalDFS(b *testing.B) {
	triangle := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
		{7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20, 21},
		{22, 23, 24, 25, 26, 27, 28},
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		minimumTotalDFS(triangle)
	}
}