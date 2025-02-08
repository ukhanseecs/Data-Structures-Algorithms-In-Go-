func (t *Tree) DeleteNode(value int) {
    t.root = DeleteNode(t.root, value)
}

func DeleteNode(node *Node, value int) *Node {
    //Implement your solution here
    if node == nil {
      return nil
    }
    
    if value < node.value {
      node.left = DeleteNode(node.left, value)
    } else if value > node.value {
      node.right = DeleteNode(node.right, value)
    } else { //value == node.value
      if node.left == nil && node.right == nil{
        return nil
      } else if node.left != nil && node.right == nil {
        return node.left
      } else if node.left == nil && node.right != nil {
        return node.right 
      } else {
        rep := FindMin(node.right)
        node.value = rep.value
        node.right = DeleteNode(node.right, rep.value)
      }
    }
    return node
}