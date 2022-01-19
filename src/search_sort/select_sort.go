package search_sort

func SelectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		//swap
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
