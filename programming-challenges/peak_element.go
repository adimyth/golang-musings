/*
Given an array that is initially increasing and then decreasing (also known as a mountain array), find the peak element. The peak element is the element that is greater than its neighbors.

Ex: 2, 4, 5, 6, 8, 9, 4, 2, 1 - output is 9
*/
package challenges

func FindPeakElement(numbers []int) int {
	/*
		Find peak element in a mountain array using binary search:
		Algorithm:
		a. Initialize left = 0 and right = length of array - 1
		b. While left < right:
			i. Calculate mid = (left + right) / 2
			ii. If arr[mid] < arr[mid + 1], peak is on the right side:
			iii. Set left = mid + 1
			iv. Else, peak is on the left side or at mid:
			v. Set right = mid
		c. Return left as the index of the peak element
	*/

	left, right := 0, len(numbers)-1

	for left < right {
		mid := (left + right) / 2

		if numbers[mid] < numbers[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return numbers[left]
}
