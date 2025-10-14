package leetcode

/*
1 -> 1 -> 2 -> 3 -> 3 -> 4
|		  |
c	->	  c

start with c

	for c.next != nil {
		if (c.next.val == c.val) {
			c.next = c.next.next
		}
		c = c.next

}
*/
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	curr := head
	var prevNode *ListNode

	for curr != nil {
		// check if the current node is still equal to prevNode val
		if prevNode != nil && curr.Val == prevNode.Val {
			prevNode.Next = curr.Next
			curr = curr.Next
			continue
		}

		if curr.Next == nil {
			break
		}
		if curr.Val == curr.Next.Val {
			prevNode = curr
			// skip the node
			curr.Next = curr.Next.Next
		}
		curr = curr.Next
	}

	return head
}
