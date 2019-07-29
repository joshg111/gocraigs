package insequence

import (
	"fmt"
	"testing"
)

func TestInsequence(t *testing.T) {
	
	// res := Insequence("Chevy Prism (For Parts) sedan", "infiniti")
	res := Insequence("van", "land")
	// res := Insequence("ii", "i")
	fmt.Println(res)
}