package string

import (
	"strings"
	"testing"
)

// next=x00012
var s3 = "abcabcabaabcabxc"
var t3 = "abcabx"

// next=x01
var s2 = "acabbaaaab"
var t2 = "aab"

// next=x01234
var s1 = "acabbaaaaxaaaaaab"
var t1 = "aaaaaab"

// next=x0100120
var s0 = "acabbaaaaxxaabaaaaabaaxxaabb"
var t0 = "aaxxaabb"

func equal(t *testing.T, b bool, format string, args ...interface{}) {
	if !b {
		t.Errorf(format, args...)
	}
}

func Test_KMP(t *testing.T) {
	equal(t, KMP_1(s3, t3) == strings.Index(s3, t3), "%v in %v find error", s3, t3)
	equal(t, KMP_2(s3, t3) == strings.Index(s3, t3), "%v in %v find error", s3, t3)

	equal(t, KMP_1(s2, t2) == strings.Index(s2, t2), "%v in %v find error", s2, t2)
	equal(t, KMP_2(s2, t2) == strings.Index(s2, t2), "%v in %v find error", s2, t2)

	equal(t, KMP_1(s1, t1) == strings.Index(s1, t1), "%v in %v find error", s1, t1)
	equal(t, KMP_2(s1, t1) == strings.Index(s1, t1), "%v in %v find error", s1, t1)

	equal(t, KMP_1(s0, t0) == strings.Index(s0, t0), "%v in %v find error", s0, t0)
	equal(t, KMP_2(s0, t0) == strings.Index(s0, t0), "%v in %v find error", s0, t0)
}

func Benchmark_KMP_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//KMP_2(s3, t3)
		//KMP_2(s2, t2)
		KMP_2(s1, t1)
		//KMP_2(s0, t0)
	}
}

func Benchmark_KMP_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//KMP_1(s3, t3)
		//KMP_1(s2, t2)
		KMP_1(s1, t1)
		//KMP_1(s0, t0)
	}
}
