
func swap(data []int, x int, y int) {
	data[x], data[y] = data[y], data[x]
}
func WaveArray(arr []int) {
	//Implement your solution here
	for i := 1; i < len(arr); i += 2 {
		if arr[i] > arr[i-1] {
			swap(arr, i, i-1)
		}
		if i+1 < len(arr) && arr[i] > arr[i+1] {
			swap(arr, i, i+1)
		}
	}
	//Update the same array
}