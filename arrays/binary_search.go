func BinarySearch(data []int, value int) bool {
	//Implement your solution here
	l := 0
	r := len(data) - 1
	for i := 0; i < len(data); i++ {
		middle := (l + r) / 2
		if data[middle] == value {
			return true
		} else if data[middle] < value {
			l = middle
		} else if data[middle] > value {
			r = middle
		}

	}
	return false //Return false if value not found
}
