class LRUCache:
    
    def __init__(self, capacity: int):
        self.cap = capacity
        self.map = {}
        self.head = DLinkedNode()
        self.tail = DLinkedNode()
        self.head.next = self.tail
        self.tail.prev = self.head


    def get(self, key: int) -> int:
        if key not in self.map:
            return -1
        self.moveToHead(self.map[key])
        
        return self.map[key].val
            

    def put(self, key: int, value: int) -> None:
        node = self.map.get(key)
        if node != None:
            node.val = value
            self.moveToHead(node)
            return
        if len(self.map) + 1 > self.cap:
            del self.map[self.tail.prev.key]
            self.removeNode(self.tail.prev)
        node = DLinkedNode(key, value)
        self.map[key] = node
        self.addToHead(node)

        
    def addToHead(self, node) -> None:
        node.next = self.head.next
        self.head.next.prev = node
        self.head.next = node
        node.prev = self.head
     
    
    def removeNode(self, node) -> None:
        if node == self.head or node == self.tail:
            return
        node.prev.next = node.next
        node.next.prev = node.prev
    
    
    def moveToHead(self, node) -> None:
        if node == self.head.next:
            return
        self.removeNode(node)
        self.addToHead(node)


class DLinkedNode:
    def __init__(self, key=0, value=0):
        self.key = key
        self.val = value
        
# 华为