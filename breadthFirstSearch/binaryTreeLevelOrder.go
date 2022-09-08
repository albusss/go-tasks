package breadthFirstSearch

import (
	"fmt"
	"github.com/albus/EduG/helpers"
)

//
//type node struct {
//	left *node
//	right *node
//	val int
//}

func BinaryTree() {
	root := &helpers.Tree{Val: 12}
	root.Left = &helpers.Tree{Val: 7}
	root.Right = &helpers.Tree{Val: 1}
	root.Left.Left = &helpers.Tree{Val: 9}
	root.Right.Left = &helpers.Tree{Val: 10}
	root.Right.Right = &helpers.Tree{Val: 5}

	fmt.Println(traversal(root))
}

func traversal(root *helpers.Tree) [][]int {
	var result [][]int

	if root == nil {
		return result
	}
	var queue helpers.Queue
	queue.Enqueue(root)

	for len(queue.Q) > 0 {
		levelSize := len(queue.Q)
		currentLevel := []int{}
		for i := 0; i < levelSize; i++ {
			val2 := queue.Denqueue().(*helpers.Tree)
			currentLevel = append(currentLevel, val2.Val)

			if val2.Left != nil {
				queue.Enqueue(val2.Left)
			}
			if val2.Right != nil {
				queue.Enqueue(val2.Right)
			}
		}
		result = append(result, currentLevel)
	}

	return result
}
