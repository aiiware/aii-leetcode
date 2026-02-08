package arrays

import (
    "testing"
)

func TestLongestPalindromicSubstring(t *testing.T) {
    tests := []struct {
        name string
        // Add your test parameters here
        want int
    }{
        {
            name: "Example 1",
            // Add test parameters here
            want: 0,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // result := LongestPalindromicSubstring()
            // assert.Equal(t, tt.want, result)
            t.Skip("Not implemented")
        })
    }
}

func BenchmarkLongestPalindromicSubstring(b *testing.B) {
    // TODO: Implement benchmark
    b.Skip("Not implemented")
}