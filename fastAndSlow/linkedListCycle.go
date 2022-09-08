package fastAndSlow

import (
	"github.com/albus/EduG/helpers"
)

func LinkedListCycle() {
}

func hasCycle(node *helpers.Node) bool {
	slow := node
	fast := node

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow.val == fast.val {
			return true
		}
	}

	return false
}
