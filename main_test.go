package main

import "testing"

func TestMath(t *testing.T) {
	if 1+1 == 2 {
		t.Error("Math is broken!") // 只要这里报错，GitHub 的合并按钮就会被锁死
	}
}
