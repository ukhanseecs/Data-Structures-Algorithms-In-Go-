package main
import "fmt"


func (t *Tree) PrintDataInRange(min int, max int) {
    printDataInRange(t.root, min, max)
}

func printDataInRange(root *Node, min int, max int) {
    if root == nil {
      return
    }
    if root.value < min {
      printDataInRange(root.right, min, max)
    } else if root.value > max {
      printDataInRange(root.left, min, max)
    } 
    if root.value >= min && root.value <= max {
      printDataInRange(root.left, min, max)
      fmt.Print(root.value," ")
      printDataInRange(root.right, min, max)
    }
    

    //Implement your solution here
}
