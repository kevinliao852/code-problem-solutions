class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        l = [[0] * len(text2) for i in range(len(text1))]

        for i in range(len(text1)):
            for j in range(len(text2)):

                # case1
                if i > 0:
                    l[i][j] = l[i - 1][j]
                if j > 0:
                    l[i][j] = max(l[i][j - 1], l[i][j])

                # case2
                s = 1
                if text1[i] == text2[j] and i > 0 and j > 0:
                    s = max(l[i - 1][j - 1] + 1, l[i][j])

                if text1[i] == text2[j]:
                    l[i][j] = max(s, l[i][j])

        # print(l)
        return l[-1][-1]
