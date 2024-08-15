package main

import (
	"fmt"
	"sort"
)

/*
Merge Join

Merge Join is a join algorithm that works by first sorting the two tables to be joined on the join key. Once the tables are sorted, the join key is compared for each row in the two tables. If the join key matches, the rows are joined and returned as the result.

1. Sort the two tables to be joined on the join key.
2. Initialize two pointers, one for each table.
3. Compare the join key values:
	a. If the join key values match:
		i. join the rows
		ii. store the result in the resultset
		iii. increment the right table pointers to check for more matches (since multiple rows in the right table may have the same key).
	b. If the join key value from the left table is less than the right table, increment the left pointer.
	c. If the join key value from the right table is less than the left table, increment the right pointer.
4. Repeat this process until all rows are joined.

Time Complexity:
1. Both tables need to be sorted. Assuming n & m are the number of rows in the two tables, the sorting step takes O(n log n) and O(m log m) time.
2. The merge join step takes O(n + m) time.
3. Overall, the time complexity is O(n log n + m log m + n + m) = O(n log n + m log m).

Space Complexity: O(n+m)

Pseudocode:
1. Sort the left and right tables on the join key.
2. Initialize pointers i and j for the left and right tables.
3. While i < len(left) and j < len(right):
   a. If left[i].key == right[j].key:
      - Append the joined result to the output.
      - Increment j to process the next row in the right table.
      - If right[j].key still matches left[i].key, repeat this step.
   b. If left[i].key < right[j].key:
      - Increment i to process the next row in the left table.
   c. If right[j].key < left[i].key:
      - Increment j to process the next row in the right table.
4. Repeat until all rows are processed.


Example:
In this example, we will create two tables, `employees` and `departments`, and join them on the `department_id` column using the Merge Join algorithm.

Here, I will be using strucuts to represent the employees and departments tables.
*/

type Employee struct {
	id            int
	name          string
	department_id int
}

type Department struct {
	id   int
	name string
}

type Result struct {
	EmployeeID     int
	EmployeeName   string
	DepartmentID   int
	DepartmentName string
}

// Implementing sort.Interface for the Employee struct by "department_id". This is required to sort the Employee slice by the "department_id" field.
// It requires us to implement the following methods: Len, Less, and Swap.
// Internally, golang uses pdqsort for sorting. PDQSort stands for Pattern-Defeating Quicksort. It is a hybrid sorting algorithm that combines quicksort and insertion sort.
type Employees []Employee

// Len returns the length of the Employees slice.
func (e Employees) Len() int {
	return len(e)
}

// Less(i, j) returns true if the department_id of the Employee at index i is less than the department_id of the Employee at index j.
func (e Employees) Less(i, j int) bool {
	return e[i].department_id < e[j].department_id
}

// Swap swaps the Employees at index i and j.
func (e Employees) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// Implementing sort.Interface for the Department struct by "id". This is required to sort the Department slice by the "id" field.
type Departments []Department

// Len returns the length of the Departments slice.
func (d Departments) Len() int {
	return len(d)
}

// Less(i, j) returns true if the id of the Department at index i is less than the id of the Department at index j.
func (d Departments) Less(i, j int) bool {
	return d[i].id < d[j].id
}

// Swap swaps the Departments at index i and j.
func (d Departments) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// MergeJoin performs a merge join on the employees and departments tables based on the department_id & returns the joined result.
// The result contains the EmployeeID, EmployeeName, DepartmentID, and DepartmentName.
func MergeJoin(employees Employees, departments Departments) []Result {

	// Sort the employees and departments tables by the join key. department_id for employees and id for departments.
	sort.Sort(Employees(employees))
	sort.Sort(Departments(departments))

	// Initialize the result slice to store the joined rows.
	var result []Result

	// Initialize the pointers for the employees and departments tables.
	i, j := 0, 0

	// NOTE: Databases perform an optimisation where it prefers to have the smaller table on the left side of the join. This is because the smaller table can be loaded into memory and used for lookups during the join operation.

	// Perform the merge join operation. This is the merge step of the merge sort algorithm.
	// This is an inner join operation where we only include rows where the join key matches.
	// Time Complexity: O(n) where n is the number of rows in the employees and departments tables.
	for i < departments.Len() && j < employees.Len() {
		if departments[i].id == employees[j].department_id {
			result = append(result, Result{
				EmployeeID:     employees[j].id,
				EmployeeName:   employees[j].name,
				DepartmentID:   employees[j].department_id,
				DepartmentName: departments[i].name,
			})
			j++
		} else if departments[i].id < employees[j].department_id {
			i++
		} else {
			j++
		}
	}

	return result
}

func main() {
	// Create sample data for employees and departments
	employees := Employees{
		Employee{id: 1, name: "Alice", department_id: 3},
		Employee{id: 2, name: "Bob", department_id: 1},
		Employee{id: 3, name: "Charlie", department_id: 1},
		Employee{id: 4, name: "David", department_id: 2},
	}

	departments := Departments{
		Department{id: 1, name: "Engineering"},
		Department{id: 2, name: "Marketing"},
		Department{id: 3, name: "Sales"},
	}

	// Perform the merge join operation
	result := MergeJoin(employees, departments)

	// Print the joined result
	for _, r := range result {
		fmt.Printf("Employee ID: %d, Employee Name: %s, Department ID: %d, Department Name: %s\n", r.EmployeeID, r.EmployeeName, r.DepartmentID, r.DepartmentName)
	}
}
