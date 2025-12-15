package main

type page struct {
	url   string
	count int
}

func MergeSort(pages []page) []page {

	if len(pages) < 2 {
		return pages
	}
	mid := len(pages) / 2
	left := MergeSort(pages[:mid])
	right := MergeSort(pages[mid:])
	return merge(left, right)

}

func merge(left []page, right []page) []page {
	i, j := 0, 0

	result := make([]page, len(left)+len(right))

	for i < len(left) && j < len(right) {
		if left[i].count <= right[j].count {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i])
	result = append(result, right[j])

	return result
}
