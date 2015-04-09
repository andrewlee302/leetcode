class Solution:
    # @return a string
    def minWindow(self, S, T):
        charMap = {}
        for t in T:
            if t not in charMap.keys():
                charMap[t] = 1
            else:
                charMap[t] += 1
        start = 0
        pointer = start

        flag = False
        minLength = len(S)
        minStr = ""

        for i in range(0, len(S)):
            s = S[i]
            flag = False
            if s not in T:
                continue
            start = i
            if self.isValid(charMap):
                flag = True
            else:
                while pointer < len(S):
                    if S[pointer] in T:
                        charMap[S[pointer]]-=1
                        if self.isValid(charMap):
                            flag = True
                            pointer+=1
                            break
                    pointer+=1
            if flag:
                length = pointer - start
                if length <= minLength:
                    minLength = length
                    minStr = S[start:pointer]
                charMap[S[start]] += 1
            else:
                break
        return minStr

    def isValid(self, charMap):
        for v in charMap.values():
            if v > 0:
                return False
        return True

if "__main__" == __name__:
    s = Solution()
    T = "aa"
    S = "aa"
    print s.minWindow(S,T)
