const MOD int64 = 1000000007

func modPow(base, exp int64) int64 {
	result := int64(1)
	base %= MOD

	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % MOD
		}
		base = (base * base) % MOD
		exp /= 2
	}

	return result
}

func countGoodNumbers(n int64) int {
	evenPositions := (n + 1) / 2
	oddPositions := n / 2

	evenPart := modPow(5, evenPositions)
	oddPart := modPow(4, oddPositions)

	return int((evenPart * oddPart) % MOD)
}

