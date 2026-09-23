/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil{
		return list1
	}

	smaller := list1
	larger := list2
	if list1.Val > list2.Val {
		smaller = list2
		larger = list1
	}
	smaller.Next = mergeTwoLists(smaller.Next, larger)
	return smaller
}
