# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def averageOfSubtree(self, root: Optional[TreeNode]) -> int:
        res, avg, cnt = self.find(root)
        return res

    def find(self, node: Optional[TreeNode]) -> Tuple[int, int, int]:
        if not node:
            return 0, 0, 0

        l_res, l_sum, l_cnt = self.find(node.left)
        r_res, r_sum, r_cnt = self.find(node.right)

        cnt = l_cnt + r_cnt + 1
        n_sum = l_sum + r_sum + node.val

        res = l_res + r_res
        if n_sum // cnt == node.val:
            res += 1

        return res, n_sum, cnt
