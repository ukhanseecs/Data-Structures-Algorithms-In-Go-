package stack

func reverseKElementInStack(stk *Stack, k int) {
    //Implement your solution here
    queue := new(Queue)
    
    for k != 0 {
      k--
      queue.Enqueue(stk.Pop())
    }
    for (queue.Length()!=0) {
      stk.Push(queue.Dequeue().(int))
    }
}
