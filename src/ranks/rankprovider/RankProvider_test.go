package rankprovider

import (
	"testing"
	"fmt"
)

func TestRankProvider(t *testing.T) {
	// target - 2 / 8
	// source - 3 / 5
	res := Get("3 Series 328i Coupe 2D",
"BMW 328i Coupe M Sport Pakage Low Mileage")
	fmt.Println(res)

	// target - 2 / 8
	// source - 
	res = Get("M3 Coupe 2D",
"BMW 328i Coupe M Sport Pakage Low Mileage")
	fmt.Println(res)

	// r1 := _findSourceTokens([]string{"M3"}, "BMW 328i Coupe M Sport Pakage Low Mileage");
	// fmt.Println(r1)
}