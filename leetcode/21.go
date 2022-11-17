/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var list *ListNode = &ListNode{}
	var temp1 = list1
	var temp2 = list2
	var res = list

	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	for temp1 != nil && temp2 != nil {

		if temp1.Val <= temp2.Val {
			list.Val = temp1.Val
			temp1 = temp1.Next
		} else {
			list.Val = temp2.Val
			temp2 = temp2.Next
		}

		if temp1 == nil || temp2 == nil {
			break
		}

		list.Next = &ListNode{}
		list = list.Next
	}

	if temp1 != nil {
		list.Next = temp1
	} else {
		list.Next = temp2
	}

	return res
}
