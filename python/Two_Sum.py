# 1 Two Sum Medium
class Solution:
    # @return a tuple, (index1, index2)
    def twoSum(self, num, target):
        hash = {}
        for i in range(0, len(num)):
            if not (target - num[i]) in hash:
                hash[num[i]] = i
            else:
                j = hash[target - num[i]]
                return (min(i+1,j+1), max(i+1,j+1))
