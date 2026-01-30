package sql


/*
183. Customers Who Never Order

Table: Customers

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| name        | varchar |
+-------------+---------+
id is the primary key for this table.
Each row of this table indicates the ID and name of a customer.

Table: Orders

+-------------+------+
| Column Name | Type |
+-------------+------+
| id          | int  |
| customerId  | int  |
+-------------+------+
id is the primary key for this table.
customerId is a foreign key of the ID from the Customers table.
Each row of this table indicates the ID of an order and the ID of the customer who ordered it.

Write an SQL query to report all customers who never order anything.

Return the result table in any order.

Example 1:
Input: 
Customers table:
+----+-------+
| id | name  |
+----+-------+
| 1  | Joe   |
| 2  | Henry |
| 3  | Sam   |
| 4  | Max   |
+----+-------+
Orders table:
+----+------------+
| id | customerId |
+----+------------+
| 1  | 3          |
| 2  | 1          |
+----+------------+
Output: 
+-----------+
| Customers |
+-----------+
| Henry     |
| Max       |
+-----------+

Difficulty: Easy
Tags: Database
Companies: Amazon, Google, Microsoft
*/

// Customer183 represents a row in the Customers table for problem 183
type Customer183 struct {
	ID   int
	Name string
}

// Order183 represents a row in the Orders table for problem 183
type Order183 struct {
	ID         int
	CustomerID int
}

// CustomersWhoNeverOrder finds customers who have never placed an order.
// This is the Go equivalent of the SQL solution for LeetCode 0183.
//
// Algorithm:
// 1. Create a set of customer IDs who have placed orders
// 2. Iterate through all customers and check if their ID is in the set
// 3. If not, add their name to the result
//
// Time complexity: O(n + m) where n is number of customers, m is number of orders
// Space complexity: O(m) for the set of customer IDs with orders
func CustomersWhoNeverOrder(customers []Customer183, orders []Order183) []string {
	// Create a set of customer IDs who have placed orders
	customersWithOrders := make(map[int]bool)
	for _, order := range orders {
		customersWithOrders[order.CustomerID] = true
	}

	// Find customers who have never placed an order
	result := []string{}
	for _, customer := range customers {
		if !customersWithOrders[customer.ID] {
			result = append(result, customer.Name)
		}
	}

	return result
}