package funk

// ContainsInt return true if an int is present in a iteratee.
func ContainsInt(s []int, v int) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

// ContainsInt64 return true if an int64 is present in a iteratee.
func ContainsInt64(s []int64, v int64) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

// ContainsString return true if a string is present in a iteratee.
func ContainsString(s []string, v string) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

// ContainsFloat32 return true if a float32 is present in a iteratee.
func ContainsFloat32(s []float32, v float32) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

// ContainsFloat64 return true if a float64 is present in a iteratee.
func ContainsFloat64(s []float64, v float64) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

// SumInt64 sums a int64 iteratee and returns the sum of all elements
func SumInt64(s []int64) (sum int64) {
	for _, v := range s {
		sum += v
	}
	return
}

// SumInt sums a int iteratee and returns the sum of all elements
func SumInt(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}
	return
}

// SumFloat64 sums a float64 iteratee and returns the sum of all elements
func SumFloat64(s []float64) (sum float64) {
	for _, v := range s {
		sum += v
	}
	return
}

// SumFloat32 sums a float32 iteratee and returns the sum of all elements
func SumFloat32(s []float32) (sum float32) {
	for _, v := range s {
		sum += v
	}
	return
}

// ReverseString reverses an array of string
func ReverseString(s []string) []string {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// ReverseInt reverses an array of int
func ReverseInt(s []int) []int {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// ReverseInt64 reverses an array of int64
func ReverseInt64(s []int64) []int64 {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// ReverseFloat64 reverses an array of float64
func ReverseFloat64(s []float64) []float64 {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// ReverseFloat32 reverses an array of float32
func ReverseFloat32(s []float32) []float32 {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
