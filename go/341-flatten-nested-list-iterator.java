/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * public interface NestedInteger {
 *
 *     // @return true if this NestedInteger holds a single integer, rather than a nested list.
 *     public boolean isInteger();
 *
 *     // @return the single integer that this NestedInteger holds, if it holds a single integer
 *     // Return null if this NestedInteger holds a nested list
 *     public Integer getInteger();
 *
 *     // @return the nested list that this NestedInteger holds, if it holds a nested list
 *     // Return null if this NestedInteger holds a single integer
 *     public List<NestedInteger> getList();
 * }
 */
public class NestedIterator implements Iterator<Integer> {
    Stack<Iterator<NestedInteger>> stack = new Stack<Iterator<NestedInteger>>();
    Integer curr;
    
    public NestedIterator(List<NestedInteger> nestedList) {
        if (nestedList != null) {
            stack.push(nestedList.iterator());    
        }   
    }

    @Override
    public Integer next() {
        return curr;
    }

    @Override
    public boolean hasNext() {
        while (!stack.isEmpty()) { 
            Iterator<NestedInteger> iter = stack.peek();
            if (iter.hasNext()) {
                 NestedInteger ni = iter.next();
                if (ni.isInteger()) {
                    curr = ni.getInteger();
                    return true;
                } else {
                    stack.push(ni.getList().iterator());
                }
            } else {
                stack.pop();
            }   
        }
        return false;
    }
}

/**
 * Your NestedIterator object will be instantiated and called as such:
 * NestedIterator i = new NestedIterator(nestedList);
 * while (i.hasNext()) v[f()] = i.next();
 */
