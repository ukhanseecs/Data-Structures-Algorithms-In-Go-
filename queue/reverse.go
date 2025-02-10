func reverseQueue(que *Queue) {
    //Implement your solution here
    stk := new(Stack)
    
    for que.Length() != 0 {
      stk.Push(que.Dequeue().(int))
    }
    for stk.Length() != 0 {
      que.Enqueue(stk.Pop())
    }  
  
}