package main

var array []struct{}

type Page struct {
	url string
	val int
}

func MergeSort(array []Page) []Page {
	var i = 0
	var j = 0
	var k = 0

	var mid = len(array)/2 - 1
	var left = array[:mid]
	var right = array[mid:]

	MergeSort(left)
	MergeSort(right)

	for i < len(left) && j < len(right) {
		if left[i].val < right[j].val {
			array[k] = left[i]
			i++
		} else {
			array[k] = right[j]
			j++
		}
		k++

	}

	return array
}
