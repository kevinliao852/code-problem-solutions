class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        d = {}
        # only one solution
        for idx, v in enumerate(nums):
            if v not in d:
                d[target - v] = idx
            else:
                return [d[v], idx]
