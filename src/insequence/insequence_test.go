package insequence

import (
	"fmt"
	"testing"
)

func TestInsequence(t *testing.T) {
	
	res := Insequence("M3",
"BMW 328i Coupe M Sport Pakage Low Mileage")
	fmt.Println(res)
}