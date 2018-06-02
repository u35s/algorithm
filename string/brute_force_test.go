package string

import "testing"

func Test_BruteForce(t *testing.T) {
	s1 := "aaaabaaaac"
	t1 := "aaaac"
	t.Log(BruteForce(s1, t1))

	s2 := "aaaabaaaac"
	t2 := "aaaad"
	t.Log(BruteForce(s2, t2))
}
