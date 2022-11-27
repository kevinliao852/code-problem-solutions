class Solution:
    def countPairs(self, nums: List[int], k: int) -> int:
        count = 0

        for i in range(1, len(nums)):
            for j in range(i):
                if nums[i] == nums[j] and i * j % k == 0:
                    count += 1

        return count
