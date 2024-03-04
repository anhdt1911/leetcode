package main

import "fmt"

func main() {
	m := Constructor()
	m.Push(2)
	m.Push(0)
	m.Push(3)
	m.Push(0)
	m.Pop()
	m.Pop()
	m.Pop()
	fmt.Println(m.GetMin())
}

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (m *MinStack) Push(val int) {
	m.stack = append(m.stack, val)
	if len(m.minStack) == 0 || val < m.minStack[len(m.minStack)-1] {
		m.minStack = append(m.minStack, val)
	} else {
		m.minStack = append(m.minStack, m.minStack[len(m.minStack)-1])
	}
}

func (m *MinStack) Pop() {
	m.minStack = m.minStack[:len(m.minStack)-1]
	m.stack = m.stack[:len(m.stack)-1]
}

func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

func (m *MinStack) GetMin() int {
	return m.minStack[len(m.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
