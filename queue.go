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

   

//=================================================================================================================================





