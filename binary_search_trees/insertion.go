func (t *Tree) Add(value int) {
    t.root = addNode(t.root, value)
}

func addNode(n *Node, value int) *Node {
    //Implement your solution here
    if n == nil {
      n = new(Node)
      n.value = value
      return n
    } 
    if value < n.value {
      n.left = addNode(n.left, value)
    } else {
      n.right = addNode(n.right, value)
    }
    return n
}
