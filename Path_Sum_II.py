# 113 Path Sum II Medium
# Definition for a  binary tree node
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    # @param root, a tree node
    # @param sum, an integer
    # @return a list of lists of integers
    def pathSum(self, root, sum):
        if not root:
            return []
        return self.DFS(root, sum-root.val)


    def DFS(self, parent, num):
        tail = []
        if not parent.left and not parent.right:
            if 0 == num:
                tail.append([parent.val])
            return tail
        if parent.left:
            leftResult = self.DFS(parent.left, num - parent.left.val)
            if len(leftResult) != 0:
                for i in range(0, len(leftResult)):
                    leftResult[i].insert(0, parent.val)
                    tail.append(leftResult[i])
        if parent.right:
            rightResult = self.DFS(parent.right, num - parent.right.val)
            if len(rightResult) != 0:
                for i in range(0, len(rightResult)):
                    rightResult[i].insert(0, parent.val)
                    tail.append(rightResult[i])
        return tail
