class Solution:
    def tribonacci(self, n: int) -> int:
        if n == 0:
            return 0
        if n == 1:
            return 1
        if n == 2:
            return 1

        a = 1
        b = 1
        c = 0

        for i in range(2, n):
            a, b, c = a + b + c, a, b

        return a
