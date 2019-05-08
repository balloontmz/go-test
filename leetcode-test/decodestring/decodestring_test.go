package decodestring

import (
	"testing"
)

//DecodeString 解码字符串
func TestDecodeString(t *testing.T) {
	/*
	s = "3[a]2[bc]", return "aaabcbc".
	s = "3[a2[c]]", return "accaccacc".
	s = "2[abc]3[cd]ef", return "abcabccdcdcdef".
	*/
	s := "3[a]2[bc]"
	r := DecodeString(s)

	if r != "aaabcbc" {
		t.Errorf("want %v, get %v", "aaabcbc", r)
	}

	s = "3[a2[c]]"
	r = DecodeString(s)

	if r != "aaabcbc" {
		t.Errorf("want %v, get %v", "accaccacc", r)
	}

	s = "2[abc]3[cd]ef"
	r = DecodeString(s)

	if r != "aaabcbc" {
		t.Errorf("want %v, get %v", "abcabccdcdcdef", r)
	}
}