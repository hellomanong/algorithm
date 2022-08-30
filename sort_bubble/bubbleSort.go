package sort_bubble

func BubbleSort(arr []int64) {
	length := len(arr)
	if length <= 1 {
		return
	}

	for i := 0; i < length-1; i++ {
		flag := false
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				//tmp := arr[j]
				//arr[j] = arr[j+1]
				//arr[j+1] = tmp

				arr[j], arr[j+1] = arr[j+1], arr[j]

				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
