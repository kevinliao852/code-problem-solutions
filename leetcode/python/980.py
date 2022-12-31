class Solution:
    def uniquePathsIII(self, grid: List[List[int]]) -> int:
        paths = 0
        # use dfs to find the path
        x = 0
        y = 0
        c = 0

        for i, l in enumerate(grid):
            for j, v in enumerate(l):
                if v == 1:
                    x = i
                    y = j
                if v != -1:
                    c += 1

        print(c)

        def dfs(grid, x, y, c, cnt=0):

            cnt += 1

            if x >= len(grid) or x < 0 or y >= len(grid[0]) or y < 0:
                return 0

            if grid[x][y] == 2 and cnt == c:
                return 1
            elif grid[x][y] == 2:
                return 0

            if grid[x][y] == -1:
                return 0

            prev = grid[x][y]
            grid[x][y] = -1

            s = (
                dfs(grid, x, y + 1, c, cnt)
                + dfs(grid, x + 1, y, c, cnt)
                + dfs(grid, x, y - 1, c, cnt)
                + dfs(grid, x - 1, y, c, cnt)
            )

            # crucial part
            grid[x][y] = prev

            return s

        paths = dfs(copy.deepcopy(grid), x, y, c)

        return paths
