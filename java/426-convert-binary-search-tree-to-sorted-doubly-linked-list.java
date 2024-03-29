/*
// Definition for a Node.
class Node {
    public int val;
    public Node left;
    public Node right;

    public Node() {}

    public Node(int _val,Node _left,Node _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};
*/
class Solution {
    Node dummy = new Node();
    Node prev = dummy;
    public Node treeToDoublyList(Node root) {
        if (root == null) return null;
        inOrderTreeTraversal(root);
        Node head = dummy.right;
        head.left = prev;
        prev.right = head;
        return head;
    }
    
    // root could be null
    public void inOrderTreeTraversal(Node root) {
        if (root == null) return;
        inOrderTreeTraversal(root.left);
        prev.right = root;
        root.left = prev;
        prev = root;
        inOrderTreeTraversal(root.right);
    }
}
