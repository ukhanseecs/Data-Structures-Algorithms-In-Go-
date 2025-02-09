package arrays

func sortedInsert(stk *Stack, element int) {
	//Implement your solution here
	// base case
	top := stk.Top()
	if stk.Length() == 0 || top < element {
		stk.Push(element)
	} else {
		// recursive case
		popped := stk.Pop()
		sortedInsert(stk, element)
		stk.Push(popped)
	}
}
