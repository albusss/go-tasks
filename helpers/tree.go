package helpers

type Tree struct {
	Left  *Tree
	Right *Tree
	Val   int
}

func GetTree() *Tree {
	root := &Tree{Val: 1}
	root.Left = &Tree{Val: 2}
	root.Right = &Tree{Val: 3}
	root.Left.Left = &Tree{Val: 4}
	root.Left.Right = &Tree{Val: 5}
	root.Right.Left = &Tree{Val: 6}
	root.Right.Right = &Tree{Val: 7}

	return root
}
