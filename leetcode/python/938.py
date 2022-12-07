# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    count = 0

    def rangeSumBST(self, root: Optional[TreeNode], low: int, high: int) -> int:
        self.TraverseAndCount(root, low, high)

        return self.count

    def TraverseAndCount(self, node: TreeNode, low: int, high: int):
        if node:
            if node.val >= low and node.val <= high:
                self.count += node.val

            self.TraverseAndCount(node.left, low, high)
            self.TraverseAndCount(node.right, low, high)
