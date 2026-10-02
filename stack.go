package main

// en go, se utiliza slice para implementar stacks.
// inicializar el stack con un slice vacio de ints.
type MyStack struct {
	elements []int
}

func NewStack() *MyStack {
	return &MyStack{elements: []int{}}
}

func (stack *MyStack) isEmpty() bool {
	return len(stack.elements) == 0
}

func (stack *MyStack) push(item int) {
	stack.elements = append(stack.elements, item)
}

func (stack *MyStack) pop() (int, bool) {
	if stack.isEmpty() {
		return 0, false
	}
	var index int
	var item int
	index = len(stack.elements) - 1
	item = stack.elements[index]
	stack.elements[index] = 0               //libera la referencia; el array subyacente conserva la celda mas alla de len
	stack.elements = stack.elements[:index] //slice desde el inicio hasta el index del ultimo elemento
	return item, true
}

func (stack *MyStack) peek() (int, bool) {
	if stack.isEmpty() {
		return 0, false
	}
	var index int
	index = len(stack.elements) - 1
	return stack.elements[index], true
}
