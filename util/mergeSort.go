package util

type Page struct {
	Url   string
	Count int
}

func MergeSort(pages []Page) []Page {

	if len(pages) < 2 {
		return pages
	}
	mid := len(pages) / 2
	left := MergeSort(pages[:mid])
	right := MergeSort(pages[mid:])
	return merge(left, right)

}

func merge(left []Page, right []Page) []Page {
	i, j := 0, 0

	result := make([]Page, len(left)+len(right))

	for i < len(left) && j < len(right) {
		if left[i].Count <= right[j].Count {
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
