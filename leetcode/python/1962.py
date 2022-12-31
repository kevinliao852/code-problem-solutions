class Solution:
    def minStoneSum(self, piles: List[int], k: int) -> int:
        # Build a heap
        piles = [-p for p in piles]
        heapq.heapify(piles)

        for i in range(k):
            # Using heap to get the the highest piles
            n = heapq.heappop(piles)
            n //= 2
            heapq.heappush(piles, n)

        return -1 * sum(piles)
