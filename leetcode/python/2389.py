class Solution:
    def answerQueries(self, nums: List[int], queries: List[int]) -> List[int]:
        nums.sort()

        res = [0] * len(queries)

        for i, q in enumerate(queries):
            sum = 0
            cnt = 0
            for n in nums:
                sum += n
                cnt += 1

                if sum <= q:
                    res[i] = cnt
                else:
                    break

        return res
