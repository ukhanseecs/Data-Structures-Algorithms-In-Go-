func ArrayIndexMaxDiff(arr []int, size int) int {
	//Implement your solution here

	diff := 0
	maxdiff := 0
	for i := 0; i < size; i++ {
		for j := 1; j < size; j++ {
			if i < j && arr[i] < arr[j] {
				diff = j - i
				if diff > maxdiff {
					maxdiff = diff

				}
			}
		}
	}
	return maxdiff
}
