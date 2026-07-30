package main

import (
	"fmt"
)

func main() {
	// Example sorted array for testing
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	// Test binary search
	target := 7
	index := BinarySearch(arr, target)
	fmt.Printf("BinarySearch: %d found at index %d\n", target, index)

	// Test recursive binary search
	recursiveIndex := BinarySearchRecursive(arr, target, 0, len(arr)-1)
	fmt.Printf("BinarySearchRecursive: %d found at index %d\n", target, recursiveIndex)

	// Test find insert position
	insertTarget := 8
	insertPos := FindInsertPosition(arr, insertTarget)
	fmt.Printf("FindInsertPosition: %d should be inserted at index %d\n", insertTarget, insertPos)
}

// BinarySearch performs a standard binary search to find the target in the sorted array.
// Returns the index of the target if found, or -1 if not found.
func BinarySearch(arr []int, target int) int {
	// TODO: Implement this function
	for i:=0 ; i<len(arr);i++{
	    if arr[i] == target{
	        return i
	    }
	}
	return -1
}

// BinarySearchRecursive performs binary search using recursion.
// Returns the index of the target if found, or -1 if not found.
func BinarySearchRecursive(arr []int, target int, left int, right int) int {
	// TODO: Implement this function
	left = 0
	right = len(arr)-1
	if len(arr) ==1 && target ==  arr[0]{
	    return 0
	}
	mid := left + (right-left)/2
	for left<right{
	    if arr[mid] == target {
		return mid
	}
	    if target == arr[left]{
	        return left
	    }else if target == arr[right] {
	        return right
	    }else{
	        left ++
	        right--
	    }
	}
	return -1
}

// FindInsertPosition returns the index where the target should be inserted
// to maintain the sorted order of the array.
func FindInsertPosition(arr []int, target int) int {
	// TODO: Implement this function
	if len(arr) == 0 {
	    return 0
	}
	for i:=0 ; i<len(arr);i++{
	    
	    if target == arr[i]{
	        return i
	    }else if target <= arr[0]{
	        return 0
	    }else if target > arr[len(arr)-1]{
	        return len(arr)
	    }else{
	        if arr[i]<target && arr[i+1]>=target{
	            return i+1
	    }
	    
	  
	  }
	}
	
	return -1
}


