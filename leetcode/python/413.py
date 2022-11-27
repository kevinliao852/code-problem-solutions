class Solution:
    def numberOfArithmeticSlices(self, nums: List[int]) -> int:

        offset = 0
        offset_count = 0
        count = 0

        for i in range(len(nums) - 1):
            if offset == nums[i + 1] - nums[i]:
                offset_count += 1
                if offset_count >= 2:
                    count += offset_count - 1
            else:
                offset_count = 1
            offset = nums[i + 1] - nums[i]

        return count
