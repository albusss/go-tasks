package reverseList

import (
	"github.com/albus/EduG/helpers"
)

func InplaceReverse() {
	head := helpers.GetList()

	var current *helpers.Node
	var prev *helpers.Node
	current = head

	for current != nil {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
	}
}
