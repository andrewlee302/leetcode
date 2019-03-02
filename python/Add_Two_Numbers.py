# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    # @return a ListNode
    def addTwoNumbers(self, l1, l2):
        cur1 = l1
        cur2 = l2
        tmp = cur1.val + cur2.val
        result = None
        advance = 0
        if tmp > 9:
            result = ListNode(tmp - 10)
            advance = 1
        else:
            result = ListNode(tmp)
        curRes = result
        while cur1.next and cur2.next:
            cur1 = cur1.next
            cur2 = cur2.next
            tmp = cur1.val + cur2.val + advance
            tmpNode = None
            if tmp > 9:
                tmpNode = ListNode(tmp - 10)
                advance = 1
            else:
                tmpNode = ListNode(tmp)
                advance = 0
            curRes.next = tmpNode
            curRes = tmpNode
        ano = None
        if not cur1.next:
            ano = cur2
        else:
            ano = cur1
        while ano.next:
                ano = ano.next
                tmp = ano.val + advance
                if tmp > 9:
                    tmpNode = ListNode(tmp - 10)
                    advance = 1
                else:
                    tmpNode = ListNode(tmp)
                    advance = 0
                curRes.next = tmpNode
                curRes = tmpNode
        if advance == 1:
            curRes.next = ListNode(1)
        return result
