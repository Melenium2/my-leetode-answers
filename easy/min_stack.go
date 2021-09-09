package easy

import "math"

// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
//
// Implement the minStack class:
//
// minStack() initializes the stack object.
// void push(val) pushes the element val onto the stack.
// void pop() removes the element on the top of the stack.
// int top() gets the top element of the stack.
// int getMin() retrieves the minimum element in the stack.
//
//
// Example 1:
//
// Input
// ["minStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
//
// Output
// [null,null,null,null,-3,null,0,-2]
//
// Explanation
// minStack minStack = new minStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin(); // return -3
// minStack.pop();
// minStack.top();    // return 0
// minStack.getMin(); // return -2
//
// Link to answer
// https://www.geeksforgeeks.org/design-a-stack-that-supports-getmin-in-o1-time-and-o1-extra-space/
type minStack struct {
	stack []int
	min   int
}

func constructor() minStack {
	return minStack{
		stack: make([]int, 0),
		min:   math.MaxInt32,
	}
}

func (ms *minStack) Push(val int) {
	if len(ms.stack) == 0 {
		ms.min = val
		ms.stack = append(ms.stack, val)

		return
	}

	if ms.min > val {
		lastMin := ms.min
		ms.min = val
		val = 2*val - lastMin
	}

	ms.stack = append(ms.stack, val)
}

func (ms *minStack) Pop() {
	last := len(ms.stack)-1
	lastElem := ms.stack[last]

	if lastElem < ms.min {
		ms.min = 2*ms.min - lastElem
	}

	ms.stack = ms.stack[:last]
}

func (ms *minStack) Top() int {
	elem := ms.stack[len(ms.stack)-1]

	if elem < ms.min {
		return ms.min
	}

	return elem
}

func (ms *minStack) GetMin() int {
	return ms.min
}