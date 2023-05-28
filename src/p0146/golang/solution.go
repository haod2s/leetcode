type LRUCache struct {
	cap  int
	m    map[int]*DLinkedNode
	head *DLinkedNode
	tail *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	obj := LRUCache{
		cap:  capacity,
		m:    make(map[int]*DLinkedNode),
		head: &DLinkedNode{},
		tail: &DLinkedNode{},
	}
	obj.head.next = obj.tail
	obj.tail.prev = obj.head

	return obj
}

func (this *LRUCache) Get(key int) int {
	o, ok := this.m[key]
	if !ok {
		return -1
	}
	this.moveToHead(o)
	return o.val
}

func (this *LRUCache) Put(key int, value int) {
	o, ok := this.m[key]
	if ok {
		o.val = value
		this.moveToHead(o)
		return
	}
	if len(this.m)+1 > this.cap {
		delete(this.m, this.tail.prev.key)
		this.removeNode(this.tail.prev)
	}
	node := &DLinkedNode{
		key: key,
		val: value,
	}
	this.addToHead(node)
	this.m[key] = node
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	if this.head.next == node {
		return
	}
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	if node == nil || node == this.head || node == this.tail {
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	o := this.head.next
	node.next = o
	o.prev = node
	this.head.next = node
	node.prev = this.head
}

type DLinkedNode struct {
	key  int
	val  int
	prev *DLinkedNode
	next *DLinkedNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// 腾讯