/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	temp := head
	prev := head

	for temp != nil {
		next := temp.Next
		if temp != prev {
			temp.Next = prev
		} else {
			temp.Next = nil
		}

		prev = temp
		temp = next
	}

	return prev
}
