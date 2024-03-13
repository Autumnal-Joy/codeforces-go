// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P6329
func Test_p6329(t *testing.T) {
	testCases := [][2]string{
		{
			`8 1
1 10 100 1000 10000 100000 1000000 10000000
1 2
1 3
2 4
2 5
3 6
3 7
3 8
0 3 1`,
			`11100101`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p6329)
}
