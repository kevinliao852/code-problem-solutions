/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	t := head
	r := head
	cnt := 0

	if t == nil || t.Next == nil || t.Next.Next == nil {
		return nil
	}

	for t != nil && r != nil && r.Next != nil {
		cnt++
		t = t.Next
		r = r.Next.Next

		if t == r {
			if head == t {
				return t
			}
			t = head
			for t != r {
				t = t.Next
				r = r.Next
			}
			return t
		}
	}

	return nil
}
