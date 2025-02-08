func (t *Tree) FindMinNode() *Node {
    
    var curr *Node = t.root
    
    for  curr.left != nil {
      curr = curr.left
    }
    return curr
}