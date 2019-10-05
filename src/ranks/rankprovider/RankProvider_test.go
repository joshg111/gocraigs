package rankprovider

import (
	"testing"
	"fmt"
)

func TestRankProvider(t *testing.T) {
	// target - 2 / 8
	// source - 8 / 
	// res := Get("Mercedes Benz S430 AMG $1499 OBO sedan", "E Class E 55 AMG Sedan 4D")
	res := Get("Mercedes Benz S430 AMG $1499 OBO sedan", "S Class S 430 Sedan 4D")
	// res := Get("Mercedes Benz S430", "S Class S 430")
	// res := Get("CHEVY EQUINOX LT AWD 4DR SUV/// FACTORY WARRANTY SUV", "Equinox LS Sport Utility 4D")
	fmt.Println(res)

	// target - 2 / 8
	// source - 
// 	res = Get("M3 Coupe 2D",
// "BMW 328i Coupe M Sport Pakage Low Mileage")
// 	fmt.Println(res)

	// r1 := _findSourceTokens([]string{"infiniti"}, "Chevy Prism (For Parts) sedan");
	// r1 := _findSourceTokens([]string{"chevrolet"}, "Chevy 2500 Cargo Van");
	// r1 := _findSourceTokens([]string{"land", "rover"}, "Chevy 2500 Cargo Van");
	// r1 := _findSourceTokens([]string{"land", "rover"}, "2007 Chevy 2500 Cargo Van");
	// r1 := _findSourceTokens([]string{"land", "rover"}, "2007 Chevy 2500 Cargo Van");
	// r1 := _findSourceTokens([]string{"van"}, "an");
	
	// fmt.Println(r1)
}