func SmallestPositiveMissingNumber(arr []int, size int) int {
	//Implement your solution here
	temp := 0
	for i := 0; i < size; i++ {
		for arr[i] != i+1 && arr[i] > 0 && arr[i] <= size {
			temp = arr[i]
			arr[i] = arr[arr[i]-1]
			arr[temp-1] = temp
		}
	}

	for i := 0; i < size; i++ {
		if arr[i] != i+1 {
			return i + 1
		}
	}
	//Implement your solution here
	return -1 //Return -1 if missing number not found
}
