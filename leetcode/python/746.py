class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:

        for i, _ in enumerate(cost):
            if i < 2:
                continue

            if cost[i - 1] > cost[i - 2]:
                cost[i] += cost[i - 2]
            else:
                cost[i] += cost[i - 1]

        return cost[-1] if cost[-1] < cost[-2] else cost[-2]
