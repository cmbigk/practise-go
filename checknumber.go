package piscine

func checknumber(k string) bool {
	for _, c := range k {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}
