package util

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var test = []Page{
		{"", 2}, {"", 1}, {"", 4}, {"", 3}, {"", 6},
		{"", 0}, {"", 9}, {"", 8}, {"", 7}, {"", 5},
	}

	result := MergeSort(test)
	fmt.Println(result)

}
