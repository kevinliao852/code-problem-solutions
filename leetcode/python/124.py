# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
MIN_NUM = -1001


class Solution:
    max_sum = MIN_NUM

    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        def find(root):
            if not root:
                return MIN_NUM

            left_sum = find(root.left)
            right_sum = find(root.right)

            s: int

            if left_sum == MIN_NUM and right_sum == MIN_NUM:
                s = root.val
            elif left_sum == MIN_NUM:
                s = max(right_sum + root.val, root.val)
            elif right_sum == MIN_NUM:
                s = max(left_sum + root.val, root.val)
            else:
                s = max(root.val, right_sum + root.val, left_sum + root.val)

            self.max_sum = max(self.max_sum, s, left_sum + right_sum + root.val)
            return s

        find(root)
        return self.max_sum
