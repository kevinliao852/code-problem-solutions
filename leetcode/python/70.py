class Solution:
    def climbStairs(self, n: int) -> int:
        if n == 1 or n == 2:
            return n

        l = [0 for i in range(n + 1)]
        l[1] = 1
        l[2] = 2

        for i in range(3, n + 1):
            l[i] = l[i - 1] + l[i - 2]

        return l[-1]
