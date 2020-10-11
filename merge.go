package merge

import (
	"errors"
)

// ┏┳┓┏━╸┏━┓┏━╸┏━╸   ┏━┓┏━┓┏━┓╺┳╸   ╻┏┓╻   ┏━╸┏━┓╻  ┏━┓┏┓╻┏━╸
// ┃┃┃┣╸ ┣┳┛┃╺┓┣╸    ┗━┓┃ ┃┣┳┛ ┃    ┃┃┗┫   ┃╺┓┃ ┃┃  ┣━┫┃┗┫┃╺┓
// ╹ ╹┗━╸╹┗╸┗━┛┗━╸   ┗━┛┗━┛╹┗╸ ╹    ╹╹ ╹   ┗━┛┗━┛┗━╸╹ ╹╹ ╹┗━┛

func MergeSort(list []int) []int {
	left, right := split(list)
	print("%v %v", left, right)
	if len(left) > 1 {
		left = mergeSort(left)
	}

	if len(right) > 1 {
		right = mergeSort(right)
	}

	return merge(left, right)
}

// This function assumes that left and right have already been sorted
func merge(left []int, right []int) []int {
	// If the last element of the left array is smaller than the first
	// element of the right array, assuming both are already sorted,
	// there is nothing left to do.
	if left[len(left)-1] < right[0] {
		return append(left, right...)
	}
	merged := []int{}
	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			removed, _ := remove(0, &right)
			merged = append(merged, removed)
		} else {
			removed, _ := remove(0, &left)
			merged = append(merged, removed)
		}
	}
	merged = append(merged, left...)
	merged = append(merged, right...)
	return merged
}

func split(nums []int) ([]int, []int) {
	mid := len(nums) / 2

	left := nums[:mid]
	right := nums[mid:]
	return left, right
}

func remove(index int, list *[]int) (int, error) {
	templist := *list
	if index > len(templist)-1 {
		return 0, errors.New("Index out of bounds")
	}
	removed := templist[index]
	left := templist[:index]
	right := templist[index+1:]
	*list = append(left, right...)
	return removed, nil
}
