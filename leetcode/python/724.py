class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        s = sum(nums)
        p = 0

        for i, v in enumerate(nums):

            if s - p * 2 == v:
                return i
            p += v

        return -1
