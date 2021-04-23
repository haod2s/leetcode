// 给你一个由 无重复 正整数组成的集合 nums ，请你找出并返回其中最大的整除子集 answer ，子集中每一元素对 (answer[i], answer[j]) 都应当满足：
// answer[i] % answer[j] == 0 ，或
// answer[j] % answer[i] == 0
// 如果存在多个有效解子集，返回其中任何一个均可。
// 
//  
// 
// 示例 1：
// 
// 输入：nums = [1,2,3]
// 输出：[1,2]
// 解释：[1,3] 也会被视为正确答案。
// 示例 2：
// 
// 输入：nums = [1,2,4,8]
// 输出：[1,2,4,8]
//  
// 
// 提示：
// 
// 1 <= nums.length <= 1000
// 1 <= nums[i] <= 2 * 109
// nums 中的所有整数 互不相同

func largestDivisibleSubset(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	maxSize, maxVal := 1, 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] % nums[j] == 0 && dp[j]+1 > dp[i]{
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxSize {
			maxVal, maxSize = nums[i], dp[i]
		}
	}
	if maxSize == 1 {
        return []int{nums[0]}
    }
	res := make([]int, 0)

    for i := n - 1; i >= 0 && maxSize > 0; i-- {
        if dp[i] == maxSize && maxVal%nums[i] == 0 {
            res = append(res, nums[i])
            maxVal = nums[i]
            maxSize--
        }
    }

	return res
}