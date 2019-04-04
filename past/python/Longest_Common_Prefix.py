# 14 Longest Common Prefix Easy
class Solution:
    # @return a string
    def longestCommonPrefix(self, strs):
        if not strs:
            return ""
        result=""
        i = 0
        flag = True
        while True:
            tmp = ''
            for s in strs:
                if len(s) < i + 1 or tmp and tmp != s[i]:
                    flag = False
                    break
                else:
                    if not tmp:
                        tmp = s[i]
            if not flag:
                break
            flag = True
            result+=tmp
            i+=1
        return result

if "__main__" == __name__:
    s = Solution()
    sen = ["abc", "abce,", "abc1"]
    print s.longestCommonPrefix(sen)
