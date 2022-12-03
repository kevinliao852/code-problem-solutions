class Solution:
    def frequencySort(self, s: str) -> str:
        d = Counter(s)

        l = [(number, letter) for letter, number in d.most_common()]

        l.sort(reverse=True)

        res = []
        for number, letter in l:
            for _ in range(number):
                res.append(letter)

        return "".join(res)
