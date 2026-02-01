package design

import (
	"testing"
)

func TestMedianFinder(t *testing.T) {
	tests := []struct {
		name     string
		actions  []string
		inputs   []interface{}
		expected []interface{}
	}{
		{
			name:     "Example 1",
			actions:  []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"},
			inputs:   []interface{}{nil, 1, 2, nil, 3, nil},
			expected: []interface{}{nil, nil, nil, 1.5, nil, 2.0},
		},
		{
			name:     "Single element",
			actions:  []string{"MedianFinder", "addNum", "findMedian"},
			inputs:   []interface{}{nil, 5, nil},
			expected: []interface{}{nil, nil, 5.0},
		},
		{
			name:     "Two elements",
			actions:  []string{"MedianFinder", "addNum", "addNum", "findMedian"},
			inputs:   []interface{}{nil, 1, 3, nil},
			expected: []interface{}{nil, nil, nil, 2.0},
		},
		{
			name:     "Multiple elements odd",
			actions:  []string{"MedianFinder", "addNum", "addNum", "addNum", "addNum", "addNum", "findMedian"},
			inputs:   []interface{}{nil, 1, 2, 3, 4, 5, nil},
			expected: []interface{}{nil, nil, nil, nil, nil, nil, 3.0},
		},
		{
			name:     "Multiple elements even",
			actions:  []string{"MedianFinder", "addNum", "addNum", "addNum", "addNum", "findMedian"},
			inputs:   []interface{}{nil, 1, 2, 3, 4, nil},
			expected: []interface{}{nil, nil, nil, nil, nil, 2.5},
		},
		{
			name:     "Negative numbers",
			actions:  []string{"MedianFinder", "addNum", "addNum", "addNum", "findMedian"},
			inputs:   []interface{}{nil, -1, -2, -3, nil},
			expected: []interface{}{nil, nil, nil, nil, -2.0},
		},
		{
			name:     "Mixed positive and negative",
			actions:  []string{"MedianFinder", "addNum", "addNum", "addNum", "findMedian"},
			inputs:   []interface{}{nil, -1, 2, -3, nil},
			expected: []interface{}{nil, nil, nil, nil, -1.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mf MedianFinder
			for i, action := range tt.actions {
				switch action {
				case "MedianFinder":
					mf = ConstructorMedianFinder()
				case "addNum":
					mf.AddNum(tt.inputs[i].(int))
				case "findMedian":
					got := mf.FindMedian()
					expected := tt.expected[i].(float64)
					if got != expected {
						t.Errorf("FindMedian() = %v, want %v", got, expected)
					}
				}
			}
		})
	}
}

func TestMedianFinderAlt(t *testing.T) {
	tests := []struct {
		name     string
		actions  []string
		inputs   []interface{}
		expected []interface{}
	}{
		{
			name:     "Example 1",
			actions:  []string{"MedianFinderAlt", "addNum", "addNum", "findMedian", "addNum", "findMedian"},
			inputs:   []interface{}{nil, 1, 2, nil, 3, nil},
			expected: []interface{}{nil, nil, nil, 1.5, nil, 2.0},
		},
		{
			name:     "Single element",
			actions:  []string{"MedianFinderAlt", "addNum", "findMedian"},
			inputs:   []interface{}{nil, 5, nil},
			expected: []interface{}{nil, nil, 5.0},
		},
		{
			name:     "Two elements",
			actions:  []string{"MedianFinderAlt", "addNum", "addNum", "findMedian"},
			inputs:   []interface{}{nil, 1, 3, nil},
			expected: []interface{}{nil, nil, nil, 2.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mf MedianFinderAlt
			for i, action := range tt.actions {
				switch action {
				case "MedianFinderAlt":
					mf = ConstructorMedianFinderAlt()
				case "addNum":
					mf.AddNum(tt.inputs[i].(int))
				case "findMedian":
					got := mf.FindMedian()
					expected := tt.expected[i].(float64)
					if got != expected {
						t.Errorf("FindMedian() = %v, want %v", got, expected)
					}
				}
			}
		})
	}
}

func BenchmarkMedianFinder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mf := ConstructorMedianFinder()
		for j := 0; j < 1000; j++ {
			mf.AddNum(j)
			if j%100 == 0 {
				_ = mf.FindMedian()
			}
		}
	}
}

func BenchmarkMedianFinderAlt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mf := ConstructorMedianFinderAlt()
		for j := 0; j < 1000; j++ {
			mf.AddNum(j)
			if j%100 == 0 {
				_ = mf.FindMedian()
			}
		}
	}
}