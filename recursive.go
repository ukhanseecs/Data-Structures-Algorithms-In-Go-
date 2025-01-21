package arrays

import (
	"fmt"
	"strconv"
)

//================================================================================================


func Factorial(i int) int {
    // Implement your solution here
    if i<=1 {
      return 1
    }
    return i * Factorial(i-1)
}


//================================================================================================


func PrintInt(number int) {
    // Write recursive code here
    if number == 0  {
      return
    }
    char := "0123456789ABCDEF"
    PrintInt(number /16)
    fmt.Print(string(char[number%16]))
    // Do not remove the print statement although you can change it with your initialized variables.
}


//================================================================================================


func GCD(m int, n int) int {
    //Implement your solution here
    if n == 0 {
      return m
    }
    rem := m%n
    return GCD(n, rem)
}


//================================================================================================


func Fibonacci(n int) int {
    //Implement your solution here
    // base case
    if n == 0 {
      return 0
    } else if n==1{
      return 1
    }
    return Fibonacci(n-1) + Fibonacci(n-2)
}


//================================================================================================




func Permutation(data []int, i int, length int) {

	// base condition DO NOT ALTER IT
	if length == i {
		temp := "{"
		for k := 0 ; k <length ; k++{
			temp += strconv.Itoa(data[k])
			temp += " "
		}
		temp += "} "
		fmt.Print(temp)
	  return
	 }
	  // Write your code here
	  for j := i; j < length; j++ {
		data[i], data[j] = data[j], data[i] // swap
		Permutation(data, i+1, length)
		data[i], data[j] = data[j],data[i] //backtrack
	  }
  }



