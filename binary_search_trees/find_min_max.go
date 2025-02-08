func (t *Tree) FindMinNode() *Node {
    
    var curr *Node = t.root
    
    for  curr.left != nil {
      curr = curr.left
    }
    return curr
}


//========================

func (t *Tree) FindMaxNode() *Node {
  var node *Node = t.root
  
  if node == nil {
    return nil
  }
  
  //Implement your solution here
  for node.right != nil {
    node = node.right
  }
  return node
}



