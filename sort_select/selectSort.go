package sort_select

func SelectSort(arr []int64) []int64 {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	for i := 0; i < length-1; i++ {
		index := i
		for j := i + 1; j < length; j++ {
			if arr[i] > arr[j] {
				index = j
			}
		}

		if index != i {
			tmp := arr[i]
			arr[i] = arr[index]
			arr[index] = tmp
		}
	}

	return arr
}
