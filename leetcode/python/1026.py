# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def maxAncestorDiff(self, root: Optional[TreeNode]) -> int:
        return self.traverse_and_find_max(root, [], root.val, root.val)

    def traverse_and_find_max(
        self, node: Optional[TreeNode], l: list, m: int = 0, n: int = 0
    ) -> int:
        if not node:
            return abs(m - n)

        m = min(m, node.val)
        n = max(n, node.val)

        return max(
            self.traverse_and_find_max(node.left, l, m, n),
            self.traverse_and_find_max(node.right, l, m, n),
        )
