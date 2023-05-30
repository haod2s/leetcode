class Solution:
    def partitionLabels(self, s: str) -> List[int]:
        m = {}
        for i, c in enumerate(s):
            if c not in m:
                m[c] = [i, i]
            else:
                m[c][1] = i
        seq = sorted(m.values(), key=lambda x: x[0])
        merged = [seq[0]]
        for i in range(1, len(seq)):
            if seq[i][0] > merged[-1][1]:
                merged.append(seq[i])
            elif seq[i][0] < merged[-1][1] and seq[i][1] > merged[-1][1]:
                merged[-1][1] = seq[i][1]
                
        return [v[1] - v[0] + 1 for v in merged]
    
# 腾讯