package arrays
import "fmt"

func sortStack(stk *Stack) {
	//Implement your solution here
	// base case
	// if stk empty, return
	if stk.Length() > 0 {
		// recursive case
		// pop top
		popped := stk.Pop()
		// sortStack
		sortStack(stk)
		// insert top correctly
		sortedInsert(stk, popped)
	}
}