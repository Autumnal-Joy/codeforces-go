// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_a(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, evenOddBit, "a.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-337/problems/number-of-even-and-odd-bits/
// https://leetcode.cn/problems/number-of-even-and-odd-bits/
