class Solution:
    def uniqueOccurrences(self, arr: List[int]) -> bool:
        d = defaultdict(int)
        d2 = defaultdict(bool)

        for v in arr:
            d[v] += 1

        for _, v in d.items():
            if d2[v] == True:
                return False

            d2[v] = True

        return True
