class KthLargest:

    def __init__(self, k: int, nums: List[int]):
        self.h = []
        for v in nums:
            if len(self.h) < k:
                heapq.heappush(self.h, v)
            elif v > self.h[0]:
                heapq.heapreplace(self.h, v)
        for i in range(k-len(self.h)):
            heapq.heappush(self.h, -math.inf)

    def add(self, val: int) -> int:
        if val > self.h[0]:
            heapq.heapreplace(self.h, val)
        
        return self.h[0]



# Your KthLargest object will be instantiated and called as such:
# obj = KthLargest(k, nums)
# param_1 = obj.add(val)