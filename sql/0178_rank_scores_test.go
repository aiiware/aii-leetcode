package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRankScores(t *testing.T) {
	tests := []struct {
		name   string
		scores []Score
		expected []RankedScore
	}{
		{
			name: "Example 1 from LeetCode",
			scores: []Score{
				{ID: 1, Score: 3.50},
				{ID: 2, Score: 3.65},
				{ID: 3, Score: 4.00},
				{ID: 4, Score: 3.85},
				{ID: 5, Score: 4.00},
				{ID: 6, Score: 3.65},
			},
			expected: []RankedScore{
				{Score: 4.00, Rank: 1},
				{Score: 4.00, Rank: 1},
				{Score: 3.85, Rank: 2},
				{Score: 3.65, Rank: 3},
				{Score: 3.65, Rank: 3},
				{Score: 3.50, Rank: 4},
			},
		},
		{
			name: "All same scores",
			scores: []Score{
				{ID: 1, Score: 100.0},
				{ID: 2, Score: 100.0},
				{ID: 3, Score: 100.0},
			},
			expected: []RankedScore{
				{Score: 100.0, Rank: 1},
				{Score: 100.0, Rank: 1},
				{Score: 100.0, Rank: 1},
			},
		},
		{
			name: "Strictly decreasing scores",
			scores: []Score{
				{ID: 1, Score: 500.0},
				{ID: 2, Score: 400.0},
				{ID: 3, Score: 300.0},
				{ID: 4, Score: 200.0},
				{ID: 5, Score: 100.0},
			},
			expected: []RankedScore{
				{Score: 500.0, Rank: 1},
				{Score: 400.0, Rank: 2},
				{Score: 300.0, Rank: 3},
				{Score: 200.0, Rank: 4},
				{Score: 100.0, Rank: 5},
			},
		},
		{
			name: "Mixed with many ties",
			scores: []Score{
				{ID: 1, Score: 100.0},
				{ID: 2, Score: 90.0},
				{ID: 3, Score: 90.0},
				{ID: 4, Score: 80.0},
				{ID: 5, Score: 80.0},
				{ID: 6, Score: 80.0},
				{ID: 7, Score: 70.0},
			},
			expected: []RankedScore{
				{Score: 100.0, Rank: 1},
				{Score: 90.0, Rank: 2},
				{Score: 90.0, Rank: 2},
				{Score: 80.0, Rank: 3},
				{Score: 80.0, Rank: 3},
				{Score: 80.0, Rank: 3},
				{Score: 70.0, Rank: 4},
			},
		},
		{
			name: "Single score",
			scores: []Score{
				{ID: 1, Score: 50.0},
			},
			expected: []RankedScore{
				{Score: 50.0, Rank: 1},
			},
		},
		{
			name:   "Empty scores list",
			scores: []Score{},
			expected: []RankedScore{},
		},
		{
			name: "Negative scores",
			scores: []Score{
				{ID: 1, Score: -10.0},
				{ID: 2, Score: -20.0},
				{ID: 3, Score: -20.0},
				{ID: 4, Score: -30.0},
			},
			expected: []RankedScore{
				{Score: -10.0, Rank: 1},
				{Score: -20.0, Rank: 2},
				{Score: -20.0, Rank: 2},
				{Score: -30.0, Rank: 3},
			},
		},
		{
			name: "Decimal scores with precision",
			scores: []Score{
				{ID: 1, Score: 3.14159},
				{ID: 2, Score: 3.14159},
				{ID: 3, Score: 2.71828},
				{ID: 4, Score: 2.71828},
				{ID: 5, Score: 1.61803},
			},
			expected: []RankedScore{
				{Score: 3.14159, Rank: 1},
				{Score: 3.14159, Rank: 1},
				{Score: 2.71828, Rank: 2},
				{Score: 2.71828, Rank: 2},
				{Score: 1.61803, Rank: 3},
			},
		},
		{
			name: "Large dataset with alternating ties",
			scores: []Score{
				{ID: 1, Score: 100.0},
				{ID: 2, Score: 100.0},
				{ID: 3, Score: 90.0},
				{ID: 4, Score: 90.0},
				{ID: 5, Score: 80.0},
				{ID: 6, Score: 80.0},
				{ID: 7, Score: 70.0},
				{ID: 8, Score: 70.0},
				{ID: 9, Score: 60.0},
				{ID: 10, Score: 60.0},
			},
			expected: []RankedScore{
				{Score: 100.0, Rank: 1},
				{Score: 100.0, Rank: 1},
				{Score: 90.0, Rank: 2},
				{Score: 90.0, Rank: 2},
				{Score: 80.0, Rank: 3},
				{Score: 80.0, Rank: 3},
				{Score: 70.0, Rank: 4},
				{Score: 70.0, Rank: 4},
				{Score: 60.0, Rank: 5},
				{Score: 60.0, Rank: 5},
			},
		},
		{
			name: "Unsorted input should be sorted by score descending",
			scores: []Score{
				{ID: 1, Score: 50.0},
				{ID: 2, Score: 100.0},
				{ID: 3, Score: 75.0},
				{ID: 4, Score: 100.0},
				{ID: 5, Score: 25.0},
			},
			expected: []RankedScore{
				{Score: 100.0, Rank: 1},
				{Score: 100.0, Rank: 1},
				{Score: 75.0, Rank: 2},
				{Score: 50.0, Rank: 3},
				{Score: 25.0, Rank: 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RankScores(tt.scores)
			assertRankedScores(t, result, tt.expected)
		})
	}
}

func TestRankScoresOptimized(t *testing.T) {
	tests := []struct {
		name   string
		scores []Score
		expected []RankedScore
	}{
		{
			name: "Basic example",
			scores: []Score{
				{ID: 1, Score: 100.0},
				{ID: 2, Score: 90.0},
				{ID: 3, Score: 90.0},
				{ID: 4, Score: 80.0},
			},
			expected: []RankedScore{
				{Score: 100.0, Rank: 1},
				{Score: 90.0, Rank: 2},
				{Score: 90.0, Rank: 2},
				{Score: 80.0, Rank: 3},
			},
		},
		{
			name: "Single score",
			scores: []Score{
				{ID: 1, Score: 50.0},
			},
			expected: []RankedScore{
				{Score: 50.0, Rank: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RankScoresOptimized(tt.scores)
			assertRankedScores(t, result, tt.expected)
		})
	}
}

func TestRankScoresWithMap(t *testing.T) {
	tests := []struct {
		name   string
		scores []Score
		expected []RankedScore
	}{
		{
			name: "Many duplicates",
			scores: []Score{
				{ID: 1, Score: 100.0},
				{ID: 2, Score: 100.0},
				{ID: 3, Score: 100.0},
				{ID: 4, Score: 90.0},
				{ID: 5, Score: 90.0},
				{ID: 6, Score: 80.0},
			},
			expected: []RankedScore{
				{Score: 100.0, Rank: 1},
				{Score: 100.0, Rank: 1},
				{Score: 100.0, Rank: 1},
				{Score: 90.0, Rank: 2},
				{Score: 90.0, Rank: 2},
				{Score: 80.0, Rank: 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RankScoresWithMap(tt.scores)
			assertRankedScores(t, result, tt.expected)
		})
	}
}

func TestRankScoresSQLStyle(t *testing.T) {
	tests := []struct {
		name   string
		scores []Score
		expected []RankedScore
	}{
		{
			name: "SQL style dense ranking",
			scores: []Score{
				{ID: 1, Score: 100.0},
				{ID: 2, Score: 100.0},
				{ID: 3, Score: 90.0},
				{ID: 4, Score: 80.0},
				{ID: 5, Score: 80.0},
			},
			expected: []RankedScore{
				{Score: 100.0, Rank: 1},
				{Score: 100.0, Rank: 1},
				{Score: 90.0, Rank: 2},
				{Score: 80.0, Rank: 3},
				{Score: 80.0, Rank: 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RankScoresSQLStyle(tt.scores)
			assertRankedScores(t, result, tt.expected)
		})
	}
}

func TestRankScoresInPlace(t *testing.T) {
	tests := []struct {
		name   string
		scores []Score
		expected []RankedScore
	}{
		{
			name: "In-place ranking",
			scores: []Score{
				{ID: 1, Score: 100.0},
				{ID: 2, Score: 90.0},
				{ID: 3, Score: 90.0},
				{ID: 4, Score: 80.0},
			},
			expected: []RankedScore{
				{Score: 100.0, Rank: 1},
				{Score: 90.0, Rank: 2},
				{Score: 90.0, Rank: 2},
				{Score: 80.0, Rank: 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RankScoresInPlace(tt.scores)
			assertRankedScores(t, result, tt.expected)
		})
	}
}

func TestRankScores_Consistency(t *testing.T) {
	// Test that all implementations produce the same results
	testCases := []struct {
		name   string
		scores []Score
	}{
		{
			name: "Normal case",
			scores: []Score{
				{ID: 1, Score: 100.0},
				{ID: 2, Score: 90.0},
				{ID: 3, Score: 90.0},
				{ID: 4, Score: 80.0},
				{ID: 5, Score: 80.0},
				{ID: 6, Score: 70.0},
			},
		},
		{
			name: "All same scores",
			scores: []Score{
				{ID: 1, Score: 50.0},
				{ID: 2, Score: 50.0},
				{ID: 3, Score: 50.0},
			},
		},
		{
			name: "Strictly decreasing",
			scores: []Score{
				{ID: 1, Score: 300.0},
				{ID: 2, Score: 200.0},
				{ID: 3, Score: 100.0},
			},
		},
		{
			name:   "Empty list",
			scores: []Score{},
		},
		{
			name: "Single score",
			scores: []Score{
				{ID: 1, Score: 99.9},
			},
		},
		{
			name: "Many duplicates",
			scores: []Score{
				{ID: 1, Score: 100.0},
				{ID: 2, Score: 100.0},
				{ID: 3, Score: 100.0},
				{ID: 4, Score: 90.0},
				{ID: 5, Score: 90.0},
				{ID: 6, Score: 80.0},
				{ID: 7, Score: 80.0},
				{ID: 8, Score: 80.0},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Make copies for implementations that modify input
			scores1 := make([]Score, len(tc.scores))
			copy(scores1, tc.scores)
			
			scores2 := make([]Score, len(tc.scores))
			copy(scores2, tc.scores)
			
			scores3 := make([]Score, len(tc.scores))
			copy(scores3, tc.scores)
			
			scores4 := make([]Score, len(tc.scores))
			copy(scores4, tc.scores)
			
			scores5 := make([]Score, len(tc.scores))
			copy(scores5, tc.scores)

			result1 := RankScores(scores1)
			result2 := RankScoresOptimized(scores2)
			result3 := RankScoresWithMap(scores3)
			result4 := RankScoresSQLStyle(scores4)
			result5 := RankScoresInPlace(scores5)

			// All should have same length
			assert.Equal(t, len(result1), len(result2), "Result lengths should match")
			assert.Equal(t, len(result1), len(result3), "Result lengths should match")
			assert.Equal(t, len(result1), len(result4), "Result lengths should match")
			assert.Equal(t, len(result1), len(result5), "Result lengths should match")

			// All should have same values
			for i := 0; i < len(result1); i++ {
				assert.Equal(t, result1[i].Score, result2[i].Score, "Score mismatch at index %d", i)
				assert.Equal(t, result1[i].Rank, result2[i].Rank, "Rank mismatch at index %d", i)
				
				assert.Equal(t, result1[i].Score, result3[i].Score, "Score mismatch at index %d", i)
				assert.Equal(t, result1[i].Rank, result3[i].Rank, "Rank mismatch at index %d", i)
				
				assert.Equal(t, result1[i].Score, result4[i].Score, "Score mismatch at index %d", i)
				assert.Equal(t, result1[i].Rank, result4[i].Rank, "Rank mismatch at index %d", i)
				
				assert.Equal(t, result1[i].Score, result5[i].Score, "Score mismatch at index %d", i)
				assert.Equal(t, result1[i].Rank, result5[i].Rank, "Rank mismatch at index %d", i)
			}
		})
	}
}

func BenchmarkRankScores(b *testing.B) {
	// Create test data with 10000 scores
	scores := make([]Score, 10000)
	for i := 0; i < 10000; i++ {
		scores[i] = Score{
			ID:    i + 1,
			Score: float64((i * 17) % 5000) / 100.0, // Generate varied decimal scores
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a copy for each iteration since implementations may modify input
		scoresCopy := make([]Score, len(scores))
		copy(scoresCopy, scores)
		RankScores(scoresCopy)
	}
}

func BenchmarkRankScoresOptimized(b *testing.B) {
	scores := make([]Score, 10000)
	for i := 0; i < 10000; i++ {
		scores[i] = Score{
			ID:    i + 1,
			Score: float64((i * 17) % 5000) / 100.0,
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		scoresCopy := make([]Score, len(scores))
		copy(scoresCopy, scores)
		RankScoresOptimized(scoresCopy)
	}
}

func BenchmarkRankScoresWithMap(b *testing.B) {
	scores := make([]Score, 10000)
	for i := 0; i < 10000; i++ {
		scores[i] = Score{
			ID:    i + 1,
			Score: float64((i * 17) % 5000) / 100.0,
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		scoresCopy := make([]Score, len(scores))
		copy(scoresCopy, scores)
		RankScoresWithMap(scoresCopy)
	}
}

func BenchmarkRankScoresSQLStyle(b *testing.B) {
	scores := make([]Score, 10000)
	for i := 0; i < 10000; i++ {
		scores[i] = Score{
			ID:    i + 1,
			Score: float64((i * 17) % 5000) / 100.0,
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		scoresCopy := make([]Score, len(scores))
		copy(scoresCopy, scores)
		RankScoresSQLStyle(scoresCopy)
	}
}

func BenchmarkRankScoresInPlace(b *testing.B) {
	scores := make([]Score, 10000)
	for i := 0; i < 10000; i++ {
		scores[i] = Score{
			ID:    i + 1,
			Score: float64((i * 17) % 5000) / 100.0,
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		scoresCopy := make([]Score, len(scores))
		copy(scoresCopy, scores)
		RankScoresInPlace(scoresCopy)
	}
}

// Helper function to assert RankedScore results
func assertRankedScores(t *testing.T, actual, expected []RankedScore) {
	assert.Equal(t, len(expected), len(actual), "Result length mismatch")

	// Check that results are sorted by score descending
	for i := 1; i < len(actual); i++ {
		assert.GreaterOrEqual(t, actual[i-1].Score, actual[i].Score, 
			"Scores should be sorted descending at index %d", i)
	}

	// Check each element
	for i, exp := range expected {
		if i >= len(actual) {
			break
		}
		act := actual[i]
		
		// Use epsilon for float comparison
		assert.InDelta(t, exp.Score, act.Score, 0.0001, "Score mismatch at index %d", i)
		assert.Equal(t, exp.Rank, act.Rank, "Rank mismatch at index %d", i)
	}

	// Additional validation: ranks should be dense (no gaps)
	if len(actual) > 0 {
		rankSet := make(map[int]bool)
		for _, rs := range actual {
			rankSet[rs.Rank] = true
		}
		
		// Check that ranks are consecutive from 1
		maxRank := 0
		for rank := range rankSet {
			if rank > maxRank {
				maxRank = rank
			}
		}
		
		for i := 1; i <= maxRank; i++ {
			assert.True(t, rankSet[i], "Rank %d should exist (dense ranking)", i)
		}
	}
}