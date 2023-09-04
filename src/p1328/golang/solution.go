func breakPalindrome(palindrome string) string {
	if len(palindrome) == 1 {
		return ""
	}
	ans := []byte(palindrome)
	found := false
	for i := 0; !found && i < len(ans)>>1; i++ {
		if ans[i] != 'a' {
			ans[i] = 'a'
			found = true
		}
	}
	if !found {
		ans[len(ans)-1] = 'b'
	}
	return string(ans)
}