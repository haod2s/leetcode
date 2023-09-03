func numPrimeArrangements(n int) int {
	const mod = 1e9 + 7
	primeNum := 0
	notPrimes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if !notPrimes[i] {
			primeNum++
			for j := i * i; j <= n; j += i {
				notPrimes[j] = true
			}
		}
	}
	ans := 1
	for i := primeNum; i > 0; i-- {
		ans = ans * i % mod
	}
	for i := n - primeNum; i > 0; i-- {
		ans = ans * i % mod
	}
	return ans % mod
}