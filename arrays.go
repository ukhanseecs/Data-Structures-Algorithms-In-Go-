package arrays

func SumArr(data []int) int {

  total := 0

	//Implement your solution here
	for i:=0; i<len(data) ; i++ {
	  total = total + data[i]
	}
	return total
  }



func SequentialSearch(data []int, value int) bool {
    //Implement your solution here

  for i := 0; i < len(data) ; i++ {
    if value == data[i]{
      return true
    }
  }
    return false     //Return false if value not found
}


func BinarySearch(data []int, value int) bool {
    //Implement your solution here
    l := 0
    r := len(data) -1
    for i :=0; i<len(data); i++ {
      middle := (l+r)/2
      if data[middle]==value {
        return true
      } else if data[middle] < value {
        l = middle
      } else if data[middle] > value {
        r = middle
      }

    }
    return false     //Return false if value not found
}



func MaxSubArraySum(data []int) int {

    //Implement your solution here
    max := 0
    sum := 0

    for i := 0; i < len(data); i++ {
        sum = sum + data[i]
        if sum < 0{
          sum = 0
        }
        if sum > max {
          max = sum
        }
    }
    return max
}




func RotateArray(data []int, k int) {
    //Implement your solution here

  	// if array 1,2,3,4,5 and k =2
    k_elements := data[:k]     //1,2
    remaining_data := data[k:]   //3,4,5
    rotated_data := append(remaining_data, k_elements...) //3,4,5,1,2
    copy(data, rotated_data)

  }





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





func IndexArray(arr []int, size int) {
    //Implement your solution here
  temp := 0
  for i := 0; i<size ; i++ {
    for arr[i] != -1 && arr[i] != i{
      temp = arr[arr[i]]
      arr[arr[i]] = arr[i]
      arr[i] = temp
    }
  }
   //Note: Please make changes in the given array
}




func Sort1toN(arr []int, size int) {
	temp := 0
	for i := 0; i<size ; i++ {
	  for arr[i] != i+1{
		temp = arr[i]
		arr[i] = arr[arr[i]-1]
		arr[temp-1] = temp
	  }
	}
	//Implement your solution here

  }



  func SmallestPositiveMissingNumber(arr []int, size int) int {
    //Implement your solution here
    temp := 0
    for i := 0; i<size ; i++ {
      for arr[i] != i+1 && arr[i] > 0 && arr[i] <= size{
        temp = arr[i]
        arr[i] = arr[arr[i]-1]
        arr[temp-1] = temp
      }
    }

    for i := 0; i<size ; i++{
      if arr[i] != i+1 {
        return i+1
      }
    }
  //Implement your solution here
    return -1 //Return -1 if missing number not found
}






func MaxMinArr(arr []int, size int) {
	//Implement your solution here
	max := size-1
	min := 0
	arr_copy := make([]int, size)
	copy(arr_copy, arr)
	 for i:= 0; i < size; i++ {
	   if i%2 == 0 {
		 arr[i] = arr_copy[max]
		 max -= 1
	   } else {
		  arr[i] = arr_copy[min]
		  min +=1
	   }
	 }
  }






  func ArrayIndexMaxDiff(arr []int, size int) int {
	//Implement your solution here

	diff := 0
	maxdiff := 0
	for i:=0; i <size; i++ {
	  for j:=1; j < size; j++ {
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





