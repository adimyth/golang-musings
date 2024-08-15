# Nested Loop Join
Nested Loop Join is one of the simplest join algorithms used to combine rows from two tables based on a common key. The algorithm iterates over each row in the first table (outer table) and then searches for matching rows in the second table (inner table). If a match is found, the rows are combined and included in the result set.

## Algorithm

1. Initialize the result set as empty.
2. For each row in the outer table:
   1. For each row in the inner table:
   2. Compare the join key of the outer table row with the join key of the inner table row.
   3. If the keys match, combine the rows and add them to the result set.
3. Return the result set after all rows have been processed.

## Pseudocode

```plaintext
function NestedLoopJoin(outerTable, innerTable):
    result = []
    
    for outerRow in outerTable:
        for innerRow in innerTable:
            if outerRow.joinKey == innerRow.joinKey:
                result.append(combine(outerRow, innerRow))
    
    return result
```

## Time & Space Complexity
### Time Complexity

* Best Case: `O(n * m)` (where n is the number of rows in the outer table, and m is the number of rows in the inner table).
* Worst Case: `O(n * m)` (same as the best case, as each row in the outer table is compared with every row in the inner table).

### Space Complexity

`O(1)` for the algorithm itself, but `O(n * m)` for the result set in the worst case.


## Example
Employees:
| ID  | Name    | Department ID |
| --- | ------- | ------------- |
| 1   | Alice   | 3             |
| 2   | Bob     | 1             |
| 3   | Charlie | 1             |
| 4   | David   | 2             |

Departments:
| ID  | Name         |
| --- | ------------ |
| 1   | Engineering  |
| 2   | Marketing    |
| 3   | Sales        |

Join Process:

* For each employee, iterate through all departments and compare department_id:
  * Bob (department_id=1) matches with Engineering (id=1)
  * Charlie (department_id=1) matches with Engineering (id=1)
  * David (department_id=2) matches with Marketing (id=2)
  * Alice (department_id=3) matches with Sales (id=3)

Result:
| Employee ID | Employee Name | Department ID | Department Name |
| ----------- | ------------- | ------------- | --------------- |
| 2           | Bob           | 1             | Engineering     |
| 3           | Charlie       | 1             | Engineering     |
| 4           | David         | 2             | Marketing       |
| 1           | Alice         | 3             | Sales           |