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
// Simplest recursion with return value.
class Solution {
    HashMap<Node, Node> map = new HashMap<>();
    public Node cloneGraph(Node node) {
        if (map.get(node) != null)
            return map.get(node);
        Node newNode = new Node();
        newNode.val = node.val;
        newNode.neighbors = new ArrayList<Node>(node.neighbors.size());
        map.put(node, newNode);
        for (Node next: node.neighbors) {
            newNode.neighbors.add(cloneGraph(next));
        }
        return newNode;
    }
}

import java.util.Stack;

// DFS iteration
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

// DFS recursion
class Solution {
    public Node cloneGraph(Node node) {
        HashMap<Node, Node> map = new HashMap<Node, Node>();
        Node nodeCopy = new Node(node.val, new ArrayList<Node>());
        map.put(node, nodeCopy);
        DFS(node, map);
        return map.get(node);
    }
    
    void DFS(Node root, HashMap<Node, Node> map) {
        Node rootCopy = map.get(root);
        for (Node neighbor: root.neighbors) {
            Node neighborCopy = map.get(neighbor);
            if (neighborCopy == null) {
                neighborCopy = new Node(neighbor.val, new ArrayList<Node>());
                map.put(neighbor, neighborCopy);
                DFS(neighbor, map);
            }
            rootCopy.neighbors.add(neighborCopy);
        }
    }
}
