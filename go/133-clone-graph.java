/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> neighbors;

    public Node() {}

    public Node(int _val,List<Node> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
};
*/
import java.util.Stack;

class Solution {
    public Node cloneGraph(Node node) {
        HashMap<Node, Node> map = new HashMap<Node, Node>();
        Stack<Node> stack = new Stack<Node>();
        stack.push(node);
        Node nodeCopy = new Node(node.val, new ArrayList<Node>(node.neighbors.size()));
        map.put(node, nodeCopy);
        while (!stack.empty()) {
            Node v = stack.pop();
            Node vCopy = map.get(v);
            for (Node link: v.neighbors) {
                Node linkCopy = map.get(link);
                if (linkCopy == null) {
                    linkCopy = new Node(link.val, new ArrayList<Node>(link.neighbors.size()));
                    map.put(link, linkCopy);
                    stack.push(link);
                }    
                vCopy.neighbors.add(linkCopy);
            }
        }
        return map.get(node);                          
    }
}
