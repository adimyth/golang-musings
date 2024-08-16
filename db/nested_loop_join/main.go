package main

import "fmt"

/*
Nested Loop Join

Nested Loop Join is a join algorithm that works by iterating over each row in one table and comparing it with each row in the other table. If the join condition is satisfied, the rows are joined and returned as the result.

1. For each row in the left table:
	a. For each row in the right table:
		i. Check if the join condition is satisfied.
		ii. If the join condition is satisfied, join the rows and store the result in the resultset.

Time Complexity:
1. The time complexity of the Nested Loop Join algorithm is O(n * m), where n is the number of rows in the left table and m is the number of rows in the right table.

Space Complexity: O(1)

Example:
In this example, we will create two tables, `employees` and `departments`, and join them on the `department_id` column using the Nested Loop Join algorithm.

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
	employee_id     int
	employee_name   string
	department_id   int
	department_name string
}

type Employees []Employee
type Departments []Department

func NestedLoopJoin(employees Employees, departments Departments) []Result {
	var result []Result

	for i := 0; i <= len(employees)-1; i++ {
		for j := 0; j <= len(departments)-1; j++ {
			if employees[i].department_id == departments[j].id {
				result = append(result, Result{employee_id: employees[i].id, employee_name: employees[i].name, department_id: employees[i].department_id, department_name: departments[j].name})
			}
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

	// Perform the Nested Loop Join operation
	result := NestedLoopJoin(employees, departments)

	// Print the joined result
	for _, r := range result {
		fmt.Printf("Employee ID: %d, Employee Name: %s, Department ID: %d, Department Name: %s\n", r.employee_id, r.employee_name, r.department_id, r.department_name)
	}
}
