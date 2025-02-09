func SumArr(data []int) int {

	total := 0

	//Implement your solution here
	for i := 0; i < len(data); i++ {
		total = total + data[i]
	}
	return total
}