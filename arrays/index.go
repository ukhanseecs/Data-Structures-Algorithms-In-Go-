func IndexArray(arr []int, size int) {
	//Implement your solution here
	temp := 0
	for i := 0; i < size; i++ {
		for arr[i] != -1 && arr[i] != i {
			temp = arr[arr[i]]
			arr[arr[i]] = arr[i]
			arr[i] = temp
		}
	}
	//Note: Please make changes in the given array
}
