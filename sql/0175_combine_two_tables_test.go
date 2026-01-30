package sql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombineTwoTables(t *testing.T) {
	tests := []struct {
		name      string
		persons   []Person
		addresses []Address
		expected  []PersonAddress
	}{
		{
			name: "Example 1 from LeetCode",
			persons: []Person{
				{PersonID: 1, LastName: "Wang", FirstName: "Allen"},
				{PersonID: 2, LastName: "Alice", FirstName: "Bob"},
			},
			addresses: []Address{
				{AddressID: 1, PersonID: 2, City: "New York City", State: "NY"},
				{AddressID: 2, PersonID: 3, City: "Leetcode", State: "CA"},
			},
			expected: []PersonAddress{
				{FirstName: "Allen", LastName: "Wang", City: nil, State: nil},
				{FirstName: "Bob", LastName: "Alice", City: strPtr("New York City"), State: strPtr("NY")},
			},
		},
		{
			name: "All persons have addresses",
			persons: []Person{
				{PersonID: 1, LastName: "Doe", FirstName: "John"},
				{PersonID: 2, LastName: "Smith", FirstName: "Jane"},
				{PersonID: 3, LastName: "Johnson", FirstName: "Bob"},
			},
			addresses: []Address{
				{AddressID: 1, PersonID: 1, City: "New York", State: "NY"},
				{AddressID: 2, PersonID: 2, City: "Los Angeles", State: "CA"},
				{AddressID: 3, PersonID: 3, City: "Chicago", State: "IL"},
			},
			expected: []PersonAddress{
				{FirstName: "John", LastName: "Doe", City: strPtr("New York"), State: strPtr("NY")},
				{FirstName: "Jane", LastName: "Smith", City: strPtr("Los Angeles"), State: strPtr("CA")},
				{FirstName: "Bob", LastName: "Johnson", City: strPtr("Chicago"), State: strPtr("IL")},
			},
		},
		{
			name: "No persons have addresses",
			persons: []Person{
				{PersonID: 1, LastName: "A", FirstName: "Person"},
				{PersonID: 2, LastName: "B", FirstName: "Another"},
			},
			addresses: []Address{},
			expected: []PersonAddress{
				{FirstName: "Person", LastName: "A", City: nil, State: nil},
				{FirstName: "Another", LastName: "B", City: nil, State: nil},
			},
		},
		{
			name:    "Empty persons list",
			persons: []Person{},
			addresses: []Address{
				{AddressID: 1, PersonID: 1, City: "Somewhere", State: "ST"},
			},
			expected: []PersonAddress{},
		},
		{
			name: "Multiple addresses for same person (should use first found)",
			persons: []Person{
				{PersonID: 1, LastName: "Test", FirstName: "Person"},
			},
			addresses: []Address{
				{AddressID: 1, PersonID: 1, City: "First City", State: "FC"},
				{AddressID: 2, PersonID: 1, City: "Second City", State: "SC"},
			},
			expected: []PersonAddress{
				{FirstName: "Person", LastName: "Test", City: strPtr("First City"), State: strPtr("FC")},
			},
		},
		{
			name: "Mixed: some have addresses, some don't",
			persons: []Person{
				{PersonID: 1, LastName: "HasAddress", FirstName: "Person"},
				{PersonID: 2, LastName: "NoAddress", FirstName: "Person"},
				{PersonID: 3, LastName: "HasAddressToo", FirstName: "Person"},
			},
			addresses: []Address{
				{AddressID: 1, PersonID: 1, City: "City1", State: "S1"},
				{AddressID: 2, PersonID: 3, City: "City3", State: "S3"},
			},
			expected: []PersonAddress{
				{FirstName: "Person", LastName: "HasAddress", City: strPtr("City1"), State: strPtr("S1")},
				{FirstName: "Person", LastName: "NoAddress", City: nil, State: nil},
				{FirstName: "Person", LastName: "HasAddressToo", City: strPtr("City3"), State: strPtr("S3")},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CombineTwoTables(tt.persons, tt.addresses)
			assert.Equal(t, len(tt.expected), len(result), "Result length mismatch")

			for i, expected := range tt.expected {
				if i >= len(result) {
					break
				}
				actual := result[i]

				assert.Equal(t, expected.FirstName, actual.FirstName, "FirstName mismatch at index %d", i)
				assert.Equal(t, expected.LastName, actual.LastName, "LastName mismatch at index %d", i)

				// Compare City (nil-aware)
				if expected.City == nil {
					assert.Nil(t, actual.City, "City should be nil at index %d", i)
				} else {
					assert.NotNil(t, actual.City, "City should not be nil at index %d", i)
					if actual.City != nil {
						assert.Equal(t, *expected.City, *actual.City, "City value mismatch at index %d", i)
					}
				}

				// Compare State (nil-aware)
				if expected.State == nil {
					assert.Nil(t, actual.State, "State should be nil at index %d", i)
				} else {
					assert.NotNil(t, actual.State, "State should not be nil at index %d", i)
					if actual.State != nil {
						assert.Equal(t, *expected.State, *actual.State, "State value mismatch at index %d", i)
					}
				}
			}
		})
	}
}

