func findDuplicateParenthesis(expn string, size int) bool {

    //Implement your solution here
    stk := new(Stack)
 
    for i:=0; i<size ; i++ {
      if expn[i] != ')' {
        stk.Push(int(expn[i]))
      } else {
        if !stk.IsEmpty() && stk.Top() == '(' {
          return true
        }
        
        operatorFound := false
        for !stk.IsEmpty() && stk.Top() != '(' {
            top := stk.Top()
            if top == '+' || top == '-' || top == '*' || top == '/' {
                operatorFound = true
            }
            stk.Pop()
        }
        
        if !stk.IsEmpty() {
          stk.Pop()
        }
          
        if !operatorFound {
            return true
        }
      }
    }
      
    return false

}