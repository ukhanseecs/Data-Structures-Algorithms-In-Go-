package main

func (t *Tree) CeilBST(val int) int {
    
    ceil := math.MinInt32
    curr := t.root
  
    
    for (curr != nil) {
      if curr.value <= val {
        curr = curr.right
      } else if curr.value > val {
        ceil = curr.value
        curr = curr.left
      }
    }
    if ceil < 0 {
      return -1
    }
    //Implement your solution here

    return ceil
}
