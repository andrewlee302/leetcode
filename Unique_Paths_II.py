# 63 Unique Paths II Medium
class Solution:
    # @param obstacleGrid, a list of lists of integers
    # @return an integer
    def uniquePathsWithObstacles(self, obstacleGrid):
        m = len(obstacleGrid)
        n = len(obstacleGrid[0])
        D = [ [0 for j in range(0,n)] for i in range(0,m)]
        if 0 == obstacleGrid[0][0]:
            D[0][0] = 1
        else:
            D[0][0] = 0
        for i in range(1, m):
            if 0 == obstacleGrid[i][0]:
                D[i][0] = D[i-1][0]
            else:
                D[i][0] = 0
        for i in range(1, n):
            if 0 == obstacleGrid[0][i]:
                D[0][i] = D[0][i-1]
            else:
                D[0][i] = 0
        for i in range(1, m):
            for j in range(1, n):
                if 0 == obstacleGrid[i][j]:
                    D[i][j] = D[i][j-1] + D[i-1][j]
                else:
                    D[i][j] = 0
        return D[m-1][n-1]
