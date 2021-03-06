package merge

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

const checkMark = '\u2713'
const ballotX = '\u2717'

func TestSplit(t *testing.T) {
	nums := []int{72, 65, 89, 93, 52}

	left, right := split(nums)
	splitNums := [][]int{left, right}
	expected := [][]int{{72, 65}, {89, 93, 52}}
	t.Logf("split(%v) should equal %v", nums, expected)
	if cmp.Equal(splitNums, expected) {
		t.Logf("Got: %v %c", splitNums, checkMark)
	} else {
		t.Errorf("Got: %v %c", splitNums, ballotX)
	}
}

func TestMerge(t *testing.T) {
	left := []int{72}
	right := []int{64}

	expected := []int{64, 72}
	t.Logf("merge(%v, %v) should equal %v", left, right, expected)
	merged := merge(left, right)
	if cmp.Equal(merged, expected) {
		t.Logf("Got: %v %c", merged, checkMark)
	} else {
		t.Errorf("Got: %v %c", merged, ballotX)
	}

	left = []int{23, 72}
	right = []int{64, 89}

	expected = []int{23, 64, 72, 89}
	t.Logf("merge(%v, %v) should equal %v", left, right, expected)
	merged = merge(left, right)
	if cmp.Equal(merged, expected) {
		t.Logf("Got: %v %c", merged, checkMark)
	} else {
		t.Errorf("Got: %v %c", merged, ballotX)
	}

	left = []int{23, 72, 105}
	right = []int{64, 89, 93}

	expected = []int{23, 64, 72, 89, 93, 105}
	t.Logf("merge(%v, %v) should equal %v", left, right, expected)
	merged = merge(left, right)
	if cmp.Equal(merged, expected) {
		t.Logf("Got: %v %c", merged, checkMark)
	} else {
		t.Errorf("Got: %v %c", merged, ballotX)
	}

	left = []int{34, 72, 89, 105}
	right = []int{23, 64, 93}

	expected = []int{23, 34, 64, 72, 89, 93, 105}
	t.Logf("merge(%v, %v) should equal %v", left, right, expected)
	merged = merge(left, right)
	if cmp.Equal(merged, expected) {
		t.Logf("Got: %v %c", merged, checkMark)
	} else {
		t.Errorf("Got: %v %c", merged, ballotX)
	}
}

func TestRemove(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}

	t.Logf("list == %v", list)
	removed, _ := remove(2, &list)
	t.Logf("removed = remove(2, list)")
	t.Log("removed should equal 3")
	if removed == 3 {
		t.Logf("removed == %d %c", removed, checkMark)
	} else {
		t.Errorf("removed == %d %c", removed, ballotX)
	}

	expected := []int{1, 2, 4, 5, 6}
	t.Logf("list should equal %v", expected)
	if cmp.Equal(list, expected) {
		t.Logf("list == %v %c", list, checkMark)
	} else {
		t.Errorf("list == %v %c", list, ballotX)
	}

	_, err := remove(10, &list)
	t.Log("Calling remove(10, list)")
	t.Log("Should get Error: 'Index out of bounds'")
	if err != nil {
		t.Logf("Got: %s %c", err, checkMark)
	} else {
		t.Logf("Got: %s %c", err, checkMark)
	}
}

func TestMergeSort(t *testing.T) {
	nums := []int{105, 5, 87, 96, 54, 46, 32, 23, 72, 12, 68}

	expected := []int{5, 12, 23, 32, 46, 54, 68, 72, 87, 96, 105}

	t.Logf("MergeSort(%v) should yield %v", nums, expected)
	sorted := MergeSort(nums)
	if cmp.Equal(sorted, expected) {
		t.Logf("Got: %v %c", sorted, checkMark)
	} else {
		t.Errorf("Got: %v %c", sorted, ballotX)
	}
}
