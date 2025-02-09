func MaxMinArr(arr []int, size int) {
	//Implement your solution here
	max := size - 1
	min := 0
	arr_copy := make([]int, size)
	copy(arr_copy, arr)
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			arr[i] = arr_copy[max]
			max -= 1
		} else {
			arr[i] = arr_copy[min]
			min += 1
		}
	}
}
