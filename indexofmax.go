package funk

import "strings"

// IndexOfMaxInt returns the index of the maximum element in an array/slice.
// If there are duplicate occurrences of max element, only return the first one.
// This function implements the argMax functionality requested in GitHub issue #139.
// It accepts []int.
// It returns int.
func IndexOfMaxInt(arr []int) int {
	if len(arr) == 0 {
		panic("arg is an empty array/slice")
	}

	var maxValue int
	var maxIndex int

	for idx := 0; idx < len(arr); idx++ {
		item := arr[idx]
		if idx == 0 {
			maxValue = item
			maxIndex = 0

			continue
		}

		if item > maxValue {
			maxValue = item
			maxIndex = idx
		}
	}

	return maxIndex
}

// IndexOfMaxInt8 returns the index of the maximum element in an array/slice.
// If there are duplicate occurrences of max element, only return the first one.
// It accepts []int8.
// It returns int.
func IndexOfMaxInt8(arr []int8) int {
	if len(arr) == 0 {
		panic("arg is an empty array/slice")
	}

	var maxValue int8
	var maxIndex int

	for idx := 0; idx < len(arr); idx++ {
		item := arr[idx]
		if idx == 0 {
			maxValue = item
			maxIndex = 0

			continue
		}

		if item > maxValue {
			maxValue = item
			maxIndex = idx
		}
	}

	return maxIndex
}

// IndexOfMaxInt16 returns the index of the maximum element in an array/slice.
// If there are duplicate occurrences of max element, only return the first one.
// It accepts []int16.
// It returns int.
func IndexOfMaxInt16(arr []int16) int {
	if len(arr) == 0 {
		panic("arg is an empty array/slice")
	}

	var maxValue int16
	var maxIndex int

	for idx := 0; idx < len(arr); idx++ {
		item := arr[idx]
		if idx == 0 {
			maxValue = item
			maxIndex = 0

			continue
		}

		if item > maxValue {
			maxValue = item
			maxIndex = idx
		}
	}

	return maxIndex
}

// IndexOfMaxInt32 returns the index of the maximum element in an array/slice.
// If there are duplicate occurrences of max element, only return the first one.
// It accepts []int32.
// It returns int.
func IndexOfMaxInt32(arr []int32) int {
	if len(arr) == 0 {
		panic("arg is an empty array/slice")
	}

	var maxValue int32

	var maxIndex int

	for idx := 0; idx < len(arr); idx++ {
		item := arr[idx]
		if idx == 0 {
			maxValue = item
			maxIndex = 0

			continue
		}

		if item > maxValue {
			maxValue = item
			maxIndex = idx
		}
	}

	return maxIndex
}

// IndexOfMaxInt64 returns the index of the maximum element in an array/slice.
// If there are duplicate occurrences of max element, only return the first one.
// It accepts []int64.
// It returns int.
func IndexOfMaxInt64(arr []int64) int {
	if len(arr) == 0 {
		panic("arg is an empty array/slice")
	}

	var maxValue int64

	var maxIndex int

	for idx := 0; idx < len(arr); idx++ {
		item := arr[idx]
		if idx == 0 {
			maxValue = item
			maxIndex = 0

			continue
		}

		if item > maxValue {
			maxValue = item
			maxIndex = idx
		}
	}

	return maxIndex
}

// IndexOfMaxFloat32 returns the index of the maximum element in an array/slice.
// If there are duplicate occurrences of max element, only return the first one.
// It accepts []float32.
// It returns int.
func IndexOfMaxFloat32(arr []float32) int {
	if len(arr) == 0 {
		panic("arg is an empty array/slice")
	}

	var maxValue float32

	var maxIndex int

	for idx := 0; idx < len(arr); idx++ {
		item := arr[idx]
		if idx == 0 {
			maxValue = item
			maxIndex = 0

			continue
		}

		if item > maxValue {
			maxValue = item
			maxIndex = idx
		}
	}

	return maxIndex
}

// IndexOfMaxFloat64 returns the index of the maximum element in an array/slice.
// If there are duplicate occurrences of max element, only return the first one.
// It accepts []float64.
// It returns int.
func IndexOfMaxFloat64(arr []float64) int {
	if len(arr) == 0 {
		panic("arg is an empty array/slice")
	}

	var maxValue float64
	var maxIndex int

	for idx := 0; idx < len(arr); idx++ {
		item := arr[idx]
		if idx == 0 {
			maxValue = item
			maxIndex = 0

			continue
		}

		if item > maxValue {
			maxValue = item
			maxIndex = idx
		}
	}

	return maxIndex
}

// IndexOfMaxString returns the index of the maximum element in an array/slice.
// If there are duplicate occurrences of max element, only return the first one.
// It accepts []string.
// It returns int.
func IndexOfMaxString(arr []string) int {
	if len(arr) == 0 {
		panic("arg is an empty array/slice")
	}

	var maxValue string
	var maxIndex int

	for idx := 0; idx < len(arr); idx++ {
		item := arr[idx]
		if idx == 0 {
			maxValue = item
			maxIndex = 0

			continue
		}

		if compareStringsIndexOfMax(maxValue, item) == item {
			maxValue = item
			maxIndex = idx
		}
	}

	return maxIndex
}

// compareStringsIndexOfMax uses the strings.Compare method to compare two strings, and returns the greater one.
func compareStringsIndexOfMax(maxValue, current string) string {
	r := strings.Compare(strings.ToLower(maxValue), strings.ToLower(current))
	if r >= 0 {
		return maxValue
	}

	return current
}
