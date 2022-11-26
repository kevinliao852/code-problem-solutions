class Solution:
    def getHint(self, secret: str, guess: str) -> str:
        d = {}
        a = 0
        b = 0

        for s in secret:
            if s not in d:
                d[s] = 1
            else:
                d[s] += 1

        for i, s in enumerate(guess):
            if s in d and s == secret[i]:
                d[s] -= 1
                a += 1

        for i, s in enumerate(guess):
            if s in d and d[s] > 0 and s != secret[i]:
                d[s] -= 1
                b += 1

        return self.output_str(a, b)

    def output_str(self, a: int, b: int) -> str:
        return f"{a}A{b}B"
