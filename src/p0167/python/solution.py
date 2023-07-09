class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        i, j = 1, len(numbers) 
        while i < j:
            res = numbers[i-1] + numbers[j-1]
            if res == target:
                return [i, j]
            elif res > target:
                j -= 1
            else:
                i += 1
                
        return None