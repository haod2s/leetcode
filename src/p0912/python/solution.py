class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
        self.mergeSort(nums, 0, len(nums)-1)
        return nums

    def mergeSort(self, nums: List[int], start: int, end: int):
        if start == end:
            return
        mid = start + (end-start) // 2
        self.mergeSort(nums, start, mid)
        self.mergeSort(nums, mid+1, end)
        self.merge(nums, start, end)

    def merge(self, nums: List[int], start: int, end: int):
        a = [0] * (end - start + 1)
        mid = start + (end-start) // 2
        i, j, k = start, mid+1, 0
        while i <= mid and j <= end:
            if nums[i] < nums[j]:
                a[k] = nums[i]
                i += 1
            else:
                a[k] = nums[j]
                j += 1
            k += 1
        while i <= mid:
            a[k] = nums[i]
            k += 1
            i += 1
        while j <= end:
            a[k] = nums[j]
            k += 1
            j += 1
        for i in range(len(a)):
            nums[start+i] = a[i]
            
# 华为