class Solution:
    def minimumAverageDifference(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return 0

        left = 0
        s = 0
        min_res = 100001
        idx = 0

        for n in nums:
            s += n

        for i in range(0, len(nums)):
            n = nums[i]
            left += n
            right = s - left
            j = len(nums) - i - 1

            if j == 0:
                diff = abs(left // (i + 1))
            else:
                diff = abs(left // (i + 1) - right // j)

            if min_res > diff:
                idx = i
                min_res = min(min_res, diff)

        return idx
