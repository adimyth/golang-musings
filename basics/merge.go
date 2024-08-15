package basics

/*
Merge Sort is a divide-and-conquer algorithm that divides the input array into two halves, recursively sorts them, and then merges the sorted halves.

1. Divide Step:
	1. Split the array into two halves:
	2. If the array has one or zero elements, it's already sorted, so return.
	3. Otherwise, find the middle index: mid = (start + end) / 2.
	4. Recursively sort the left half: mergeSort(arr, start, mid).
	5. Recursively sort the right half: mergeSort(arr, mid + 1, end).
2. Conquer Step (Merge Process):
	1. Merge the two halves back together:
	2. Create two temporary arrays: left for the first half, and right for the second half.
	3. Initialize three pointers:
		a. i for traversing left array.
		b. j for traversing right array.
		c. k for traversing the original array.
	4. Compare elements from left and right:
		a. If left[i] <= right[j], place left[i] into the original array and increment i.
		b. Otherwise, place right[j] into the original array and increment j.
	Continue this process until all elements from left and right are placed back into the original array.
3. Combine Step:
	1. After merging, the array will be sorted.
	2. This merging process is repeated recursively, and the size of the array to be merged grows larger with each level of recursion.


Time Complexity:
	1. Best Case: O(n log n)
	2. Average Case: O(n log n)
	3. Worst Case: O(n log n)

Space Complexity: O(n)


Pseudocode:
function mergeSort(arr):
    if length of arr <= 1:
        return arr

    mid = length of arr / 2
    leftHalf = arr[0:mid]
    rightHalf = arr[mid:length of arr]

    leftSorted = mergeSort(leftHalf)
    rightSorted = mergeSort(rightHalf)

    return merge(leftSorted, rightSorted)

function merge(left, right):
    result = []
    i, j = 0, 0

    while i < length of left and j < length of right:
        if left[i] <= right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1

    // Append remaining elements from left or right
    result.extend(left[i:])
    result.extend(right[j:])

    return result
*/

func merge(numbers []int, left int, mid int, right int) {

	// create a temporary array to store the merged array
	leftArray := make([]int, mid-left+1)
	rightArray := make([]int, right-mid)

	// Copy data from mid to left into leftArray and mid+1 to right into rightArray
	for i := 0; i < len(leftArray); i++ {
		leftArray[i] = numbers[i+left]
	}
	for j := 0; j < len(rightArray); j++ {
		rightArray[j] = numbers[j+mid+1]
	}

	// Set indexes for leftArray, rightArray, and numbers
	i := 0
	j := 0
	k := left

	// Merge the leftArray and rightArray into numbers
	for i < len(leftArray) && j < len(rightArray) {
		if leftArray[i] <= rightArray[j] {
			numbers[k] = leftArray[i]
			i++
		} else {
			numbers[k] = rightArray[j]
			j++
		}
		k++
	}

	// Copy the remaining elements of leftArray into numbers
	for i < len(leftArray) {
		numbers[k] = leftArray[i]
		k++
		i++
	}

	// Copy the remaining elements of rightArray into numbers
	for j < len(rightArray) {
		numbers[k] = rightArray[j]
		k++
		j++
	}
}

func MergeSort(numbers []int, left int, right int) {
	if left < right {
		mid := (left + right) / 2

		MergeSort(numbers, left, mid)
		MergeSort(numbers, mid+1, right)
		merge(numbers, left, mid, right)
	}
}
