class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        # only go down and right
        s = [[0 for j in range(m + n)] for i in range(m + n)]

        s[0][0] = 1

        for i in range(m + n - 1):
            for j in range(i + 1):
                if i == 0 and j == 0:
                    continue
                s[i][j] = s[i - 1][j - 1] + s[i - 1][j]

        return s[m + n - 2][n - 1]
