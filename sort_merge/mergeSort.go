package sort_merge

func MergeSort(arr []int) []int {

	length := len(arr)
	if length <= 1 {
		return arr
	}

	m :=  length / 2
	c1 := MergeSort(arr[0:m])
	c2 := MergeSort(arr[m:length])

	return merge(c1, c2)

}

func merge(c1, c2 []int) []int {
	tmp := make([]int,0)
	i := 0
	j := 0
	for i < len(c1) && j < len(c2) {
		if c1[i] <= c2[j] {
			tmp = append(tmp, c1[i])
			i++
		} else {
			tmp = append(tmp, c2[j])
			j++
		}
	}

	if i < len(c1) {
		tmp = append(tmp, c1[i:]...)
	}

	if j < len(c2) {
		tmp = append(tmp, c2[j:]...)
	}

	return tmp
}
