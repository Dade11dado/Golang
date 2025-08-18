package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumps over the lazy dog"
	rev, revError := Reverse(input)
	doubleRev, doubleRevError := Reverse(rev)
	fmt.Printf("Original: %q\n", input)
	fmt.Printf("Rversed: %q, err: %v\n", rev, revError)
	fmt.Printf("Reversed again: %q, err %v\n", doubleRev, doubleRevError)
}

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("Input is not validUuTF-8")
	}
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}
