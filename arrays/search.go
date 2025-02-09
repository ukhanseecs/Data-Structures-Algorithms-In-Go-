func SequentialSearch(data []int, value int) bool {
	//Implement your solution here

	for i := 0; i < len(data); i++ {
		if value == data[i] {
			return true
		}
	}
	return false //Return false if value not found
}