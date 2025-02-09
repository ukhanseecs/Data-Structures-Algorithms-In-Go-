
func RotateArray(data []int, k int) {
	//Implement your solution here

	// if array 1,2,3,4,5 and k =2
	k_elements := data[:k]                                //1,2
	remaining_data := data[k:]                            //3,4,5
	rotated_data := append(remaining_data, k_elements...) //3,4,5,1,2
	copy(data, rotated_data)

}