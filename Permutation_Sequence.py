# 60 Permutation Sequence Medium
class Solution:
    # @return a string
    def getPermutation(self, n, k):
        if(n==1):
            return "1"
        divisor =[]
        result = ""
        product = 1
        for i in range(1, n):
            product *= i
            divisor.append(product)
        for i in range(0, n):
            # when i = n-1, k = 1, order = 1
            order = (k - 1) / divisor[n-2-i] + 1
            v = 1
            while 0 < order:
                flag = True
                if result.find(str(v)) != -1:
                    flag = False
                if flag:
                    order -= 1
                v+=1
            result += str(v-1)
            k = (k - 1) % divisor[n-2-i] + 1
        return result

if "__main__" == __name__:
    s = Solution()
    n = 6
    k = 1
    print s.getPermutation(n,k)
