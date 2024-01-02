class Solution:
    def findMatrix(self, nums: List[int]) -> List[List[int]]:
        hashmap = {}
        r = []

        for n in nums:
            if n not in hashmap:
                hashmap[n] = 1
            else:
                hashmap[n] += 1

            if len(r) < hashmap[n]:
                r.append([n])
            else:
                r[hashmap[n] - 1].append(n)
        return r
