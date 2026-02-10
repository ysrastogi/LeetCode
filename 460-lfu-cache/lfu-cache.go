type Node struct{
    key int
    val int
    freq int
    prev *Node
    next *Node
}
type DoublyLinkedList struct {
    head *Node
    tail *Node
}
func (l *DoublyLinkedList) removeFromTail() *Node {
	if l.tail.prev == nil || l.tail.prev.prev == nil {
		return nil
	}
	node := l.tail.prev
	l.tail.prev = node.prev
	node.prev.next = l.tail

	node.prev = nil
	node.next = nil
	return node
}
func (l *DoublyLinkedList) addNode(node *Node) {
	first := l.head.next

	l.head.next = node
	node.prev = l.head
	node.next = first
	first.prev = node
}
func (l *DoublyLinkedList) len() int {
	p := l.head.next
	length := 0
	for p != l.tail {
		length += 1
		p = p.next
	}
	return length
}

func (l *DoublyLinkedList) moveToFront(k int) {
	p := l.head.next

	for p != l.tail && p.key != k {
		p = p.next
	}

}
func removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev

	node.prev = nil
	node.next = nil
}

type LFUCache struct {
	capacity int
	size     int
	cache    map[int]*Node
	freqMap  map[int]*DoublyLinkedList
	minFreq  int
}

func Constructor(capacity int) LFUCache {
    return LFUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		freqMap:  make(map[int]*DoublyLinkedList),
	}
    
}


func (c *LFUCache) Get(key int) int {
    if node, ok := c.cache[key]; ok {

		c.cache[key].freq += 1
		freq := node.freq
		removeNode(node)
		if c.freqMap[freq] == nil {
			c.freqMap[freq] = &DoublyLinkedList{
				head: &Node{},
				tail: &Node{},
			}
			c.freqMap[freq].head.next = c.freqMap[freq].tail
			c.freqMap[freq].tail.prev = c.freqMap[freq].head
		}
		c.freqMap[freq].addNode(node)

		if c.freqMap[c.minFreq] != nil && c.freqMap[c.minFreq].len() == 0 {
			c.minFreq += 1
		}

		return node.val
	}
	return -1
    
}


func (c *LFUCache) Put(key int, value int)  {
    if c.capacity == 0 {
		return
	}
	if node, ok := c.cache[key]; ok {
		node.val = value
		c.Get(key)
		return
	}
	if c.size == c.capacity {
		evictedNode := c.freqMap[c.minFreq].removeFromTail()
		delete(c.cache, evictedNode.key)
		c.size -= 1
	}
	newNode := &Node{
		key:  key,
		val:  value,
		freq: 1,
	}
	c.cache[key] = newNode
	if c.freqMap[1] == nil {
		c.freqMap[1] = &DoublyLinkedList{
			head: &Node{},
			tail: &Node{},
		}
		c.freqMap[1].head.next = c.freqMap[1].tail
		c.freqMap[1].tail.prev = c.freqMap[1].head
	}
	c.freqMap[1].addNode(newNode)
	c.minFreq = 1
	c.size += 1
    
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */