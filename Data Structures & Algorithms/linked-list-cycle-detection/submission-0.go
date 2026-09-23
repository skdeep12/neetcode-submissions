/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
    faster := head
	slower := head
	for true {
		slower = slower.Next
		if slower == nil {
			return false
		}
		faster = faster.Next
		if faster == nil {
			return false
		}
		faster = faster.Next
		if faster == nil {
			return false
		}
		if slower==faster{
			break
		}
	}
	return true
}
