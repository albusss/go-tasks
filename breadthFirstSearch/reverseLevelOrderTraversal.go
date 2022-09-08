package breadthFirstSearch

import (
	"fmt"
	"github.com/albus/EduG/helpers"
)

func ReverseLevel() {
	root := &helpers.Tree{Val: 1}
	root.Left = &helpers.Tree{Val: 2}
	root.Right = &helpers.Tree{Val: 3}
	root.Left.Left = &helpers.Tree{Val: 4}
	root.Left.Right = &helpers.Tree{Val: 5}
	root.Right.Left = &helpers.Tree{Val: 6}
	root.Right.Right = &helpers.Tree{Val: 7}

	var result [][]int
	var queue helpers.Queue
	queue.Enqueue(root)

	for len(queue.Q) > 0 {
		currentSize := len(queue.Q)
		currentLevel := []int{}
		for i := 0; i < currentSize; i++ {
			node := queue.Denqueue().(*helpers.Tree)
			val := node.Val
			currentLevel = append(currentLevel, val)

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
		//copy(result[1:], result)
		//result[0] = currentLevel
		result = append([][]int{currentLevel}, result...)
	}
	fmt.Println(result)
}
