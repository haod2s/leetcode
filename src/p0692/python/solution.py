class Solution:
    def topKFrequent(self, words: List[str], k: int) -> List[str]:
        m = dict()
        for w in words:
            if w not in m:
                m[w] = 1
            else:
                m[w] += 1
        temp = sorted(m.items(), key=lambda x: (-x[1], x[0]))
        ans = [t[0] for t in temp[:k]]
        
        return ans