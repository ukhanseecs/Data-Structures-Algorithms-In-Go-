func Partition01(arr []int, size int) int {
	//Write your code here
	l := 0
	r := size-1
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
		}
	  }
	}

	return count
  }


  //==============================================



