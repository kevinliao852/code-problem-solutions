class Solution:
    def maximumBags(
        self, capacity: List[int], rocks: List[int], additionalRocks: int
    ) -> int:
        offsets = [0] * len(capacity)
        for i, c in enumerate(capacity):
            offsets[i] = c - rocks[i]

        offsets.sort()

        cnt = 0

        for o in offsets:
            if additionalRocks - o < 0:
                break
            additionalRocks -= o
            cnt += 1

        return cnt
