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
        if (head == null) {
            head = new Node(insertVal, null);
            head.next = head;
            return head;
        }
        Node curr = head.next, prev = head;
        while (curr != head) {
            // Note!!!
            if (curr.val < prev.val && (insertVal < curr.val || insertVal > prev.val) || insertVal >= prev.val && insertVal <= curr.val){
                Node node = new Node(insertVal, curr);
                prev.next = node;
                break;
            }
            prev = curr;
            curr = curr.next;
        }
        if (curr == head) {
            Node node = new Node(insertVal, curr);
            prev.next = node;
        }
        return head;
    } 
}
