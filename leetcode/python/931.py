class Solution:
    def minFallingPathSum(self, matrix: List[List[int]]) -> int:
        m = matrix
        n = len(matrix)

        dp = [[float('inf')] * n for i in range(n)]

        for i in range(n):
            dp[0][i] = m[0][i]

        for i in range(n - 1):
            for j in range(n):
                if j < n - 1:
                    dp[i + 1][j + 1] = min(dp[i][j] + m[i + 1][j + 1], dp[i + 1][j + 1])
                if j != 0:
                    dp[i + 1][j - 1] = min(dp[i][j] + m[i + 1][j - 1], dp[i + 1][j - 1])

                dp[i + 1][j] = min(dp[i][j] + m[i + 1][j], dp[i + 1][j])

        return min(dp[n - 1])
