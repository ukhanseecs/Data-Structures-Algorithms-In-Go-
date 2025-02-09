func maxDepthParenthesis(expn string, size int) int {
	//Implement your solution here
	
	//Uncomment the ch variable and use it to move in an expression array
	
	
	stk := new(Stack)
	count := 0
	maxCount := 0
	var ch byte 
	
	for i := 0 ; i< size; i ++ {
		ch = expn[i]
		if ch == '(' {
		  stk.Push(ch)
		  count ++
		} else if ch == ')' {
		  stk.Pop()
		  count --
		}
		if count > maxCount {
		  maxCount = count
		}
	}
	return maxCount //Return 0 if depth of parenthesis is zero
  }