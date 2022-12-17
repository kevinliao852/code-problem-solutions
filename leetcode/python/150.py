class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []

        operator = {
            "+": True,
            "-": True,
            "*": True,
            "/": True,
        }

        if len(tokens) == 1:
            return int(tokens[0])

        for token in tokens:
            if token in operator:
                res = stack.pop()
                n = stack.pop()

                if token == "+":
                    res = n + res
                if token == "-":
                    res = n - res
                if token == "*":
                    res = n * res
                if token == "/":
                    res = n / res

                res = int(res)
                stack.append(res)

            else:
                stack.append(int(token))

        return stack[-1]
