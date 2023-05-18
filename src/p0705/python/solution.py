class MyHashSet:

    def __init__(self):
        self.arr = bytearray(10**6+1)

    def add(self, key: int) -> None:
        self.arr[key] = 1        

    def remove(self, key: int) -> None:
        self.arr[key] = 0


    def contains(self, key: int) -> bool:
        return self.arr[key] == 1

# 华为