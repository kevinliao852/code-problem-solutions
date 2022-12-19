class Solution:
    def validPath(
        self, n: int, edges: List[List[int]], source: int, destination: int
    ) -> bool:
        vertices = [[] for i in range(n)]

        for edge in edges:
            v1, v2 = edge[0], edge[1]
            vertices[v2].append(v1)
            vertices[v1].append(v2)

        res = False

        hashmap = {}

        stack = [source]

        while len(stack):
            v = stack.pop()
            hashmap[v] = True

            if v == destination:
                res = True
                break

            for n in vertices[v]:
                if n not in hashmap:
                    stack.append(n)

        return res
