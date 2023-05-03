class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        t = [0] * (m + n)
        i, j, k = 0, 0, 0
        while i < m and j < n:
            if nums1[i] < nums2[j]:
                t[k] = nums1[i]
                i += 1
            else:
                t[k] = nums2[j]
                j += 1
            k += 1
        while i < m:
            t[k] = nums1[i]
            i += 1
            k += 1
        while j < n:
            t[k] = nums2[j]
            j += 1
            k += 1
        for i in range(m+n):
            nums1[i] = t[i]