package queue

import "fmt"

type Node struct {
	Val          int
	Next, Before *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (l *LinkedList) push(val int) {
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

func (l *LinkedList) pop() {
	if l.head != nil {
		l.head = l.head.Before
		l.head.Next = nil
		l.size -= 1
	}
}

func main() {
	linkedlist := LinkedList{}
	linkedlist.push(5)
	linkedlist.push(10)
	linkedlist.pop()
	linkedlist.push(9)
	cur := linkedlist.tail
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

// berarti pada linked list ketika update menggunakan O(n)
