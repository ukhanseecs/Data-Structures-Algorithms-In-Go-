package arrays
import "fmt"

type Node struct {
    value int
    next  *Node
}

type StackLinkedList struct {
    head *Node
    size int
}

//Size() function will return the size of the linked list.
func (s *StackLinkedList) Size() int {
    //Implement your solution here
    return s.size //Return size
}

/* IsEmpty() function will return true is size of the linked list is
equal to zero or false in all other cases. */
func (s *StackLinkedList) IsEmpty() bool {
    //Implement your solution here
    return s.size == 0 //Return true if stack is empty
}

/*First, the Peek() function will check if the stack is empty.
If not, it will return the peek value of stack i.e., will return the
head value of the linked list. */
func (s *StackLinkedList) Peek() (int, bool) {
    //Implement your solution here
    if s.size == 0 {
      return 0, true
    } else {
      return s.head.value, false
    }
}

//Push() function  will add new value at the start of the linked list.
func (s *StackLinkedList) Push(value int) {
    //Implement your solution here
    s.head = &Node{value, s.head}
    s.size++
}

/*In the pop() function, first it will check that the stack is not empty.
Then it will pop the value from the linked list and return it. */
func (s *StackLinkedList) Pop() (int, bool) {
    //Implement your solution here
    if s.size == 0 {
      return 0, true //Return true if stack is empty
    } else {
      value := s.head.value
      s.head = s.head.next
      s.size--
      return value, false

    }
}

/* Print() function will print the elements of the linked list. */
func (s *StackLinkedList) Print() {
    //Implement your solution here
    for s.head != nil {
      fmt.Print(s.head.value, " ")
      s.head = s.head.next
    }
    fmt.Println()
}
