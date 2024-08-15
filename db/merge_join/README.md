# Merge Join

Merge Join is a highly efficient join algorithm that works by first sorting the two tables on the join key. Once sorted, the algorithm iterates through both tables, comparing the join keys. When a match is found, the rows are combined, and the process continues until all matching rows have been processed.

## Algorithm
1. Sort the Input Tables:
   1. Ensure that both the left table (employees) and the right table (departments) are sorted based on the join key (department_id and id, respectively).
2. Initialize Pointers:
   1. Use two pointers i and j to traverse the left table and right table, respectively.
3. Comparison Loop:
   1. Compare the join key of the current row in the left table (pointed by i) with the join key of the current row in the right table (pointed by j).
   2. If they match:
      1. Join the rows.
      2. Store the result.
      3. Move the right pointer (j) forward to check for additional matches.
   3. If the left table's join key is smaller:
      1. Move the left pointer (i) forward.
   4. If the right table's join key is smaller:
   5. Move the right pointer (j) forward.
4. End Condition: Repeat the comparison until one of the pointers reaches the end of its table.


## Pseudocode

```plaintext
function MergeJoin(employees, departments):
    sort(employees by department_id)
    sort(departments by id)
    
    i, j = 0, 0
    result = []

    while i < length(employees) and j < length(departments):
        if employees[i].department_id == departments[j].id:
            result.append(merge(employees[i], departments[j]))
            j += 1
        else if employees[i].department_id < departments[j].id:
            i += 1
        else:
            j += 1

    return result
```

## Time & Space Complexity
### Time Complexity

* Best Case: `O(n log n + m log m)` (where n is the number of rows in the left table, and m is the number of rows in the right table). This is the time complexity required for sorting the input tables.

### Space Complexity
The space complexity primarily depends on the space required for storing the sorted tables and the result set.

* Sorting Space: `O(n + m)` (for the sorted tables).

## Example
Employees:
| ID  | Name    | Department ID |
| --- | ------- | ------------- |
| 1   | Alice   | 3             |
| 2   | Bob     | 1             |
| 3   | Charlie | 1             |
| 4   | David   | 2             |


Departments:
| ID  | Name        |
| --- | ----------- |
| 1   | Engineering |
| 2   | Marketing   |
| 3   | Sales       |

Step-by-Step Execution:

**Sorting:**
* Employees are sorted by department_id: [Bob, Charlie, David, Alice]
* Departments are sorted by id: [Engineering, Marketing, Sales]

**Merging:**
* Start with pointers i=0 and j=0.
* Comparison: Bob’s department_id (1) matches Engineering’s id (1). Join and move j forward.
* Comparison: Charlie’s department_id (1) also matches Engineering’s id (1). Join and move j forward.
* Comparison: David’s department_id (2) matches Marketing’s id (2). Join and move j forward.
* Comparison: Alice’s department_id (3) matches Sales’s id (3). Join and move j forward.
* End when j reaches the end of departments.

**Result:**
| Employee ID | Employee Name | Department ID | Department Name |
| ----------- | ------------- | ------------- | --------------- |
| 2           | Bob           | 1             | Engineering     |
| 3           | Charlie       | 1             | Engineering     |
| 4           | David         | 2             | Marketing       |
| 1           | Alice         | 3             | Sales           |