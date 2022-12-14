class Solution:
    def rob(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]

        m = nums[0]
        n = max(nums[1], nums[0])

        for i in range(2, len(nums)):
            cur = nums[i]
            if i % 2 == 0:
                m += cur
                m = max(m, n)
            else:
                n += cur
                n = max(m, n)

        return max(m, n)
