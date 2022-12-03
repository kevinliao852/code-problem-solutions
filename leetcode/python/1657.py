class Solution:
    def closeStrings(self, word1: str, word2: str) -> bool:
        if len(word1) != len(word2):
            return False

        d = defaultdict(int)
        e = defaultdict(int)

        for s in word1:
            d[s] += 1
        for s in word2:
            e[s] += 1
            if s not in d:
                print(s)
                return False

        for s in word1:
            if s not in e:
                print(s)
                return False

        r1 = [v for _, v in d.items()]
        r2 = [v for _, v in e.items()]

        r1.sort()
        r2.sort()

        if len(r1) != len(r2):
            return False

        for i in range(len(r1)):
            if r1[i] != r2[i]:
                return False

        return True
