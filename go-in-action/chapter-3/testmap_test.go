package main

import (
	"testing"
)

//DecodeString 解码字符串
func TestTestMap(t *testing.T) {
	/*
		s = "3[a]2[bc]", return "aaabcbc".
		s = "3[a2[c]]", return "accaccacc".
		s = "2[abc]3[cd]ef", return "abcabccdcdcdef".
	*/
	testMap()

	if true {
		t.Errorf("want %v, get", "abcabccdcdcdef")
	}
}
