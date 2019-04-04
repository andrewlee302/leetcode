# 4 Median of Two Sorted Arrays Hard
# Should be optimized by the divide-conquer approach
class Solution:
    # @return a float
    def findMedianSortedArrays(self, A, B):
        array = []
        i = j = 0
        while True:
            if i == len(A):
                while j < len(B):
                    array.append(B[j])
                    j+=1
                break
            elif j == len(B):
                while i < len(A):
                    array.append(A[i])
                    i+=1
                break
            else:
                if A[i] < B[j]:
                    array.append(A[i])
                    i+=1
                else:
                    array.append(B[j])
                    j+=1
        num = len(array)
        if num % 2 == 0:
            return (float(array[num/2]) + array[num/2-1])/2
        else:
            return array[num/2]