func TestCombineTwoTablesSQLStyle(t *testing.T) {
	tests := []struct {
		name      string
		persons   []Person
		addresses []Address
		expected  []PersonAddress
	}{
		{
			name: "Basic test",
			persons: []Person{
				{PersonID: 1, LastName: "Test", FirstName: "Person"},
			},
			addresses: []Address{
				{AddressID: 1, PersonID: 1, City: "Test City", State: "TC"},
			},
			expected: []PersonAddress{
				{FirstName: "Person", LastName: "Test", City: strPtr("Test City"), State: strPtr("TC")},
			},
		},
		{
			name: "Person without address",
			persons: []Person{
				{PersonID: 1, LastName: "NoAddr", FirstName: "Person"},
			},
			addresses: []Address{},
			expected: []PersonAddress{
				{FirstName: "Person", LastName: "NoAddr", City: nil, State: nil},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CombineTwoTablesSQLStyle(tt.persons, tt.addresses)
			assert.Equal(t, len(tt.expected), len(result), "Result length mismatch")

			for i, expected := range tt.expected {
				if i >= len(result) {
					break
				}
				actual := result[i]

				assert.Equal(t, expected.FirstName, actual.FirstName, "FirstName mismatch at index %d")
				assert.Equal(t, expected.LastName, actual.LastName, "LastName mismatch at index %d")

				// Compare City
				if expected.City == nil {
					assert.Nil(t, actual.City, "City should be nil at index %d")
				} else if actual.City != nil {
					assert.Equal(t, *expected.City, *actual.City, "City value mismatch at index %d")
				}

				// Compare State
				if expected.State == nil {
					assert.Nil(t, actual.State, "State should be nil at index %d")
				} else if actual.State != nil {
					assert.Equal(t, *expected.State, *actual.State, "State value mismatch at index %d")
				}
			}
		})
	}
}

func TestCombineTwoTables_Consistency(t *testing.T) {
	// Test that both implementations produce the same results
	persons := []Person{
		{PersonID: 1, LastName: "A", FirstName: "First"},
		{PersonID: 2, LastName: "B", FirstName: "Second"},
		{PersonID: 3, LastName: "C", FirstName: "Third"},
	}
	addresses := []Address{
		{AddressID: 1, PersonID: 1, City: "City1", State: "S1"},
		{AddressID: 2, PersonID: 3, City: "City3", State: "S3"},
	}

	result1 := CombineTwoTables(persons, addresses)
	result2 := CombineTwoTablesSQLStyle(persons, addresses)

	assert.Equal(t, len(result1), len(result2), "Results should have same length")

	for i := 0; i < len(result1); i++ {
		r1 := result1[i]
		r2 := result2[i]

		assert.Equal(t, r1.FirstName, r2.FirstName, "FirstName mismatch at index %d", i)
		assert.Equal(t, r1.LastName, r2.LastName, "LastName mismatch at index %d", i)

		// Compare City
		if r1.City == nil {
			assert.Nil(t, r2.City, "City should both be nil at index %d", i)
		} else if r2.City != nil {
			assert.Equal(t, *r1.City, *r2.City, "City value mismatch at index %d", i)
		}

		// Compare State
		if r1.State == nil {
			assert.Nil(t, r2.State, "State should both be nil at index %d", i)
		} else if r2.State != nil {
			assert.Equal(t, *r1.State, *r2.State, "State value mismatch at index %d", i)
		}
	}
}

func BenchmarkCombineTwoTables(b *testing.B) {
	// Create test data
	persons := make([]Person, 1000)
	addresses := make([]Address, 800) // 80% of persons have addresses

	for i := 0; i < 1000; i++ {
		persons[i] = Person{
			PersonID:  i + 1,
			LastName:  fmt.Sprintf("LastName%d", i),
			FirstName: fmt.Sprintf("FirstName%d", i),
		}
	}

	for i := 0; i < 800; i++ {
		addresses[i] = Address{
			AddressID: i + 1,
			PersonID:  i + 1, // First 800 persons have addresses
			City:      fmt.Sprintf("City%d", i),
			State:     fmt.Sprintf("State%d", i),
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CombineTwoTables(persons, addresses)
	}
}

func BenchmarkCombineTwoTablesSQLStyle(b *testing.B) {
	// Create test data
	persons := make([]Person, 100)
	addresses := make([]Address, 80) // 80% of persons have addresses

	for i := 0; i < 100; i++ {
		persons[i] = Person{
			PersonID:  i + 1,
			LastName:  fmt.Sprintf("LastName%d", i),
			FirstName: fmt.Sprintf("FirstName%d", i),
		}
	}

	for i := 0; i < 80; i++ {
		addresses[i] = Address{
			AddressID: i + 1,
			PersonID:  i + 1, // First 80 persons have addresses
			City:      fmt.Sprintf("City%d", i),
			State:     fmt.Sprintf("State%d", i),
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CombineTwoTablesSQLStyle(persons, addresses)
	}
}

// Helper function to create string pointers
func strPtr(s string) *string {
	return &s
}