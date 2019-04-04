# 62 Unique Paths Medium
class Solution:
    # @return an integer
    def uniquePaths(self, m, n):
        return self.factorial(m+n-2)/(self.factorial(m-1)*self.factorial(n-1))

    def factorial(self, n):
        product = 1
        for i in range(1, n+1):
            product *= i
        return product
