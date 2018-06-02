package string

func getNext2(t string) []int {
	next := make([]int, len(t))
	next[0] = -1
	var j int = -1
	for i := 0; i+1 < len(t); {
		if j == -1 || t[i] == t[j] {
			i++
			j++
			if t[i] != t[j] {
				next[i] = j
			} else {
				next[i] = next[j]
			}
		} else {
			j = next[j]
		}
	}
	return next
}

func getNext1(t string) []int {
	next := make([]int, len(t))
	next[0] = -1
	var j int = -1
	for i := 0; i+1 < len(t); {
		if j == -1 || t[i] == t[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

func kmp(s, t string, next []int) int {
	for i := 0; i < len(s); {
		for j := 0; j < len(t); {
			if j == -1 || s[i] == t[j] {
				i++
				j++
			} else {
				// brute force
				//i = i - j + 1
				//j = 0

				// kmp
				//fmt.Printf("i:%v,j:%v,next:%v\n", i, j, next[j])
				j = next[j]
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

func KMP_1(s, t string) int {
	return kmp(s, t, getNext1(t))
}

func KMP_2(s, t string) int {
	return kmp(s, t, getNext2(t))
}
