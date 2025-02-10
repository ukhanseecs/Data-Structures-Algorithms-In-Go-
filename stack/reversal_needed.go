import "math" //You can "math" library to solve the problem

func reverseParenthesis(expn string, size int) int {
  
  //Implement your solution here

  //Uncomment this line of code and use ch to iterate in a string
  if size % 2 != 0 {
    return -1
  }
  
  
  var ch rune
  count := 0
  stk :=  new(Stack)
  

  for i:=0 ; i<size ; i++ {
      ch = rune(expn[i])
      if ch == '(' {
        stk.Push(ch)
      } else {
          if !stk.IsEmpty() && stk.Top() == '(' {
            stk.Pop()
          } else {
             count ++
          }
      }
  }

    len := stk.Length() // Remaining unmatched '('

    // Each two unmatched require one reversal
    return (len+1)/2 + (count+1)/2
}