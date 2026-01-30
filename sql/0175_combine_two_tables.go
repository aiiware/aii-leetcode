package sql

/*
0175. Combine Two Tables

Table: Person
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| personId    | int     |
| lastName    | varchar |
| firstName   | varchar |
+-------------+---------+
personId is the primary key column for this table.
This table contains information about the ID of some persons and their first and last names.

Table: Address
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| addressId   | int     |
| personId    | int     |
| city        | varchar |
| state       | varchar |
+-------------+---------+
addressId is the primary key column for this table.
Each row of this table contains information about the city and state of one person with ID = PersonId.

Write an SQL query to report the first name, last name, city, and state of each person in the Person table.
If the address of a personId is not present in the Address table, report null instead.

Return the result table in any order.

Example 1:
Input: 
Person table:
+----------+----------+-----------+
| personId | lastName | firstName |
+----------+----------+-----------+
| 1        | Wang     | Allen     |
| 2        | Alice    | Bob       |
+----------+----------+-----------+
Address table:
+-----------+----------+---------------+-------+
| addressId | personId | city          | state |
+-----------+----------+---------------+-------+
| 1         | 2        | New York City | NY    |
| 2         | 3        | Leetcode      | CA    |
+-----------+----------+---------------+-------+
Output: 
+-----------+----------+---------------+-------+
| firstName | lastName | city          | state |
+-----------+----------+---------------+-------+
| Allen     | Wang     | null          | null  |
| Bob       | Alice    | New York City | NY    |
+-----------+----------+---------------+-------+
Explanation: 
There is no address in the address table for the personId = 1 so we return null in their city and state.
addressId = 1 contains information about the address of personId = 2.

Difficulty: Easy
Tags: Database
Companies: Amazon, Apple, Google, Microsoft
*/

// Person represents a person in the Person table
type Person struct {
	PersonID  int
	LastName  string
	FirstName string
}

// Address represents an address in the Address table
type Address struct {
	AddressID int
	PersonID  int
	City      string
	State     string
}

// PersonAddress represents the combined result of Person and Address tables
type PersonAddress struct {
	FirstName string
	LastName  string
	City      *string // Use pointer to allow null
	State     *string // Use pointer to allow null
}

// CombineTwoTables simulates a LEFT JOIN between Person and Address tables
// It returns firstName, lastName, city, and state for each person
// If a person doesn't have an address, city and state will be nil
// If a person has multiple addresses, it uses the first one found
func CombineTwoTables(persons []Person, addresses []Address) []PersonAddress {
	// Create a map for quick address lookup by personId
	addressMap := make(map[int]Address)
	for _, addr := range addresses {
		// Only add if not already present (use first found)
		if _, exists := addressMap[addr.PersonID]; !exists {
			addressMap[addr.PersonID] = addr
		}
	}

	// Perform the "LEFT JOIN" operation
	var result []PersonAddress
	for _, person := range persons {
		pa := PersonAddress{
			FirstName: person.FirstName,
			LastName:  person.LastName,
		}

		// Check if person has an address
		if addr, found := addressMap[person.PersonID]; found {
			pa.City = &addr.City
			pa.State = &addr.State
		} else {
			// No address found, leave City and State as nil (null in SQL)
			pa.City = nil
			pa.State = nil
		}

		result = append(result, pa)
	}

	return result
}

// CombineTwoTablesSQLStyle provides an alternative implementation that more closely
// mimics the SQL query pattern using nested loops (like a hash join)
// If a person has multiple addresses, it uses the first one found
func CombineTwoTablesSQLStyle(persons []Person, addresses []Address) []PersonAddress {
	var result []PersonAddress

	// For each person, find matching address (if any)
	for _, person := range persons {
		pa := PersonAddress{
			FirstName: person.FirstName,
			LastName:  person.LastName,
		}

		// Find address for this person (use first found)
		found := false
		for _, addr := range addresses {
			if addr.PersonID == person.PersonID {
				pa.City = &addr.City
				pa.State = &addr.State
				found = true
				break // Use first found
			}
		}

		if !found {
			pa.City = nil
			pa.State = nil
		}

		result = append(result, pa)
	}

	return result
}