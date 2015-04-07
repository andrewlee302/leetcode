# 112 Path Sum Easy
# Definition for a  binary tree node
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    # @param root, a tree node
    # @param sum, an integer
    # @return a boolean
    def hasPathSum(self, root, sum):
        if not root:
            return False
        return self.DFS(root, sum-root.val)

    def DFS(self, parent, num):
        if not parent.left and not parent.right:
            if 0 == num:
                return True
            else:
                return False
        if parent.left and self.DFS(parent.left, num - parent.left.val):
            return True
        if parent.right and self.DFS(parent.right, num - parent.right.val):
            return True
        return False
