package arrays
import "fmt"

type StackInt struct {
    s []int
}

//isEmpty() function returns true if stack is empty or false in all other cases.
func (s *StackInt) IsEmpty() bool {
    //Implement your solution here
    length := len(s.s)
    return length == 0
    //Kindly make changes according to your needs
}

//length() function returns the number of elements in the stack.
func (s *StackInt) Length() int {
    //Implement your solution here
    length := len(s.s)
    return length
    //Kindly make changes according to your needs
}

//The print function will print the elements of the array.
func (s *StackInt) Print() {
    //Implement your solution here
    length := len(s.s)
    for i:=0 ; i<length ; i++ {
      fmt.Print(s.s[i]," ")
    }
    fmt.Println()
}

//push() function append value to the data.
func (s *StackInt) Push(value int) {
    //Implement your solution here
    s.s = append(s.s, value)
}

/* In the pop() function, first it will check that the stack is not empty.
Then it will pop the value from the data array and return it. */

func (s *StackInt) Pop() int {
    //Implement your solution here
    length := len(s.s)
    if length != 0 {
          res := s.s[length-1]
          s.s = s.s[:length-1]
          return res
    } else {
      return 0
    }
}

/*top() function will first check that the stack is not empty
then returns the value stored in the top element
of stack (does not remove it). */
func (s *StackInt) Top() int {
    //Implement your solution here
    length := len(s.s)
    if length != 0 {
        res := s.s[length-1]
        return res
    } else {
        return 0
    }
}
