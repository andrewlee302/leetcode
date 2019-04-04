# 42 Trapping Rain Water Hard
class Solution:
    # @param A, a list of integers
    # @return an integer
    def trap(self, A):
        if not A or len(A) == 1:
            return 0
        storage = []
        totalArea = 0
        start = len(A)
        for i in range(0, len(A) - 1):
            if A[i] > A[i+1]:
                storage.append((A[i], i))
                start = i+1
                break
        for i in range(start, len(A)):
            t = (A[i], i)
            leftTuple = None
            while storage and A[i] >= storage[len(storage)-1][0]:
                leftTuple = storage.pop()
            if len(storage) == 0:
                leftIndex = leftTuple[1]
                height = leftTuple[0]
                area = height * (i - leftIndex - 1)
                for j in range(leftIndex+1, i):
                    area -= A[j]
                totalArea += area
            storage.append(t)
        if len(storage) > 1:
            rightTuple = storage.pop()
            while storage:
                rightIndex = rightTuple[1]
                height = rightTuple[0]
                rightTuple = storage.pop()
                leftIndex = rightTuple[1]
                area = height * (rightIndex - leftIndex - 1)
                for j in range(leftIndex+1, rightIndex):
                    area -= A[j]
                totalArea += area
        return totalArea

if __name__ == "__main__":
    s = Solution()
    # a = [3, 0, 1, 3, 1, 2, 4, 4]
    a = [0,1,0,2,1,0,1,3,2,1,2,1]
    print s.trap(a)
