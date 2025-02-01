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

//=================================================================================================================================

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

//=================================================================================================================================


func bottomInsert(stk *Stack, element int) {
    //Implement your solution here
    if (stk.Length() == 0) {
      stk.Push(element)
    } else {
      popped := stk.Pop()
      bottomInsert(stk, element)
      stk.Push(popped)
    }

}


//=================================================================================================================================



func reverseStack(stk *Stack) {
	//Implement your solution here
	  if (stk.Length() == 0) {
		return
	  }
	  popped := stk.Pop()
	  reverseStack(stk)
	  bottomInsert(stk, popped)
	}


//=================================================================================================================================


