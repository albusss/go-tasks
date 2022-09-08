package reverseList

import (
	"github.com/albus/EduG/helpers"
)

func ReverseSubList() {
	head := helpers.GetList()
	p := 2
	q := 4
	var prev *helpers.Node
	var lastNodeFirstPart *helpers.Node
	var lastNodeSub *helpers.Node

	current := head
	i := 1
	for current != nil && i < p {
		prev = current
		current = current.Next
		i++
	}

	lastNodeFirstPart = prev
	lastNodeSub = current

	i = 0
	for current != nil && i < q-p+1 {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
		i++
	}

	lastNodeFirstPart.Next = prev
	lastNodeSub.Next = current
}
