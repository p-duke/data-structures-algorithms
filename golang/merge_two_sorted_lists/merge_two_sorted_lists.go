/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	result := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
            result.Next = list1
	        list1 = list1.Next
		} else {
            result.Next = list2
			list2 = list2.Next
		}
        result = result.Next
	}

	if list1 != nil {
		result.Next = list1
	} else {
		result.Next = list2
	}

	return head.Next
}
