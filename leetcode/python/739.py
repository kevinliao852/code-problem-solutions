class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        stack: (int, int) = []
        res: List[int] = [0] * len(temperatures)

        for i, t in enumerate(temperatures):
            while len(stack) > 0 and stack[-1][1] < t:
                idx, _ = stack.pop()
                res[idx] = i - idx
            stack.append((i, t))
        return res
