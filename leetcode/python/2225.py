class Solution:
    def findWinners(self, matches: List[List[int]]) -> List[List[int]]:
        win = []
        lose_only_once = []

        d = {}

        for w, l in matches:
            if w not in d:
                d[w] = 1

            if l not in d:
                d[l] = 0
            elif d[l] == 1:
                d[l] = 0
            elif d[l] == 0:
                d[l] = -1

        for k, v in d.items():
            if v == 0:
                lose_only_once.append(k)
            if v == 1:
                win.append(k)

        lose_only_once.sort()
        win.sort()
        return [win, lose_only_once]
