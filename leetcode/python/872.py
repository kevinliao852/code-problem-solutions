# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def leafSimilar(self, root1: Optional[TreeNode], root2: Optional[TreeNode]) -> bool:
        records = []
        self.traverse_and_record(root1, records)
        return self.traverse_and_compare(root2, records) and len(records) == 0

    def traverse_and_record(self, node: Optional[TreeNode], records: list):
        if not node:
            return

        if node.left is None and node.right is None:
            records.append(node.val)

        self.traverse_and_record(node.left, records)
        self.traverse_and_record(node.right, records)

    def traverse_and_compare(self, node: Optional[TreeNode], records: list):
        if not node:
            return True

        if node.left is None and node.right is None:
            if len(records) == 0:
                return False
            if records[0] != node.val:
                return False

            records.pop(0)

        return self.traverse_and_compare(
            node.left, records
        ) and self.traverse_and_compare(node.right, records)
