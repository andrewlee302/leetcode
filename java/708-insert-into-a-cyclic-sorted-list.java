/*
// Definition for a Node.
class Node {
    public int val;
    public Node next;

    public Node() {}

    public Node(int _val,Node _next) {
        val = _val;
        next = _next;
    }
};
*/
class Solution {
    public Node insert(Node head, int insertVal) {
        Node node = new Node(insertVal);
        node.next = node;
        if (head == null)
            return node;
        Node p = head, q = head.next;
        while (q != head) { // Can't be p != head, as it never goes into the loop.
            if (p.val <= q.val && insertVal >= p.val && insertVal <= p.next.val || p.val > p.next.val && (insertVal >= p.val || insertVal <= p.next.val)) {
                break;
            }
            p = q;
            q = q.next;
        }
        p.next = node;
        node.next = q;
        return head;
    }
}
