class Solution:
    def numberOfArithmeticSlices(self, nums: List[int]) -> int:
        d = [defaultdict(int) for _ in range(len(nums))]
        count = 0

        for i in range(1, len(nums)):
            for j in range(i):
                diff = nums[i] - nums[j]
                count += d[j][diff]

                d[i][diff] += d[j][diff] + 1

        return count
