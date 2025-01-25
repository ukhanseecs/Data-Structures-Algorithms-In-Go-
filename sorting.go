package arrays

func Partition01(arr []int, size int) int {
	//Write your code here
	l := 0
	r := size - 1
	count := 0

	for l < r {
		if arr[l] == 0 {
			l++
		} else if arr[r] == 1 {
			r--
		} else {
			//swap
			arr[l], arr[r] = arr[r], arr[l]
			count++
			l++
			r--
		}
	}

	return count
}

//================================================================================================

func Partition012(arr []int, size int) {
	//Implement your solution here
	l := 0
	m := 0
	h := size - 1

	for m <= h {
		if arr[m] == 0 {
			// swap l,m
			arr[l], arr[m] = arr[m], arr[l]
			l++
			m++
		} else if arr[m] == 1 {
			m++
		} else if arr[m] == 2 {
			//swap m,h
			arr[m], arr[h] = arr[h], arr[m]
			h--
		}
	}
}

//================================================================================================
