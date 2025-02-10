func isBSTArray( preorder[] int, size int) bool {
  stk := new(Stack)
  var value int
  root := -1;
  for i := 0; i < size; i++ {
      value = preorder[i]
      if (value < root){
          return false
      }

      for (stk.Length() > 0 && stk.Top() < value){
          root = stk.Pop()
      }
      stk.Push(value)
  }
  return true
}