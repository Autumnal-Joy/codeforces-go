// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1990/problem/C
// https://codeforces.com/problemset/status/1990/problem/C?friends=on
func Test_cf1990C(t *testing.T) {
	testCases := [][2]string{
		{
			`4
1
1
3
2 2 3
4
2 1 1 2
4
4 4 4 4`,
			`1
13
9
40`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1990C)
}
