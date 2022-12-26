class Solution:
    def canJump(self, nums: List[int]) -> bool:
        max_jump = 0
        res = False
        for i, n in enumerate(nums):
            if i == len(nums) - 1:
                res = True
                break
            if n == 0 and max_jump == 0:
                res = False
                break

            max_jump = max(n - 1, max_jump - 1)

        return res
