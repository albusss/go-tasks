package fastAndSlow

import (
	"fmt"
	"github.com/albus/EduG"
	"github.com/albus/EduG/helpers"
)

func StartOfCycle() {
	list := main.GetListWithCycleAt4()
	cycleLength := 0
	slow := list
	fast := list

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next

		if fast.val == slow.val {
			cycleLength = findLength(slow)
			break
		}
	}

	fmt.Println(findStart(list, cycleLength).val)
}

func findLength(node *helpers.Node) int {
	length := 1
	next := node.next
	for next.val != node.val {
		length++
		next = next.next
	}

	return length
}

func findStart(node *helpers.Node, length int) *helpers.Node {
	slow := node
	fast := node
	for i := 0; i < length; i++ {
		fast = fast.next
	}

	for {
		slow = slow.next
		fast = fast.next

		if slow.val == fast.val {
			return slow
		}
	}
}
