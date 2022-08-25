package sort_insert

func InsertSort(arr []int64) []int64 {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	for i := 1; i < length; i++ {
		value := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] < value {
				break
			}
			arr[j+1] = arr[j]
		}

		arr[j+1] = value
	}

	return arr
}
