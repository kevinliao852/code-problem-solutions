class Solution:
    def countGoodSubstrings(self, s: str) -> int:
        s_len = len(s)
        count = 0
        for i in range(0, s_len - 2):
            if s[i] != s[i + 1] and s[i] != s[i + 2] and s[i + 1] != s[i + 2]:
                count += 1

        return count
