class Trie:

    def __init__(self):
        self.children = [None] * 26
        self.tail = False

    def insert(self, word: str) -> None:
        p = self
        for c in word:
            i = ord(c) - ord('a')
            if p.children[i] == None:
                p.children[i] = Trie()
            p = p.children[i]
        p.tail = True 

    def search(self, word: str) -> bool:
        p = self
        for c in word:
            i = ord(c) - ord('a')
            if p.children[i] == None:
                return False
            p = p.children[i]
             
        return p != None and p.tail     
        
    def startsWith(self, prefix: str) -> bool:
        p = self
        for c in prefix:
            i = ord(c) - ord('a')
            if p.children[i] == None:
                return False
            p = p.children[i]
        
        return p != None
    
# 华为