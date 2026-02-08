package dp

import "testing"

func TestClimbStairs_0194(t *testing.T) {
	// Test case 1: n = 1
	result := climbStairs(1)
	expected := 1
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	// Test case 2: n = 2
	result = climbStairs(2)
	expected = 2
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	// Test case 3: n = 3
	result = climbStairs(3)
	expected = 3
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	// Test case 4: n = 4
	result = climbStairs(4)
	expected = 5
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
