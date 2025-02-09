package stack

func reverseStack(stk *Stack) {
	//Implement your solution here
	  if (stk.Length() == 0) {
		return
	  }
	  popped := stk.Pop()
	  reverseStack(stk)
	  bottomInsert(stk, popped)
	}