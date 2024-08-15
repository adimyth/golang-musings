/*
 Search Missing number in an array containing sequential numeric values. Ex: 1,2,3,4,5,7,8,9 - output is 6
*/

package challenges

func FindMissingNumber(numbers []int) int {
	/*
		1. Initialize left=0 and right=(length of array) - 1
		2. While left <= right
			a. mid = (left + right) / 2
			b. if arr[mid] - mid == 1, missing number is on the right side, set "left = mid + 1"
			c. else, missing number is on the left side, right = mid - 1
		3. Return left + 1 as missing number
	*/
	left, right := 0, len(numbers)-1

	for left <= right {
		mid := (left + right) / 2

		if numbers[mid]-mid == 1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left + 1
}
