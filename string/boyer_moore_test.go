package string

import (
	"strings"
	"testing"
)

//BoyerMoore

func Test_BoyerMoore(t *testing.T) {
	t.Log(BoyerMoore(s3, t3))

	equal(t, BoyerMoore(s3, t3) == strings.Index(s3, t3), "%v in %v find error", s3, t3)

	equal(t, BoyerMoore(s2, t2) == strings.Index(s2, t2), "%v in %v find error", s2, t2)

	equal(t, BoyerMoore(s1, t1) == strings.Index(s1, t1), "%v in %v find error", s1, t1)

	equal(t, BoyerMoore(s0, t0) == strings.Index(s0, t0), "%v in %v find error", s0, t0)
}
