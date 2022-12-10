# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    s = 0

    def maxProduct(self, root: Optional[TreeNode]) -> int:
        sum = self.traverse_and_sum(root)
        res = self.traverse_and_find_max_product(root, sum)

        return self.s % (10**9 + 7)

    def traverse_and_sum(self, node: Optional[TreeNode]) -> int:
        if not node:
            return 0

        s = self.traverse_and_sum(node.left) + self.traverse_and_sum(node.right)

        return node.val + s

    def traverse_and_find_max_product(self, node: Optional[TreeNode], sum: int) -> int:
        if not node:
            return 0

        m = (
            self.traverse_and_find_max_product(node.left, sum)
            + self.traverse_and_find_max_product(node.right, sum)
            + node.val
        )

        self.s = max(m * (sum - m), self.s)

        return m
