package lqarray

// Contains checks if an element exists in an array/slice.
//
// Parameters:
//   - array: The array/slice to search in
//   - elem: The element to search for
//
// Returns:
//   - bool: true if element exists, false otherwise
//
// Example:
//
//	Contains([]string{"a", "b", "c"}, "b") // returns true
//	Contains([]int{1, 2, 3}, 4) // returns false
func Contains[T comparable](array []T, elem T) bool {
	for _, v := range array {
		if v == elem {
			return true
		}
	}
	return false
}

// ContainsFunc checks if an element exists in an array/slice using a custom comparison function.
//
// Parameters:
//   - array: The array/slice to search in
//   - elem: The element to search for
//   - equals: Custom comparison function
//
// Returns:
//   - bool: true if element exists, false otherwise
//
// Example:
//
//	type Person struct { Name string }
//	people := []Person{{Name: "John"}, {Name: "Jane"}}
//	ContainsFunc(people, Person{Name: "John"}, func(a, b Person) bool {
//	    return a.Name == b.Name
//	}) // returns true
func ContainsFunc[T any](array []T, elem T, equals func(T, T) bool) bool {
	for _, v := range array {
		if equals(v, elem) {
			return true
		}
	}
	return false
}

// ContainsBy checks if any element in the array satisfies the predicate.
//
// Parameters:
//   - array: The array/slice to search in
//   - predicate: Function that returns true when element matches criteria
//
// Returns:
//   - bool: true if any element satisfies the predicate, false otherwise
//
// Example:
//
//	type User struct { Age int }
//	users := []User{{Age: 25}, {Age: 30}}
//	ContainsBy(users, func(u User) bool {
//	    return u.Age > 28
//	}) // returns true
func ContainsBy[T any](array []T, predicate func(T) bool) bool {
	for _, v := range array {
		if predicate(v) {
			return true
		}
	}
	return false
}

// Unique returns a new array/slice with duplicate elements removed.
//
// Parameters:
//   - array: The array/slice to remove duplicates from
//
// Returns:
//   - []T: New array/slice with unique elements
//
// Example:
//
//	Unique([]int{1, 2, 2, 3, 3, 4}) // returns [1, 2, 3, 4]
//	Unique([]string{"a", "b", "b", "c"}) // returns ["a", "b", "c"]
func Unique[T comparable](array []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0)

	for _, v := range array {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Chunk splits an array/slice into chunks of specified size.
//
// Parameters:
//   - array: The array/slice to split
//   - size: Size of each chunk
//
// Returns:
//   - [][]T: Array/slice of chunks
//
// Example:
//
//	Chunk([]int{1, 2, 3, 4, 5}, 2) // returns [[1, 2], [3, 4], [5]]
//	Chunk([]string{"a", "b", "c"}, 2) // returns [["a", "b"], ["c"]]
func Chunk[T any](array []T, size int) [][]T {
	if size <= 0 {
		return nil
	}

	chunks := make([][]T, 0, (len(array)+size-1)/size)
	for i := 0; i < len(array); i += size {
		end := i + size
		if end > len(array) {
			end = len(array)
		}
		chunks = append(chunks, array[i:end])
	}
	return chunks
}

// Map applies a function to each element in an array/slice.
//
// Parameters:
//   - array: The input array/slice
//   - fn: Function to apply to each element
//
// Returns:
//   - []U: New array/slice with transformed elements
//
// Example:
//
//	Map([]int{1, 2, 3}, func(x int) int { return x * 2 }) // returns [2, 4, 6]
//	Map([]string{"a", "b"}, func(s string) string { return s + "!" }) // returns ["a!", "b!"]
func Map[T any, U any](array []T, fn func(T) U) []U {
	result := make([]U, len(array))
	for i, v := range array {
		result[i] = fn(v)
	}
	return result
}

// Filter returns a new array/slice with elements that satisfy the predicate.
//
// Parameters:
//   - array: The input array/slice
//   - fn: Predicate function
//
// Returns:
//   - []T: Filtered array/slice
//
// Example:
//
//	Filter([]int{1, 2, 3, 4}, func(x int) bool { return x > 2 }) // returns [3, 4]
//	Filter([]string{"a", "b", "c"}, func(s string) bool { return s != "b" }) // returns ["a", "c"]
func Filter[T any](array []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range array {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce reduces an array/slice to a single value using an accumulator function.
//
// Parameters:
//   - array: The input array/slice
//   - fn: Accumulator function
//   - initial: Initial value
//
// Returns:
//   - T: Accumulated result
//
// Example:
//
//	Reduce([]int{1, 2, 3}, func(acc, x int) int { return acc + x }, 0) // returns 6
//	Reduce([]string{"a", "b"}, func(acc, s string) string { return acc + s }, "") // returns "ab"
func Reduce[T any, U any](array []T, fn func(U, T) U, initial U) U {
	result := initial
	for _, v := range array {
		result = fn(result, v)
	}
	return result
}

// Reverse returns a new array/slice with elements in reverse order.
//
// Parameters:
//   - array: The input array/slice
//
// Returns:
//   - []T: Reversed array/slice
//
// Example:
//
//	Reverse([]int{1, 2, 3}) // returns [3, 2, 1]
//	Reverse([]string{"a", "b", "c"}) // returns ["c", "b", "a"]
func Reverse[T any](array []T) []T {
	result := make([]T, len(array))
	for i, j := 0, len(array)-1; i < len(array); i, j = i+1, j-1 {
		result[i] = array[j]
	}
	return result
}

// IndexOf returns the index of the first occurrence of an element.
//
// Parameters:
//   - array: The array/slice to search in
//   - elem: Element to find
//
// Returns:
//   - int: Index of element (-1 if not found)
//
// Example:
//
//	IndexOf([]int{1, 2, 3}, 2) // returns 1
//	IndexOf([]string{"a", "b", "c"}, "d") // returns -1
func IndexOf[T comparable](array []T, elem T) int {
	for i, v := range array {
		if v == elem {
			return i
		}
	}
	return -1
}
