func IsBST(root *Node) bool {
    //Implement your solution here
    if root== nil {
      return true
    }
    
    
    lft := root.left
    rht := root.right
    
    if root.value > lft.value && root.value < rht.value {
      return true
    } else {
      return false
    }
    
    return (IsBST(root.left) && IsBST(root.right))
}


//=================================================


func IsBST(root *Node) bool {
    //Implement your solution here
    if root== nil {
      return true
    }
    
    
    if root.left != nil && FindMax(root.left).value > root.value {
      return false
    }
    if root.right != nil && FindMin(root.right).value < root.value {
      return false 
    }
    
    return IsBST(root.left)&&IsBST(root.right)
}