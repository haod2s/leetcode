func maskPII(s string) string {
	if strings.IndexByte(s, '@') == -1 {
		return maskPhone(s)
	}
	return maskEmail(s)
}

func maskEmail(s string) string {
	t := strings.ToLower(s)
	items := strings.Split(t, "@")
	return fmt.Sprintf("%c*****%c@%s", items[0][0], items[0][len(items[0])-1], items[1])
}

func maskPhone(s string) string {
	nums := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			nums = append(nums, s[i])
		}
	}
	localNum := string(nums[len(nums)-4:])
	switch len(nums) {
	case 10:
		return fmt.Sprintf("***-***-%s", localNum)
	case 11:
		return fmt.Sprintf("+*-***-***-%s", string(localNum))
	case 12:
		return fmt.Sprintf("+**-***-***-%s", string(localNum))
	}
	return fmt.Sprintf("+***-***-***-%s", string(localNum))
}