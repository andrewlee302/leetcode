# 3 Longest Substring Without Repeating Characters Medium
class Solution:
    # @return an integer
    def lengthOfLongestSubstring(self, s):
        charMap = {}
        start = 0
        pointer = start
        maxLength = 0
        for i in range(0, len(s)):
            start = i
            while pointer < len(s) and s[pointer] not in charMap.keys():
                charMap[s[pointer]] = 1
                pointer+=1
            length = pointer - start
            if length > maxLength:
                maxLength = length
            charMap.pop(s[start])
            if pointer == len(s):
                break
        return maxLength
