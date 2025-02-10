func printParenthesisNumber(expn string, size int) {
	//Uncomment the ch variable and use it iterate through the string
	var ch byte 
	output := "" //use output variable to save and print the output string
	
	stk := new(Stack)
	num := 1
  
	
	for i := 0; i < size; i++ {
		  ch = expn[i]
	  if ch == '(' {
		stk.Push(num)
		output += fmt.Sprintf("%v",num)
		num++
	  } else if ch == ')' {
		  output += fmt.Sprintf("%v",stk.Pop())
	  }
	}
	  fmt.Println(output)
  }