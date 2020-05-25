package goqs

//HTMLStack is a stack data structure that can hold HTMLElement objects
type HTMLStack struct {
	top  *HTMLStackNode
	size int
}

//HTMLStackNode is a linked list node that is used in HTMLStack
type HTMLStackNode struct {
	element HTMLElement
	next    *HTMLStackNode
}

//EmptyStackError is an error type that is thrown when an element from
//a HTMLStack is attempted to be accessed from an empty stack
type EmptyStackError struct {
}

func (e EmptyStackError) Error() string {
	return "Removing from Stack was attempted when it was empty"
}

//MakeStack returns a default HTMLStack object with a nil root and 0 length
func MakeStack() HTMLStack {
	return HTMLStack{nil, 0}
}

//Push adds onto the top of the stack
func (stack *HTMLStack) Push(element HTMLElement) {
	node := HTMLStackNode{element, nil}
	node.next = stack.top
	stack.top = &node
	stack.size++
}

//Pop removes the top element and returns it
func (stack *HTMLStack) Pop() (HTMLElement, error) {
	if stack.size == 0 {
		return HTMLElement{}, EmptyStackError{}
	}

	node := stack.top
	stack.top = node.next
	stack.size--
	return node.element, nil
}

//Top returns the top element if it, or gives an error if stack is empty
func (stack *HTMLStack) Top() (HTMLElement, error) {
	if stack.size == 0 {
		return HTMLElement{}, EmptyStackError{}
	}
	return stack.top.element, nil
}

//Empty returns whether or not the stack is empty
func (stack *HTMLStack) Empty() bool {
	return stack.size == 0
}

//Size returns the number of elements in the current stack
func (stack *HTMLStack) Size() int {
	return stack.size
}
