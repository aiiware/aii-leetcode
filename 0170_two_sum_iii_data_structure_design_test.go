package leetcode

import (
    "testing"
)

func TestTwoSumIII(t *testing.T) {
    tests := []struct {
        name     string
        operations []string
        inputs    [][]int
        expected  []interface{}
    }{
        {
            name: "Example 1",
            operations: []string{"TwoSumIII", "add", "add", "add", "find", "find"},
            inputs: [][]int{{}, {1}, {3}, {5}, {4}, {7}},
            expected: []interface{}{nil, nil, nil, nil, true, false},
        },
        {
            name: "Multiple same numbers",
            operations: []string{"TwoSumIII", "add", "add", "add", "find", "find", "find"},
            inputs: [][]int{{}, {1}, {1}, {2}, {2}, {4}, {3}},
            expected: []interface{}{nil, nil, nil, nil, true, false, true}, // 1+1=2, no pair sums to 4, 1+2=3
        },
        {
            name: "Negative numbers",
            operations: []string{"TwoSumIII", "add", "add", "add", "find", "find", "find"},
            inputs: [][]int{{}, {-1}, {2}, {-3}, {1}, {-4}, {-1}},
            expected: []interface{}{nil, nil, nil, nil, true, true, true}, // -1+2=1, -1+-3=-4, 2+-3=-1
        },
        {
            name: "Empty structure",
            operations: []string{"TwoSumIII", "find"},
            inputs: [][]int{{}, {0}},
            expected: []interface{}{nil, false},
        },
        {
            name: "Large numbers",
            operations: []string{"TwoSumIII", "add", "add", "find", "find"},
            inputs: [][]int{{}, {100000}, {-100000}, {0}, {200000}},
            expected: []interface{}{nil, nil, nil, true, false},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var obj *TwoSumIII
            
            for i, op := range tt.operations {
                switch op {
                case "TwoSumIII":
                    ts := NewTwoSumIII()
                    obj = &ts
                case "add":
                    obj.Add(tt.inputs[i][0])
                case "find":
                    result := obj.Find(tt.inputs[i][0])
                    expected := tt.expected[i].(bool)
                    if result != expected {
                        t.Errorf("Find(%v) = %v, want %v", tt.inputs[i][0], result, expected)
                    }
                }
            }
        })
    }
}

func BenchmarkTwoSumIIIAdd(b *testing.B) {
    obj := NewTwoSumIII()
    for i := 0; i < b.N; i++ {
        obj.Add(i % 1000)
    }
}

func BenchmarkTwoSumIIIFind(b *testing.B) {
    obj := NewTwoSumIII()
    // Add some numbers first
    for i := 0; i < 1000; i++ {
        obj.Add(i)
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        obj.Find(i % 2000)
    }
}