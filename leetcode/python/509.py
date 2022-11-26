class Solution:
    def fib(self, n: int) -> int:
        if n == 0 or n == 1:
            return n

        l = [0 for i in range(n + 1)]
        l[1] = 1

        for i in range(2, n + 1):
            l[i] = l[i - 1] + l[i - 2]

        return l[-1]
