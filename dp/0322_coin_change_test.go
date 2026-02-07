package dp

import (
	"testing"
)

func TestCoinChangeDP(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "Example 1: Standard case",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3, // 5 + 5 + 1
		},
		{
			name:   "Example 2: No solution",
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			name:   "Example 3: Zero amount",
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
		{
			name:   "Single coin exact match",
			coins:  []int{5},
			amount: 5,
			want:   1,
		},
		{
			name:   "Multiple coins, optimal uses larger denominations",
			coins:  []int{1, 3, 4},
			amount: 6,
			want:   2, // 3 + 3 (not 4 + 1 + 1)
		},
		{
			name:   "Large amount with small coins",
			coins:  []int{1, 2, 5},
			amount: 100,
			want:   20, // 20 * 5
		},
		{
			name:   "Unsorted coins",
			coins:  []int{5, 2, 1},
			amount: 11,
			want:   3,
		},
		{
			name:   "Duplicate coins in input",
			coins:  []int{1, 2, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			name:   "Amount smaller than smallest coin",
			coins:  []int{3, 7, 10},
			amount: 2,
			want:   -1,
		},
		{
			name:   "Greedy approach would fail",
			coins:  []int{1, 3, 4, 5},
			amount: 7,
			want:   2, // 3 + 4 (greedy would take 5 + 1 + 1 = 3 coins)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coinChangeDP(tt.coins, tt.amount)
			if got != tt.want {
				t.Errorf("coinChangeDP(%v, %d) = %d, want %d", tt.coins, tt.amount, got, tt.want)
			}
		})
	}
}

func TestCoinChangeBFS(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "Standard case",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			name:   "No solution",
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			name:   "Zero amount",
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
		{
			name:   "Single coin exact match",
			coins:  []int{5},
			amount: 5,
			want:   1,
		},
		{
			name:   "Large amount",
			coins:  []int{1, 2, 5},
			amount: 100,
			want:   20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coinChangeBFS(tt.coins, tt.amount)
			if got != tt.want {
				t.Errorf("coinChangeBFS(%v, %d) = %d, want %d", tt.coins, tt.amount, got, tt.want)
			}
		})
	}
}

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "Main function test 1",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			name:   "Main function test 2",
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			name:   "Main function test 3",
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CoinChange(tt.coins, tt.amount)
			if got != tt.want {
				t.Errorf("CoinChange(%v, %d) = %d, want %d", tt.coins, tt.amount, got, tt.want)
			}
		})
	}
}

func BenchmarkCoinChangeDP(b *testing.B) {
	coins := []int{1, 2, 5, 10, 20, 50, 100}
	amount := 1000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coinChangeDP(coins, amount)
	}
}

func BenchmarkCoinChangeBFS(b *testing.B) {
	coins := []int{1, 2, 5, 10, 20, 50, 100}
	amount := 1000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coinChangeBFS(coins, amount)
	}
}

func BenchmarkCoinChangeLargeAmount(b *testing.B) {
	coins := []int{1, 2, 5, 10, 25}
	amount := 10000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coinChangeDP(coins, amount)
	}
}