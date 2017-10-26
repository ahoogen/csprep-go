package stack

// Elem contains data and a pointer to the next StackElem in the chain
// or nil if last element in the chain.
type Elem struct {
	next *Elem
	data interface{}
}

// New creates and returns a new *StackElem containing data.
func New(data interface{}) *Elem {
	return &Elem{nil, data}
}

// Push adds a new Elem to the top of stack e.
// Idiomatic with append(), Elem.Push() returns the address of the new
// head of the stack.
func (e *Elem) Push(data interface{}) *Elem {
	return &Elem{e, data}
}

// Pop moves the head pointer into stack and returns the element and head
// pointer for both the Popped stack element and the new stack head.
// Pop is idiomatic with append() function and expects the user to maintain
// a copy of the stack head pointer.
func (e *Elem) Pop() (*Elem, *Elem) {
	stack := e.next
	e.next = nil
	return e, stack
}
