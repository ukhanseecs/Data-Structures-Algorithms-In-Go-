package dsa

import "fmt"

func TowerOfHanoi(num int, src byte, dst byte, temp byte) {
	// write some code here
	if num < 1 {
		return
	}
	TowerOfHanoi(num-1, src, temp, dst)
	fmt.Printf("Move %d disk from peg %c to peg %c ", num, src, dst)
	TowerOfHanoi(num-1, temp, dst, src)
	//DONOT change print statement
	// write some code here
}