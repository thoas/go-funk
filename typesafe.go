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

func indexOf(n int, f func(int) bool) int {
	for i := 0; i < n; i++ {
		if f(i) {
			return i
		}
	}
	return -1
}

// IndexOfInt gets the index at which the first occurrence of an int value is found in array or return -1
// if the value cannot be found
func IndexOfInt(a []int, x int) int {
	return indexOf(len(a), func(i int) bool { return a[i] == x })
}

// IndexOfInt64 gets the index at which the first occurrence of an int64 value is found in array or return -1
// if the value cannot be found
func IndexOfInt64(a []int64, x int64) int {
	return indexOf(len(a), func(i int) bool { return a[i] == x })
}

// IndexOfFloat64 gets the index at which the first occurrence of an float64 value is found in array or return -1
// if the value cannot be found
func IndexOfFloat64(a []float64, x float64) int {
	return indexOf(len(a), func(i int) bool { return a[i] == x })
}

// IndexOfString gets the index at which the first occurrence of a string value is found in array or return -1
// if the value cannot be found
func IndexOfString(a []string, x string) int {
	return indexOf(len(a), func(i int) bool { return a[i] == x })
}

// UniqInt64 creates an array of int64 with unique values.
func UniqInt64(a []int64) []int64 {
	length := len(a)

	seen := make(map[int64]bool, length)
	j := 0

	for i := 0; i < length; i++ {
		v := a[i]

		if _, ok := seen[v]; ok {
			continue
		}

		seen[v] = true
		a[j] = v
		j++
	}

	return a[0:j]
}

// UniqInt creates an array of int with unique values.
func UniqInt(a []int) []int {
	length := len(a)

	seen := make(map[int]bool, length)
	j := 0

	for i := 0; i < length; i++ {
		v := a[i]

		if _, ok := seen[v]; ok {
			continue
		}

		seen[v] = true
		a[j] = v
		j++
	}

	return a[0:j]
}

// UniqString creates an array of string with unique values.
func UniqString(a []string) []string {
	length := len(a)

	seen := make(map[string]bool, length)
	j := 0

	for i := 0; i < length; i++ {
		v := a[i]

		if _, ok := seen[v]; ok {
			continue
		}

		seen[v] = true
		a[j] = v
		j++
	}

	return a[0:j]
}

// UniqFloat64 creates an array of float64 with unique values.
func UniqFloat64(a []float64) []float64 {
	length := len(a)

	seen := make(map[float64]bool, length)
	j := 0

	for i := 0; i < length; i++ {
		v := a[i]

		if _, ok := seen[v]; ok {
			continue
		}

		seen[v] = true
		a[j] = v
		j++
	}

	return a[0:j]
}
