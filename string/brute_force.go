package string

func BruteForce(s, t string) int {
	for i := 0; i < len(s); {
		for j := 0; j < len(t); {
			if s[i] == t[j] {
				i++
				j++
			} else {
				i = i - j + 1
				j = 0
			}
			if j >= len(t) {
				return i - j
			}
			if i >= len(s) {
				return -1
			}
		}
	}
	return -1
}
