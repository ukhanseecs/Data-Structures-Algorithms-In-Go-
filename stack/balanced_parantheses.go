func IsBalancedParenthesis(expn string) bool {
    //Implement your solution here
    stk := new(Stack)
    
    for _, char := range expn {
      
      if (char == '{' || char == '(' || char == '[') {
        stk.Push(char)
      } else if (char == '}' || char == ')' || char == ']') {
          if stk.IsEmpty() {
           return false
          }
          
          ele := stk.Pop() 
          
          if (ele == '(' && char != ')') ||
  				(ele == '{' && char != '}') ||
  				(ele == '[' && char != ']') {
  				  return false
			    }
      }
    } 
    return stk.IsEmpty() //Return true if parentheses are balanced
}