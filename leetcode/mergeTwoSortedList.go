package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func Merge_two_sorted_lists(list1 []int, list2 []int) []int {
	retArr := make([]int, 0)
	p1 := 0
	p2 := 0

	for p1 < len(list1) && p2 < len(list2) {
		if list1[p1] > list2[p2] {
			retArr = append(retArr, list2[p2])
			p2++
		} else {
			retArr = append(retArr, list1[p1])
			p1++
		}
	}

	for p1 < len(list1) {
		retArr = append(retArr, list1[p1])
		p1++
	}

	for p2 < len(list2) {
		retArr = append(retArr, list2[p2])
		p2++
	}

	return retArr
}

func Merge_two_sorted_lists_linked_list(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	head := ListNode{Val: 0, Next: nil}
	hp := &head

	for list1.Next != nil && list2.Next != nil {
		if list1.Val > list2.Val {
			hp.Next = &ListNode{list2.Val, nil}
			hp = hp.Next
			list2 = list2.Next
		} else {
			hp.Next = &ListNode{list1.Val, nil}
			hp = hp.Next
			list1 = list1.Next
		}
	}

	for list1.Next != nil {
		hp.Next = &ListNode{list1.Val, nil}
		hp = hp.Next
		list1 = list1.Next
	}

	for list2.Next != nil {
		hp.Next = &ListNode{list2.Val, nil}
		hp = hp.Next
		list2 = list2.Next
	}

	hp.Next = nil
	return head.Next
}
