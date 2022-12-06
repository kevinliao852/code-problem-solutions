class Solution:
    def oddEvenList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head:
            return None

        temp = head
        l = []
        while head:
            l.append(head.val)
            head = head.next

        head = temp
        j = 0
        for _ in range(len(l)):
            temp.val = l[j]
            temp = temp.next
            j += 2

            if j > len(l) - 1:
                j = 1
        return head
