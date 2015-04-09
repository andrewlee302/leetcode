# 5 Longest Palindromic Substring Medium
class Solution:
    # @return a string
    def longestPalindrome(self, s):
        if not s:
            return ""
        newS = ""
        # in order to unit the boundary checking
        # not only insert "#" into the gaps, but also insert to the start and the end
        for i in range(0, len(s)):
            newS += "#"
            newS += s[i]
        newS += "#"
        P = [1]
        maxBound = 0
        maxIndex = 0
        for i in range(1, len(newS)):
            tmp = 0
            if i < maxBound:
                tmp = min(P[2*maxIndex - i], maxBound - i + 1)
            else:
                tmp = 1
            while i+tmp<len(newS) and i-tmp>=0 and newS[i+tmp] == newS[i-tmp]:
                tmp+=1
            P.append(tmp)
            if i + tmp - 1 > maxBound:
                maxBound = i + tmp - 1
                maxIndex = i
        longestLenInNewS = 0
        longestMidIndexInNewS = 0
        for i in range(0, len(P)):
            if P[i] > longestLenInNewS:
                longestLenInNewS = P[i]
                longestMidIndexInNewS = i
        leftInS = (longestMidIndexInNewS-longestLenInNewS+2-1)/2
        rightInS = (longestMidIndexInNewS+longestLenInNewS-2-1)/2
        return s[leftInS:rightInS+1]

if "__main__" == __name__:
    s = Solution()
    sentence = "a"
    #sentence = "dabazzyyzzccd"
    print s.longestPalindrome(sentence)
