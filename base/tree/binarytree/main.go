package main

import "fmt"

// Queue
type Queue struct {
	d []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		d: make([]interface{}, 0),
	}
}

func (q Queue) Len() int {
	return len(q.d)
}
func (q *Queue) Push(i interface{}) {
	q.d = append(q.d, i)
}
func (q *Queue) Pop() interface{} {
	if len(q.d) == 0 {
		return nil
	}

	m := q.d[0]
	if len(q.d) == 1 {
		q.d = make([]interface{}, 0)
		return m
	}
	q.d = q.d[1:]
	return m
}
func (q Queue) IsEmpty() bool {
	return len(q.d) == 0
}

// Node
type Node struct {
	data int
	l    *Node
	r    *Node
}

// NewNode new Node
func NewNode(data int) *Node {
	return &Node{
		data: data,
	}
}

func (n Node) String() string {
	return fmt.Sprintf("%d", n.data)
}

func (n *Node) SetData(data int) {
	n.data = data
}

func (n *Node) Print() {
	fmt.Printf("%+v  ", n)
}

// PrePrint 前序遍历
func (n *Node) PrePrint() {
	if n == nil {
		return
	}
	n.Print()
	n.l.PrePrint()
	n.r.PrePrint()
}

// MidPrint 中序遍历
func (n *Node) MidPrint() {
	if n == nil {
		return
	}
	n.l.MidPrint()
	n.Print()
	n.r.MidPrint()
}

// PostPrint 后序遍历
func (n *Node) PostPrint() {
	if n == nil {
		return
	}
	n.l.PostPrint()
	n.r.PostPrint()
	n.Print()
}

// Bforder 广度优先遍历
func (n *Node) Bforder() {
	q := NewQueue()
	q.Push(n)
	for {
		item := q.Pop()
		if node, ok := item.(*Node); ok {
			node.Print()
			if node.l != nil {
				q.Push(node.l)
			}
			if node.r != nil {
				q.Push(node.r)
			}
			if q.IsEmpty() {
				break
			}
		}
	}
}

// TreeHeight 树深
func (n *Node) TreeHeight() int {
	if n == nil {
		return 0
	}

	l := n.l.TreeHeight()
	r := n.r.TreeHeight()
	if l >= r {
		return l + 1
	}
	return r + 1
}

func main() {
	root := NewNode(0)
	root.l = NewNode(1)
	root.r = NewNode(2)
	root.l.l = NewNode(3)
	root.l.r = NewNode(4)
	root.r.l = NewNode(5)

	fmt.Println("前序遍历:")
	root.PrePrint()

	fmt.Println("\n中序遍历:")
	root.MidPrint()

	fmt.Println("\n后序遍历:")
	root.PostPrint()

	fmt.Println("\n层序遍历:")
	root.Bforder()

	fmt.Println("\n树高:", root.TreeHeight())
}
