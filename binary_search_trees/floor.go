package main


func (t *Tree) FloorBST(val int) int {

    floor := math.MaxInt32
    curr := t.root
    //Implement your solution here
    for (curr != nil) {
      if curr.value >= val {
        curr = curr.left
      } else {
        floor = curr.value
        curr = curr.right
      }
    } 
    
    if floor < 0{
      return -1
    }
    return floor
}
