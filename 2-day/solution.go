package main

// Find if array is decreasing or increasing no more than by 1 to 3
func checkArray(arr []int) bool {
	isDecreasing := true
	isIncreasing := true
	for i := 0; i < len(arr)-1; i++ {
		difference := Abs(arr[i] - arr[i+1])
		if difference > 3 {
			return false
		}
		if arr[i] >= arr[i+1] {
			isIncreasing = false
		}
		if arr[i] <= arr[i+1] {
			isDecreasing = false
		}
	}
	return (isIncreasing || isDecreasing) && !(isIncreasing && isDecreasing)
}

// Same check as above but removing a level from an unsafe report makes it safe
func checkArray2(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		// create new array without the i-th element
		arr2 := make([]int, len(arr)-1)
		copy(arr2, arr[:i])
		copy(arr2[i:], arr[i+1:])
		isSafe := checkArray(arr2)
		if isSafe {
			return true
		}
	}
	return false
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
