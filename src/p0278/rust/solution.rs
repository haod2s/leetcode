// The API isBadVersion is defined for you.
// isBadVersion(versions:i32)-> bool;
// to call it use self.isBadVersion(versions)

impl Solution {
    pub fn first_bad_version(&self, n: i32) -> i32 {
		let mut start: i32 = 1;
		let mut end: i32 = n;
		while start <= end {
			let v: i32 = start + (end - start) / 2;
			if self.isBadVersion(v) {
				if start == end { break; }
				end = v - 1;
			} else {
				start = v + 1;
			}
		}
		
		start
    }
}