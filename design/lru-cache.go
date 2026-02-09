type Node struct{
    key int
    value int
    prev *Node
    next *Node
}

type LRUCache struct {
    capacity int
    cache map[int]*Node
    head *Node
    tail *Node
    
}


func Constructor(capacity int) LRUCache {
    head := &Node{}
    tail := &Node{}

    head.next = tail
    tail.prev = head
    return LRUCache{
        capacity: capacity,
        cache: make(map[int]*Node),
        head: head,
        tail: tail,
    }
    
}

func (l *LRUCache) remove(node *Node){
    prev := node.prev
    next := node.next
    prev.next = next
    next.prev = prev
}

func (l *LRUCache) insertAtHead(node *Node){
    node.next = l.head.next
    node.prev = l.head

    l.head.next.prev = node
    l.head.next = node
}


func (this *LRUCache) Get(key int) int {
    if node, ok := this.cache[key]; ok{
        this.remove(node)
        this.insertAtHead(node)
        return node.value
    }
    return -1
    
}


func (this *LRUCache) Put(key int, value int)  {
    if node, ok := this.cache[key]; ok{
        node.value = value
        this.remove(node)
        this.insertAtHead(node)
        return
    }
    if len(this.cache) == this.capacity{
        lru := this.tail.prev
        this.remove(lru)
        delete(this.cache, lru.key)
    }
    newNode := &Node{
        key:key,
        value:value,
    }
    this.insertAtHead(newNode)
    this.cache[key]= newNode

}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */