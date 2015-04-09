# 20 Valid Parentheses Easy
class Solution:
    # @return a boolean
    def isValid(self, s):
        match = {'(':')', '{':'}', '[':']'}
        stack = []

        for i in range(0, len(s)):
            if not stack:
                stack.append(s[i])
                continue
            top = stack[len(stack)-1]
            if match.has_key(top) and match[top] == s[i]:
                stack.pop()
            else:
                stack.append(s[i])
        if not stack:
            return True
        else:
            return False

if "__main__" == __name__:
    s = Solution()
    mark = "(({}[()]))"
    #mark = "(({}[)(]))"
    print s.isValid(mark)
