func removeKdigits(num string, k int) string {
	n := len(num)
	if k >= n {
		return "0"
	}

	st := []byte{}

	for i := 0; i < n; i++ {
		d := num[i]

		for len(st) > 0 && k > 0 && st[len(st)-1] > d {
			st = st[:len(st)-1]
			k--
		}

		st = append(st, d)
	}

	// If k still > 0, remove from end
	st = st[:len(st)-k]

	// Remove leading zeros
	idx := 0
	for idx < len(st) && st[idx] == '0' {
		idx++
	}

	if idx == len(st) {
		return "0"
	}

	return string(st[idx:])
}
