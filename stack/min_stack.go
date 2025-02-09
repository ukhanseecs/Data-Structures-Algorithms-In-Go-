
type MinStack struct {
    maxSize int
	// Initialize your data structures here
	mainstk Stack
	minstk Stack
}

//removes and returns value from stack
func (m *MinStack) Pop() int{
	// write your code here
    _ = m.minstk.Pop()
    top := m.mainstk.Top()
    _ = m.mainstk.Pop()
    return top
}

//Pushes value into the stack
func (m *MinStack) Push(value int){
	// write your code here
	//1. Push value in mainStack and check value with the top value of minStack
	//2. If value is smaller than top, then Push top in minStack
	//else Push value in minStack
	m.mainstk.Push(value)
	if value > m.minstk.Top() && !m.minstk.IsEmpty() {
	  m.minstk.Push(m.minstk.Top())
	} else {
	  m.minstk.Push(value)
	}
}

//returns minimum value in O(1)
func (m *MinStack) Min() int{
	// write your code here 
    return m.minstk.Top()
}