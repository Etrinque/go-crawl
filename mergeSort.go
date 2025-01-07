package main

type Page struct {
	url string
	val int
}

func MergeSort(array []Page) []Page {

	if len(array) < 2 {
		return array
	}

	left := MergeSort(array[:len(array)/2])
	right := MergeSort(array[len(array)/2:])
	return merge(left, right)

}

func merge(left []Page, right []Page) []Page {
	var i = 0
	var j = 0

	var tmp []Page

	for i < len(left) && j < len(right) {
		if left[i].val < right[j].val {
			tmp = append(tmp, left[i])
			i++
		} else {
			tmp = append(tmp, right[j])
			j++
		}
	}
	for ; i < len(left); i++ {
		tmp = append(tmp, left[i])
	}
	for ; j < len(right); j++ {
		tmp = append(tmp, right[j])
	}

	return tmp
}
