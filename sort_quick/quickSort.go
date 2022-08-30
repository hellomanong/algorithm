package sort_quick

func QuickSort(arr []int, p, r int) []int {
	if p >= r-1 {
		return arr
	}

	q := partition(arr, p, r-1)
	QuickSort(arr, p, q)
	QuickSort(arr, q+1, r)

	return arr
}

func partition(arr []int, p, r int) int {
	value := arr[r]
	i := p
	for ; p < r; p++ {
		if arr[p] < value {
			if i != p {
				arr[p], arr[i] = arr[i], arr[p]
			}
			i++
		}
	}
	arr[r], arr[i] = arr[i], arr[r]
	return i
}
