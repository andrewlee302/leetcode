# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

    def __str__(self):
        cursor = self
        if not cursor:
            return "None"
        strPrint = "%d" %cursor.val
        print cursor.val
        cursor = cursor.next
        while cursor:
            strPrint += ("->%d" %cursor.val)
            print cursor.val
            cursor = cursor.next
        return strPrint

class Solution:
    # @param head, a ListNode
    # @param k, an integer
    # @return a ListNode
    def reverseKGroup(self, head, k):
        newHead = head
        former_tail = None
        group_head = head
        tmp_node = None
        i_node = head
        j_node = None
        if i_node:
            j_node = i_node.next
        while i_node:
            tmp_node = None
            for i in range(k):
                if not i_node:
                    i_node = tmp_node
                    tmp_node = None
                    j_node = i_node.next
                    while i_node:
                        i_node.next = tmp_node
                        tmp_node = i_node
                        i_node = j_node
                        if not i_node:
                            j_node = None
                        else:
                            j_node = i_node.next
                    break
                else:
                    i_node.next = tmp_node
                    tmp_node = i_node
                    i_node = j_node
                    if not i_node:
                        j_node = None
                    else:
                        j_node = i_node.next
            if newHead == head and k != 1:
                newHead = tmp_node
            if former_tail:
                former_tail.next = tmp_node
            former_tail = group_head
            if i_node:
                group_head = i_node
                j_node = i_node.next
        return newHead

if __name__ == "__main__":
    head = ListNode(1)
    cursor = head
    for i in range(2, 10):
        tmp = ListNode(i)
        cursor.next = tmp
        cursor = tmp
    print head
    solu = Solution()
    newList = solu.reverseKGroup(head, 1)
    print "newList.val=%d" %newList.val
    print newList


