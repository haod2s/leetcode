struct NumArray {
	dp: Vec<i32>,
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl NumArray {

    fn new(nums: Vec<i32>) -> Self {
		let mut dp: Vec<i32> = vec![0; nums.len()];
		dp[0] = nums[0];
		for i in 1..nums.len() {
			dp[i] = dp[i-1] + nums[i];
		}

		NumArray {dp: dp}
    }
    
    fn sum_range(&self, left: i32, right: i32) -> i32 {
		if left == 0 {
			return self.dp[right as usize];
		}
		let i = (left - 1) as usize;
		let j = right as usize;
		
		self.dp[j] - self.dp[i]
    }
}

/**
 * Your NumArray object will be instantiated and called as such:
 * let obj = NumArray::new(nums);
 * let ret_1: i32 = obj.sum_range(left, right);
 */