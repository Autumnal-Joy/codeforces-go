// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/2028/C
// https://codeforces.com/problemset/status/2028/problem/C?friends=on
func Test_cf2028C(t *testing.T) {
	testCases := [][2]string{
		{
			`7
6 2 1
1 1 10 1 1 10
6 2 2
1 1 10 1 1 10
6 2 3
1 1 10 1 1 10
6 2 10
1 1 10 1 1 10
6 2 11
1 1 10 1 1 10
6 2 12
1 1 10 1 1 10
6 2 12
1 1 1 1 10 10`,
			`22
12
2
2
2
0
-1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2028C)
}
