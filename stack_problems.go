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


type QueueUsingStack struct {
	stk1 Stack
	stk2 Stack
	// add more values here if need be
   
   }
   
   func (que *QueueUsingStack) Add(value int) {
	 //Implement your solution here
	 que.stk1.Push(value)
   }
   
   func (que *QueueUsingStack) Remove() int {
	 //Implement your solution here
	 var value int
	 if que.stk2.IsEmpty() == false {
	   value = que.stk2.Pop()
	   return value
	 }
	 for que.stk1.IsEmpty() == false {
	   value = que.stk1.Pop()
	   que.stk2.Push(value)
	 }
	 value = que.stk2.Pop()
	 return value
   }
   
   func (que *QueueUsingStack) Length() int {
	 //Implement your solution here
	 return (que.stk1.Length() + que.stk2.Length())
   }
   
   func (que *QueueUsingStack) IsEmpty() bool {
	 //Implement your solution here
	 if (que.stk1.IsEmpty() ==true && que.stk2.IsEmpty() ==true) {
	   return true
	 } else {
	   return false
	 }
   } 

   

   