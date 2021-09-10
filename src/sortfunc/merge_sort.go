package sortfunc

// MergeSort performs a merge sort on a slice of integers.
// This algorithm is extremely efficient as n becomes large.
func MergeSort(slice []int) ([]int, error) {
	length := len(slice)
	// If our slice isn't long enough to require splitting,
	// just copy, sort, and return .
	if length == 2 && slice[0] > slice[1] {
		ret := []int{slice[1], slice[0]}
		return ret, nil
	} else if length <= 2 {
		ret := make([]int, length)
		copy(ret, slice)
		return ret, nil
	// If our algorithm needs to split the array, recursively split it until every section is of length 1 or 2.
	// Splitting is done in concurrent goroutines, and responses are sent back up the stack through channels,
	// and then merged in their parent goroutines .
	} else {
		halfLength := length / 2
		left := slice[:halfLength]
		right := slice[halfLength:]
		leftChan := make(chan []int, 2)
		rightChan := make(chan []int, 2)

		func(splitSlice chan []int) {
			val, _ := MergeSort(left)
			splitSlice <- val
		}(leftChan)

		func(splitSlice chan []int) {
			val, _ := MergeSort(right)
			splitSlice <- val
		}(rightChan)

		left = <-leftChan
		right = <-rightChan
		//Return the arrays merged together in order.
		return merge(left, right), nil
	}
}

func merge(left []int, right []int) []int {
	merged := make([]int, len(left)+len(right))
	leftOverflow := false
	rightOverflow := false
	for i, j, k := 0, 0, 0; k < len(merged); k++ {
		leftOverflow = i >= len(left)
		rightOverflow = j >= len(right)
		if !leftOverflow && !rightOverflow {
			if left[i] < right[j] {
				merged[k] = left[i]
				i++
			} else {
				merged[k] = right[j]
				j++
			}
		} else if leftOverflow {
			merged[k] = right[j]
			j++
		} else {
			merged[k] = left[i]
			i++
		}
	}
	return merged
}
