/*
Hash Join
=========

A Hash Join is an efficient algorithm used to perform join operations in relational databases. It is particularly useful when joining large datasets and is often faster than nested loop joins or merge joins, especially when the datasets are not sorted.
The basic idea behind a hash join is to use a hash table to partition the smaller of the two tables (referred to as the "build" table), and then probe the hash table with the larger table (referred to as the "probe" table) to find matching rows.

Algorithm:
1. Build a hash table from the smaller of the two tables based on the join key.
2. Iterate over the larger table and probe the hash table to find matching rows.
3. For each matching row, combine the rows from both tables to form the result.

Time Complexity:
1. Building the hash table takes O(n) time, where n is the number of rows in the smaller table.
2. Probing the hash table takes O(m) time, where m is the number of rows in the larger table.
3. Overall, the time complexity is O(n + m).

Space Complexity: O(n)

Additional Notes:
1. This is well suited for "equi-joins" where the join condition is based on equality. Ex - `employees.department_id = departments.id`.
2. Hash table requires additional memory, so it may not be suitable for very large datasets that do not fit in memory.

Example:
In this example, we will create two tables, `employees` and `departments`, and join them on the `department_id` column using the Hash Join algorithm.
*/
package main

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

type Employees []Employee
type Departments []Department

func HashJoin(employees Employees, departments Departments) []Result {
	var result []Result

	// NOTE: Here, I am creating a HashMap with the department_id as the key and the department name as the value. In reality, it's more complex, and you might need to handle hash collisions, etc.
	// Assuming that the "departments" table is smaller than the "employees" table.
	// Build a hash table from the smaller table (departments)
	departmentHash := make(map[int]string)
	for _, dept := range departments {
		departmentHash[dept.id] = dept.name
	}

	// Probe the hash table with the larger table (employees)
	for _, emp := range employees {
		if deptName, ok := departmentHash[emp.department_id]; ok {
			result = append(result, Result{EmployeeID: emp.id, EmployeeName: emp.name, DepartmentID: emp.department_id, DepartmentName: deptName})
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

	// Perform Hash Join on the employees and departments tables
	result := HashJoin(employees, departments)

	// Print the joined result
	for _, res := range result {
		println(res.EmployeeID, res.EmployeeName, res.DepartmentID, res.DepartmentName)
	}
}
