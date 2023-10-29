class Solution:
    def isHappy(self, n: int) -> bool:
        hashmap = {}

        str_n = str(n)
        s = 0

        while s != 1:
            for c in str_n:
                s += int(c) * int(c)

            print(s)
            if s == 1:
                return True

            if s in hashmap:
                return False

            hashmap[s] = True

            str_n = str(s)
            s = 0
