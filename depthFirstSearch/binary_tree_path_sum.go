package depthFirstSearch

import (
	"fmt"
	"github.com/albus/EduG/helpers"
)

func BinaryPathSum() {
	tree := helpers.GetTree()
	s := 10

	fmt.Println(search(tree, s))
}

func search(tree *helpers.Tree, s int) bool {

	if tree == nil {
		return false
	}

	val := tree.Val

	if val == s && tree.Left == nil && tree.Right == nil {
		return true
	}

	return search(tree.Left, s-val) || search(tree.Right, s-val)
}
