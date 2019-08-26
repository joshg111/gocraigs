package insequence

import (
	"fmt"
	"testing"
)

func TestInsequence(t *testing.T) {
	
	// res := Insequence("Chevy Prism (For Parts) sedan", "infiniti")
	// res := Insequence("Equinox", "CHEVY EQUINOX LT AWD 4DR SUV/// FACTORY WARRANTY SUV")
	res := Insequence("EQ", "EEQ")
	// res := Insequence("ii", "i")
	fmt.Println(res)
}