// 146 LRU_Cache Hard
import java.util.*;

public class LRUCache {
    HashMap<Integer, Tuple> maps;
    Tuple tail = new Tuple(-1, -1);
    Tuple head = new Tuple(-1, -1);
    int capacity;
    
    public LRUCache(int capacity) {
        this.maps = new HashMap<Integer, Tuple>(capacity);
        this.capacity = capacity;
        this.tail.next = head;
        this.head.last = tail;
    }
    
    public int get(int key) {
        Tuple t = maps.get(key);
        if(t == null)return -1;
        t.last.next = t.next;
        t.next.last = t.last;
        t.next = head;
        t.last = head.last;
        head.last.next = t;
        head.last = t;
        return t.value;
    }
    
    public void set(int key, int value) {
        if(maps.containsKey(key)){
            Tuple t = maps.get(key);
            t.value = value;
            t.last.next = t.next;
            t.next.last = t.last;
            t.next = head;
            t.last = head.last;
            head.last.next = t;
            head.last = t;
        }else {
            if(maps.size() == capacity){
                int tmpK = tail.next.key;
                tail.next = tail.next.next;
                tail.next.last = tail;
                maps.remove(tmpK);
            }
            Tuple t = new Tuple(key, value);
            t.last = head.last;
            t.next = head;
            head.last.next = t;
            head.last = t;
            maps.put(key, t);
        }
    }
}

class Tuple {
    int key;
    int value;
    Tuple last;
    Tuple next;
    
    Tuple(int key, int value){
        this.key = key;
        this.value= value;
        last = null;
        next = null;
    }
}
