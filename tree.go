// complete binary tree using dfs and queue


func LevelOrderBinaryTree(arr []int) *Tree {
    tree := new(Tree)
    tree.root = levelOrderBinaryTree(arr, 0, len(arr))
    return tree
}

func levelOrderBinaryTree(arr []int, start int, size int) *Node {
    //Implement your solution here
    root := &Node{value: arr[0]}
    queue := []*Node{root}
    
    i := 1
    for i < len(arr) {
      curr := queue[0]
      queue = queue[1:]
      
      // assign left
      if i < len(arr) {
        curr.left = &Node{value: arr[i]}
        queue = append(queue, curr.left)
        i++
      }
      
      if i < len(arr) {
        curr.right = &Node{value: arr[i]}
        queue = append(queue, curr.right)
        i++
      }
    }
    
    return root
}



//=================================================================================================================================


func (t *Tree) PrintPreOrder() {
    printPreOrder(t.root)
}

func printPreOrder(n *Node) {

    //Implement your solution here
    if n == nil {
      return
    }
    fmt.Print(n.value, " ")
    printPreOrder(n.left)
    printPreOrder(n.right)

    //Implement your solution here
    
}


//=================================================================================================================================


func (t *Tree) PrintPostOrder() {
    printPostOrder(t.root)
}

func printPostOrder(n *Node) {

    //Implement your solution here
    if n == nil {
      return
    }
    printPostOrder(n.left)
    printPostOrder(n.right)
    fmt.Print(n.value, " ")

    //Implement your solution here

}



//=================================================================================================================================


func (t *Tree) PrintInOrder() {
    printInOrder(t.root)
}


func printInOrder(n *Node) {
    
    //Implement your solution here
    if n == nil {
      return
    }
    printInOrder(n.left)
    fmt.Print(n.value, " ")
    printInOrder(n.right)


    //Implement your solution here
}



//=================================================================================================================================



