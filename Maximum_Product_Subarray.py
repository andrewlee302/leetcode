# 152 Maximum Product Subarray Medium
class Solution:
    # @param A, a list of integers
    # @return an integer
    def maxProduct(self, A):
        # The first judge is important. If len(A) != 1, the posiLargest is valid(may be 0). Otherwise, posiLargest mightn't exist.
        if len(A) == 1:
            return A[0]
        posiLargest, negaLargest = self.maxProductInterval(A, 0, len(A))
        return posiLargest

    def maxProductInterval(self, A, start, length):
        if length == 1:
            if A[start] >= 0:
                return (A[start], 0)
            else:
                return (0, A[start])
        divideIndex = start + length/2
        leftPosiLargest, leftNegaLargest = self.maxProductInterval(A, start, length/2)
        rightPosiLargest, rightNegaLargest = self.maxProductInterval(A, divideIndex, length-length/2)
        leftPartNum = rightPartNum = 1
        leftPosiPart = leftNegaPart = rightPosiPart = rightNegaPart = 0
        flag = True # positive
        for i in range(0, length/2):
            index = start + length/2 - i - 1
            if A[index] == 0:
                break
            elif A[index] < 0:
                flag = not flag
            leftPartNum *= A[index]
            if flag:
                leftPosiPart = leftPartNum
            else:
                leftNegaPart = leftPartNum
        flag = True
        for i in range(0, length-length/2):
            index = divideIndex + i
            if A[index] == 0:
                break
            elif A[index] < 0:
                flag = not flag
            rightPartNum *= A[index]
            if flag:
                rightPosiPart = rightPartNum
            else:
                rightNegaPart = rightPartNum
        posiLargest = max(leftPosiLargest, rightPosiLargest, leftPosiPart*rightPosiPart, leftNegaPart*rightNegaPart)
        negaLargest = -max(-leftNegaLargest, -rightNegaLargest, -leftPosiPart*rightNegaPart, -leftPosiPart*rightNegaPart)
        return (posiLargest, negaLargest)

if "__main__" == __name__:
    s = Solution()
    A = [2,3,-2,4]
    print s.maxProduct(A)
