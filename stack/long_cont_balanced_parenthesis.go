func longestContBalParen(str string, size int) int {
    //Implement your solution here
    stk := new(Stack)
    stk.Push(-1)
    maxLength :=  0
    
    for i:=0 ; i<size ; i++ {
      if str[i] == '(' {
        stk.Push(i)
      } else {
        stk.Pop()
        if stk.IsEmpty() {
          stk.Push(i)
        } else if !stk.IsEmpty() {
          if maxLength < i - stk.Top() {
            maxLength = i - stk.Top()
          }
        }
      }
    }
    return maxLength //Return 0 if size of str is zero
}