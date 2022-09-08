package helpers

type Node struct {
	Next *Node
	Val  int
}

func GetList() *Node {
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2}
	head.Next.Next = &Node{Val: 3}
	head.Next.Next.Next = &Node{Val: 4}
	head.Next.Next.Next.Next = &Node{Val: 5}
	head.Next.Next.Next.Next.Next = &Node{Val: 6}
	head.Next.Next.Next.Next.Next.Next = &Node{Val: 7}
	head.Next.Next.Next.Next.Next.Next.Next = &Node{Val: 8}
	head.Next.Next.Next.Next.Next.Next.Next.Next = &Node{Val: 9}

	return head
}

func GetListWithCycleAt4() *Node {
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2}
	head.Next.Next = &Node{Val: 3}
	head.Next.Next.Next = &Node{Val: 4}
	head.Next.Next.Next.Next = &Node{Val: 5}
	head.Next.Next.Next.Next.Next = &Node{Val: 6}
	head.Next.Next.Next.Next.Next.Next = &Node{Val: 7}
	head.Next.Next.Next.Next.Next.Next.Next = &Node{Val: 8}
	head.Next.Next.Next.Next.Next.Next.Next.Next = head.Next.Next.Next

	return head
}
