class Solution:
    def mostWordsFound(self, sentences: List[str]) -> int:
        max_num = 0

        for s in sentences:
            cnt = 0
            for w in s:
                if w == ' ':
                    cnt += 1
            max_num = max(cnt, max_num)

        return max_num + 1
