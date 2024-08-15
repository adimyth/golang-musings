package main

import (
	"fmt"
	"musings/basics"
	datastructures "musings/data-structures"
	"musings/maths"
	challenges "musings/programming-challenges"
)

func main() {
	fmt.Print("=====================================\n")
	fmt.Println("Golang Basics")
	fmt.Print("=====================================\n")
	fmt.Println("Define a Shape interface with Area & Perimeter methods. Implement the interface for Circle & Rectangle shapes. The Circle & Rectangle should have Area & Perimeter methods.")
	r := basics.Rectangle{Length: 10, Breadth: 5}
	fmt.Printf("[Rectangle] Area: %.2f\n", r.Area())
	fmt.Printf("[Rectangle] Perimeter: %.2f\n", r.Perimeter())

	c := basics.Circle{Radius: 5}
	fmt.Printf("[Circle] Area: %.2f\n", c.Area())
	fmt.Printf("[Circle] Perimeter: %.2f\n", c.Perimeter())
	fmt.Print("\n\n")

	numbers := []int{10, 5, 3, 7, 8, 2, 1, 6}
	fmt.Println("Merge Sort: Sort the given array of numbers using Merge Sort algorithm.")
	fmt.Printf("Input: %v\n", numbers)
	basics.MergeSort(numbers, 0, len(numbers)-1)
	fmt.Printf("Output: %v\n", numbers)
	fmt.Print("\n\n")

	fmt.Print("=====================================\n")
	fmt.Println("Maths")
	fmt.Print("=====================================\n")
	n := 10
	fmt.Printf("Iterative Fibonacci (%d): %d\n", n, maths.Iterative(n))
	fmt.Printf("Recursive Fibonacci (%d): %d\n", n, maths.Recursive(n))
	fmt.Print("\n\n")

	fmt.Printf("Sive of Erasothesnes Primes (%d): %v\n", n, maths.Sieve(n))
	fmt.Print("\n\n")

	fmt.Print("=====================================\n")
	fmt.Println("Programming Challenges")
	fmt.Print("=====================================\n")
	fmt.Println("Search Missing number in an array containing sequential numeric values. Ex: 1,2,3,4,5,7,8,9 - output is 6")
	numbers = []int{1, 2, 3, 4, 5, 6, 8, 9}
	fmt.Printf("Numbers: %v, Missing Number: %d\n", numbers, challenges.FindMissingNumber(numbers))
	fmt.Print("\n\n")

	fmt.Println("Given an array that is initially increasing and then decreasing (also known as a mountain array), find the peak element. The peak element is the element that is greater than its neighbors. Ex: 2, 4, 5, 6, 8, 9, 4, 2, 1 - output is 9")
	numbers = []int{2, 4, 5, 6, 8, 9, 10, 4, 2, 1}
	fmt.Printf("Numbers: %v, Peak Element: %d\n", numbers, challenges.FindPeakElement(numbers))
	fmt.Print("\n\n")

	fmt.Println("Given an array arr[] of size N, the task is to find the length of the Longest Increasing Subsequence (LIS) i.e., the longest possible subsequence in which the elements of the subsequence are sorted in increasing order. Ex: Input: [3, 10, 2, 1, 20], Output: 3")
	numbers = []int{3, 10, 2, 1, 20}
	fmt.Printf("Input: %v, Output: %d\n", numbers, challenges.LongestIncreasingSubsequence(numbers))
	numbers = []int{50, 3, 10, 7, 40, 80}
	fmt.Printf("Input: %v, Output: %d", numbers, challenges.LongestIncreasingSubsequence(numbers))
	fmt.Print("\n\n")

	fmt.Print("=====================================\n")
	fmt.Println("Data Structures")
	fmt.Print("=====================================\n")
	fmt.Println("Implement a Trie (Prefix Tree) with insert & search methods. Ex: Insert: 'apple', Search: 'apple' - output is true, Search: 'app' - output is false")
	trie := datastructures.Constructor()
	fruits := []string{"apple", "banana", "mango", "orange", "grapes"}
	for _, fruit := range fruits {
		trie.Insert(fruit)
	}
	fmt.Printf("Fruits inserted into the Trie: %v\n", fruits)
	fmt.Printf("Search: 'apple' - %t\n", trie.Search("apple"))
	fmt.Printf("Search: 'app' - %t\n", trie.Search("app"))
	fmt.Printf("Search: 'banana' - %t\n", trie.Search("banana"))
	fmt.Printf("Search: 'mango' - %t\n", trie.Search("ban"))
	fmt.Print("\n\n")

	fmt.Println("Implement a Stack data structure with push, pop, peek & isEmpty methods. It should be generic and support any data type. Ex: Push: 1,2,3, Pop: 3,2, Peek: 1, IsEmpty: false")
	num_stack := datastructures.Stack[int]{}
	num_stack.Push(1)
	num_stack.Push(2)
	num_stack.Push(3)
	fmt.Printf("Stack: %v\n", num_stack)
	item, _ := num_stack.Pop()
	fmt.Printf("Pop: %d, Size: %d\n", item, num_stack.Size())
	item, _ = num_stack.Pop()
	fmt.Printf("Pop: %d, Size: %d\n", item, num_stack.Size())
	fmt.Printf("IsEmpty: %t\n", num_stack.IsEmpty())
	num_stack.Clear()
	fmt.Printf("Clear -> IsEmpty: %t\n", num_stack.IsEmpty())
	fmt.Print("\n\n")

	fmt.Println("Implement a Binary Search Tree with insert, search, delete & traversal methods.",
		"Ex: Insert: 10, 5, 15, 3, 7, 12, 18, Search: 7 - output is true, Search: 8 - output is false",
		"InOrderTraversal: 3, 5, 7, 10, 12, 15, 18",
		"PreOrderTraversal: 10, 5, 3, 7, 15, 12, 18",
		"PostOrderTraversal: 3, 7, 5, 12, 18, 15, 10",
		"Delete: 5, InOrderTraversal: 3, 7, 10, 12, 15, 18")
	bst := datastructures.BinarySearchTree{}
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(12)
	bst.Insert(18)
	fmt.Printf("Search: 7 - %t\n", bst.Search(7))
	fmt.Printf("Search: 8 - %t\n", bst.Search(8))
	fmt.Printf("InOrderTraversal: ")
	bst.Traversal()
	fmt.Printf("PreOrderTraversal: ")
	bst.PreOrderTraversal()
	fmt.Printf("PostOrderTraversal: ")
	bst.PostOrderTraversal()
	bst.Delete(5)
	fmt.Printf("Delete: 5, InOrderTraversal: ")
	bst.Traversal()
	fmt.Print("\n\n")
}
