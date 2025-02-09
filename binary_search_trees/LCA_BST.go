package main

func (t *Tree) LcaBST(first int, second int) (int, bool) {
    return LcaBST(t.root, first, second)
}

func LcaBST(curr *Node, first int, second int) (int, bool) {
    //Implement your solution here
    if curr == nil {
      return 0, false
    }
    if first < curr.value && second < curr.value {
      return LcaBST(curr.left, first, second)
    } else if first > curr.value && second > curr.value {
      return LcaBST(curr.right, first, second)
    } else {
      return curr.value, true
    }
}