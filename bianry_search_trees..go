func CreateBinarySearchTree(arr []int) *Tree {
    t := new(Tree)
    size := len(arr)
    t.root = createBinarySearchTreeUtil(arr, 0, size-1)
    return t
}

func createBinarySearchTreeUtil(arr []int, start int, end int) *Node { 
    if start > end {
      return nil
    }
    
    middle := (start+end)/2
    curr := new(Node)

    curr.value = arr[middle]
    curr.left = createBinarySearchTreeUtil(arr, start, middle-1)
    curr.right = createBinarySearchTreeUtil(arr, middle+1, end)
  
  //Implement your solution here

    return curr
}

//=================================================================================================================================
