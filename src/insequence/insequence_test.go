package insequence

import (
	"fmt"
	"testing"
)

func TestInsequence(t *testing.T) {
	res := Insequence("abcdefg hijklmn opqr", "acdefg")
	fmt.Println(res)
	res = Insequence("abc d", "aee abc eec")
	fmt.Println(res)
	res = Insequence("abcd", "ab cd")
	fmt.Println(res)
}