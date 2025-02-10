func reverseKElementInQueue(que *Queue, k int) {
    //Implement your solution here
    stk := new(Stack)
    
    for i:=0 ; i<k ; i++ {
      stk.Push(que.Dequeue().(int))
    }
    for stk.Length() != 0 {
      que.Enqueue(stk.Pop())
    }
    not_k := que.Length() - k
    for i:=0 ; i<not_k ; i++ {
      ele := que.Dequeue()
      que.Enqueue(ele)
    }
}