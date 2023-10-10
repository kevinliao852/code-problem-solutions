class Solution:
    def targetIndices(self, nums: List[int], target: int) -> List[int]:
        nums.sort()
        lo, hi = 0, len(nums)

        while lo < hi:
            m = lo + (hi - lo) // 2
            num = nums[m]

            if target == num:
                i = m - 1
                j = m + 1
                while i >= 0 and target == nums[i]:
                    i -= 1
                while j < len(nums) and target == nums[j]:
                    j += 1

                return [n for n in range(i + 1, j)]

            elif target > num:
                lo = m + 1
            else:
                hi = m

        return []
