package main

import (
	"fmt"

	"github.com/scbrickley/merge"
)

func main() {
	nums := []int{
		50,
		45,
		42,
		43,
		27,
		34,
		16,
		13,
		9,
		2,
		5,
	}

	fmt.Printf("Start: %v\n", nums)
	sorted := merge.MergeSort(nums)
	fmt.Printf("Finish: %v\n", sorted)
}
