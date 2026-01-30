package sql

/*
0178. Rank Scores

Table: Scores
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| score       | decimal |
+-------------+---------+
id is the primary key column for this table.
Each row of this table contains the score of a game.
Score is a floating point value with two decimal places.

Write an SQL query to rank the scores. The ranking should be calculated according to the following rules:
1. The scores should be ranked from highest to lowest.
2. If there is a tie between two scores, both should have the same ranking.
3. After a tie, the next ranking number should be the next consecutive integer value.
   In other words, there should be no "holes" between ranks.

Return the result table ordered by score in descending order.

Example 1:
Input: 
Scores table:
+----+-------+
| id | score |
+----+-------+
| 1  | 3.50  |
| 2  | 3.65  |
| 3  | 4.00  |
| 4  | 3.85  |
| 5  | 4.00  |
| 6  | 3.65  |
+----+-------+
Output: 
+-------+------+
| score | rank |
+-------+------+
| 4.00  | 1    |
| 4.00  | 1    |
| 3.85  | 2    |
| 3.65  | 3    |
| 3.65  | 3    |
| 3.50  | 4    |
+-------+------+

Constraints:
- 1 <= Scores.length <= 10^4
- -10^4 <= score <= 10^4

Difficulty: Medium
Tags: Database
Companies: Amazon, Google, Microsoft, Facebook, Apple
*/

import (
	"sort"
)

// Score represents a score in the Scores table
type Score struct {
	ID    int
	Score float64
}

// RankedScore represents a score with its rank
type RankedScore struct {
	Score float64
	Rank  int
}

// RankScores ranks scores according to the specified rules
// Returns scores sorted by score descending with dense ranking
func RankScores(scores []Score) []RankedScore {
	if len(scores) == 0 {
		return []RankedScore{}
	}

	// Create a copy to avoid modifying input
	scoresCopy := make([]Score, len(scores))
	copy(scoresCopy, scores)

	// Sort by score descending
	sort.Slice(scoresCopy, func(i, j int) bool {
		return scoresCopy[i].Score > scoresCopy[j].Score
	})

	// Apply dense ranking
	result := make([]RankedScore, len(scoresCopy))
	currentRank := 1
	currentScore := scoresCopy[0].Score

	for i, score := range scoresCopy {
		if i == 0 {
			// First element gets rank 1
			result[i] = RankedScore{Score: score.Score, Rank: currentRank}
			continue
		}

		// If score is different from previous, increment rank
		if score.Score != currentScore {
			currentRank++
			currentScore = score.Score
		}

		// Same rank for ties
		result[i] = RankedScore{Score: score.Score, Rank: currentRank}
	}

	return result
}

// RankScoresOptimized provides an optimized version with single pass after sorting
func RankScoresOptimized(scores []Score) []RankedScore {
	if len(scores) == 0 {
		return []RankedScore{}
	}

	// Sort by score descending
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	result := make([]RankedScore, len(scores))
	rank := 1
	prevScore := scores[0].Score

	for i, score := range scores {
		if i > 0 && score.Score != prevScore {
			rank++
			prevScore = score.Score
		}
		result[i] = RankedScore{Score: score.Score, Rank: rank}
	}

	return result
}

// RankScoresWithMap uses a map to track unique scores and their ranks
// More memory efficient for large datasets with many duplicates
func RankScoresWithMap(scores []Score) []RankedScore {
	if len(scores) == 0 {
		return []RankedScore{}
	}

	// Get unique scores
	scoreSet := make(map[float64]bool)
	for _, s := range scores {
		scoreSet[s.Score] = true
	}

	// Convert to slice and sort descending
	uniqueScores := make([]float64, 0, len(scoreSet))
	for score := range scoreSet {
		uniqueScores = append(uniqueScores, score)
	}
	sort.Slice(uniqueScores, func(i, j int) bool {
		return uniqueScores[i] > uniqueScores[j]
	})

	// Create map from score to rank
	rankMap := make(map[float64]int)
	for i, score := range uniqueScores {
		rankMap[score] = i + 1 // 1-indexed ranks
	}

	// Sort original scores by score descending
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	// Create result with ranks
	result := make([]RankedScore, len(scores))
	for i, score := range scores {
		result[i] = RankedScore{
			Score: score.Score,
			Rank:  rankMap[score.Score],
		}
	}

	return result
}

// RankScoresSQLStyle simulates the SQL approach using window functions
// This mimics the RANK() or DENSE_RANK() window function behavior
func RankScoresSQLStyle(scores []Score) []RankedScore {
	if len(scores) == 0 {
		return []RankedScore{}
	}

	// Sort by score descending
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	result := make([]RankedScore, len(scores))
	denseRank := 1
	prevScore := scores[0].Score

	// First pass: assign dense ranks
	for i, score := range scores {
		if i > 0 && score.Score != prevScore {
			denseRank++
			prevScore = score.Score
		}
		result[i] = RankedScore{Score: score.Score, Rank: denseRank}
	}

	return result
}

// RankScoresInPlace performs ranking without creating new slices
// Most memory efficient version
func RankScoresInPlace(scores []Score) []RankedScore {
	if len(scores) == 0 {
		return []RankedScore{}
	}

	// Sort by score descending
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	result := make([]RankedScore, len(scores))
	rank := 1
	currentScore := scores[0].Score

	for i := 0; i < len(scores); i++ {
		if i > 0 && scores[i].Score != currentScore {
			rank++
			currentScore = scores[i].Score
		}
		result[i] = RankedScore{Score: scores[i].Score, Rank: rank}
	}

	return result
}