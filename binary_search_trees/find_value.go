func (t *Tree) Find(value int) bool {
    //Implement your solution here
    var curr *Node = t.root
    for curr != nil {
      if curr.value == value {
        return true
      } else {
        if curr.value < value {
          curr = curr.right
        } else {
          curr = curr.left
        }
      }
    }
    return false

}