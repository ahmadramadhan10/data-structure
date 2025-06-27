package stack

type Node struct {
	Val          int
	Next, Before *Node
}

type Stack struct {
	head *Node
	tail *Node
	size int
}

func (l *Stack) push(val int) {
	newNode := &Node{Val: val, Next: nil, Before: nil}
	if l.tail == nil {
		l.tail = newNode
		l.head = newNode
		l.size = 1
	} else {
		l.head.Next = newNode
		temp := l.head
		l.head = newNode
		l.head.Before = temp
		l.size += 1
	}
}

func (l *Stack) pop() {
	if l.head != nil {
		l.head = l.head.Before
		l.head.Next = nil
		l.size -= 1
	}
}

func (l *Stack) top() int {
	return l.head.Val
}
