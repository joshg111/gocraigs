package insequence

import (
	"fmt"
	"testing"
)

func TestInsequence(t *testing.T) {
	
	// res := Insequence("Chevy Prism (For Parts) sedan", "infiniti")
	// res := Insequence("Equinox", "CHEVY EQUINOX LT AWD 4DR SUV/// FACTORY WARRANTY SUV")
	// res := Insequence("S430", "S 430")
	res := Insequence("S4", "S 4")
	// res := Insequence("ii", "i")
	fmt.Println(res)
}