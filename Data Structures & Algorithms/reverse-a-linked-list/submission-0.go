/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    if head == nil {
		return head
	}
	h,_ := recurse(head)
	head.Next = nil
	return h
}

func recurse(head *ListNode)(*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	}
	fhead, prev := recurse(head.Next)
	prev.Next = head
	return fhead, head
}
