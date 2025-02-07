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


func (t *Tree) PrintBreadthFirst() {
    
    //Implement your solution here
    queue := []*Node{t.root}
    
    for len(queue) > 0 {
      curr := queue[0]
      queue = queue[1:]
      
      fmt.Print(curr.value, " ")
      
      if curr.left != nil {
        queue = append(queue, curr.left)
      }
      
      if curr.right != nil {
        queue = append(queue, curr.right)
      }
    }
}



//=================================================================================================================================



func (t *Tree) PrintDepthFirst(){

    //Implement your solution here
    stack := []*Node{t.root}
    
    for len(stack) > 0 {
      curr := stack[len(stack)-1]
      stack = stack[:len(stack)-1]
      
      fmt.Print(curr.value, " ")
      
      if curr.right != nil {
        stack = append(stack, curr.right)
      }
            
      if curr.left != nil {
        stack = append(stack, curr.left)
      }
    }
}



//=================================================================================================================================

// print level order line by line

func (t *Tree) PrintLevelOrderLineByLine() {
    //Implement your solution here
    queue1 := []*Node{t.root}
    queue2 := []*Node{}
    
    
    for len(queue1) > 0 {
      curr := queue1[0]
      queue1 = queue1[1:]
     
      fmt.Print(curr.value, " ")
   
  		if curr.left != nil {
  			queue2 = append(queue2, curr.left)
  		}
  		if curr.right != nil {
  			queue2 = append(queue2, curr.right)
  		}
  		
      if len(queue1) == 0 {
        fmt.Print("; ") 
        queue1, queue2 = queue2, queue1
      }
    }
    fmt.Println("; ")
}




//=================================================================================================================================

func (t *Tree) PrintLevelOrderLineByLine2() {
  //Implement your solution here
  queue1 := []*Node{t.root, nil}
  
  for len(queue1) > 0 {
    curr := queue1[0]
    queue1 = queue1[1:]
   
 
    if curr == nil {
          fmt.Print("; ")
          if len(queue1) > 0 { 
              queue1 = append(queue1, nil)
          }
          continue
      }
    
    fmt.Print(curr.value, " ")
    
    if curr.left != nil {
      queue1 = append(queue1, curr.left)
    }
    if curr.right != nil {
      queue1 = append(queue1, curr.right)
    }
  }
  fmt.Print(";")
}




//================================================================================================================


func (t *Tree) PrintSpiralTree() {
  stack1 := []*Node{t.root}
  stack2 := []*Node{}
  
  for len(stack1) > 0 || len(stack2) > 0{

    for len(stack1) > 0 {
      curr := stack1[len(stack1)-1]
      stack1 = stack1[:len(stack1)-1]
      fmt.Print(curr.value, " ")
      
      if curr.left != nil {
        stack2 = append(stack2, curr.left)
      }
      if curr.right != nil {
        stack2 = append(stack2, curr.right)
      }
    }
    
    if len(stack1) ==0 || len(stack2) == 0 {
        fmt.Print("; ")
    }
    
    for len(stack2) > 0 {
      curr := stack2[len(stack2)-1]
      stack2 = stack2[:len(stack2)-1]
      fmt.Print(curr.value, " ")
      
      
      if curr.right != nil {
        stack1 = append(stack1, curr.right)
      }
      if curr.left != nil {
        stack1 = append(stack1, curr.left)
      }
    }
    
    if len(stack1) ==0 || len(stack2) == 0 {
        fmt.Print("; ")
    }
  }
  fmt.Print("; ")
}




//================================================================================================================



func (t *Tree) NthPreOrder(index int) {
    var counter int
    nthPreOrder(t.root, index, &counter)
}

func nthPreOrder(node *Node, index int, counter *int) {
   //Implement your solution here
    if node != nil {
      (*counter)++
      if *counter == index {
        fmt.Println(node.value)
      }
      
      nthPreOrder(node.left,index, counter)
      nthPreOrder(node.right, index, counter)
}}



//================================================================================================================


func (t *Tree) NthPostOrder(index int) {
  var counter int
  nthPostOrder(t.root, index, &counter)
}

func nthPostOrder(node *Node, index int, counter *int) {
  //Implement your solution here
  
  if node != nil {
    nthPostOrder(node.left, index, counter)
    nthPostOrder(node.right, index, counter)
    (*counter)++
    if index == *counter {
      fmt.Println(node.value)
    }
  }
  


  //Implement your solution here
}



//================================================================================================================


func (t *Tree) NthInOrder(index int) {
  var counter int
  nthInOrder(t.root, index, &counter)
}

func nthInOrder(node *Node, index int, counter *int) {
  //Implement your solution here
  if node != nil {
    nthInOrder(node.left, index, counter)
    (*counter)++
    if index == *counter {
      fmt.Println(node.value)
    }
    nthInOrder(node.right, index, counter)
  }

  //Implement your solution here
}



//================================================================================================================

func (t *Tree) PrintAllPath() {
  stk := new(Stack)
  printAllPath(t.root, stk)
}

func printAllPath(curr *Node, stk *Stack) {
  //Implement your solution here
  if curr == nil {
      return
  }
  
  stk.Push(curr.value)
  if curr.left == nil && curr.right == nil {
    stk.Print()
    fmt.Print("; ")
    stk.Pop()
    return
  }
  printAllPath(curr.right, stk)
  printAllPath(curr.left, stk)
  stk.Pop()
}

//================================================================================================================



func (t *Tree) NumNodes() int {
  return numNodes(t.root)
}

func numNodes(curr *Node) int {
  //Implement your solution here
  count := 0
  if curr == nil {
    return 0
  }
  count ++

  return count +  numNodes(curr.left) + numNodes(curr.right) //Kindly change the return value as per your needs
}



//================================================================================================================


func (t *Tree) SumAllBT() int {
  return sumAllBT(t.root)
}

func sumAllBT(curr *Node) int {
  //Implement your solution here
  sum := 0
  if curr == nil {
    return 0
  }
  sum += curr.value
  
  return sum + sumAllBT(curr.left) + sumAllBT(curr.right) //Kindly change the return value as per need
}


///================================================================================================================









