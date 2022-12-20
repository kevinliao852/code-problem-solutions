class Solution:
    def canVisitAllRooms(self, rooms: List[List[int]]) -> bool:
        # the data at each index in the rooms is dintict key

        res = False
        visited = {}
        stack = [0]

        while len(stack) > 0:
            idx = stack.pop(-1)

            visited[idx] = True

            if len(visited) == len(rooms):
                res = True
                break

            for i in rooms[idx]:
                if i not in visited:
                    stack.append(i)
        return res
