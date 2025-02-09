func (t *Tree) TrimOutsidedataRange(min int, max int) {
    t.root = trimOutsidedataRange(t.root, min, max)
}

func trimOutsidedataRange(curr *Node, min int, max int) *Node {
    
    if curr == nil {
      return nil 
    }
    
    if curr.value < min {
      // delete curr and all on its left
      return trimOutsidedataRange(curr.right, min, max)
    } else if curr.value > max {
      // delete curr and all on its right
      return  trimOutsidedataRange(curr.left, min, max)
    }


    curr.left = trimOutsidedataRange(curr.left, min, max)
    curr.right = trimOutsidedataRange(curr.right, min, max)
    
    return curr
    
}
