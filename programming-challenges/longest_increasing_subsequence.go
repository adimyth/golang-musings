/*
Problem Statement: Given an array of integers, find the length of the longest increasing subsequence. Ex: [10, 9, 2, 5, 3, 7, 101, 18] - output is 4 (2, 3, 7, 101)

Algorithm:
a. Initialize an array dp[] of the same length as the input array, all elements set to 1
b. For i from 1 to n-1:
	i. For j from 0 to i-1:
		a. If arr[i] > arr[j] and dp[i] < dp[j] + 1:
		b. Update dp[i] = dp[j] + 1
c. Return the maximum value in dp[] as the length of LIS
*/

package challenges

import "slices"

func LongestIncreasingSubsequence(numbers []int) int {
	n := len(numbers)

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if (numbers[i] > numbers[j]) && (dp[i] < dp[j]+1) {
				dp[i] = dp[j] + 1
			}
		}
	}

	return slices.Max(dp)
}
