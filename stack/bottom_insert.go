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