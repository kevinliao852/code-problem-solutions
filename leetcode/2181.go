/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
	var res ListNode
	var sum int

	head = head.Next
	temp := &res

	for head != nil {
		if head.Val == 0 && sum != 0 {
			if temp.Val != 0 {
				temp.Next = &ListNode{}
				temp = temp.Next
			}
			temp.Val = sum
			sum = 0
		}
		sum += head.Val
		head = head.Next
	}

	return &res
} 
