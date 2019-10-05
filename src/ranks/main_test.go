package main

import (
	"fmt"
	"testing"
	"encoding/json"
)

// func TestRankProvider(t *testing.T) {

// 	sources := []Source{
// 		Source("abcd efg"),
// 	}
// 	targets := []Target{
// 		Target{Text: "aaaa", Key: "1"},
// 		Target{Text: "abcd", Key: "2"},
// 		Target{Text: "abce", Key: "3"},
// 		Target{Text: "abcefg", Key: "4"},
// 		Target{Text: "abcd efg", Key: "5"},
// 		// The following also has a rank of 1 due to a limitation having to weighting subMatch against their source word,
// 		// instead of weighing against the target string. This is acceptable so far.
// 		// Target{Text: "abcdefg", Data: "6"},
// 		Target{Text: "acd fg", Key: "7"},
// 	}

// 	res := MatchSourceAndTargets(sources, targets)
// 	fmt.Println(res)
// }

func TestCombine(t *testing.T) {
	var data = `{"targets":[{"text":"C Class C 240 Sedan 4D","key":"https://www.kbb.com/mercedes-benz/c-class/2001/c-240-sedan-4d/?vehicleid=4683&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"C Class C 320 Sedan 4D","key":"https://www.kbb.com/mercedes-benz/c-class/2001/c-320-sedan-4d/?vehicleid=4687&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"CL Class CL 55 Coupe 2D","key":"https://www.kbb.com/mercedes-benz/cl-class/2001/cl-55-coupe-2d/?vehicleid=4676&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"CL Class CL 500 Coupe 2D","key":"https://www.kbb.com/mercedes-benz/cl-class/2001/cl-500-coupe-2d/?vehicleid=4670&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"CL Class CL 600 Coupe 2D","key":"https://www.kbb.com/mercedes-benz/cl-class/2001/cl-600-coupe-2d/?vehicleid=4681&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"CLK Class CLK 320 Coupe 2D","key":"https://www.kbb.com/mercedes-benz/clk-class/2001/clk-320-coupe-2d/?vehicleid=4690&intent=buy-used&category=coupe&modalview=false&pricetype=private-party&condition=good"},{"text":"CLK Class CLK 430 Coupe 2D","key":"https://www.kbb.com/mercedes-benz/clk-class/2001/clk-430-coupe-2d/?vehicleid=4695&intent=buy-used&category=coupe&modalview=false&pricetype=private-party&condition=good"},{"text":"CLK Class CLK 55 AMG Coupe 2D","key":"https://www.kbb.com/mercedes-benz/clk-class/2001/clk-55-amg-coupe-2d/?vehicleid=4701&intent=buy-used&category=coupe&modalview=false&pricetype=private-party&condition=good"},{"text":"CLK Class CLK 320 Cabriolet 2D","key":"https://www.kbb.com/mercedes-benz/clk-class/2001/clk-320-cabriolet-2d/?vehicleid=4691&intent=buy-used&category=convertible&modalview=false&pricetype=private-party&condition=good"},{"text":"CLK Class CLK 430 Cabriolet 2D","key":"https://www.kbb.com/mercedes-benz/clk-class/2001/clk-430-cabriolet-2d/?vehicleid=4696&intent=buy-used&category=convertible&modalview=false&pricetype=private-party&condition=good"},{"text":"E Class E 320 AWD Sedan 4D","key":"https://www.kbb.com/mercedes-benz/e-class/2001/e-320-awd-sedan-4d/?vehicleid=350166&intent=buy-used&category=sedan&modalview=false&pricetype=private-party&condition=good"},{"text":"E Class E 320 Sedan 4D","key":"https://www.kbb.com/mercedes-benz/e-class/2001/e-320-sedan-4d/?vehicleid=350161&intent=buy-used&category=sedan&modalview=false&pricetype=private-party&condition=good"},{"text":"E Class E 430 Sedan 4D","key":"https://www.kbb.com/mercedes-benz/e-class/2001/e-430-sedan-4d/?vehicleid=350163&intent=buy-used&category=sedan&modalview=false&pricetype=private-party&condition=good"},{"text":"E Class E 430 AWD Sedan 4D","key":"https://www.kbb.com/mercedes-benz/e-class/2001/e-430-awd-sedan-4d/?vehicleid=350159&intent=buy-used&category=sedan&modalview=false&pricetype=private-party&condition=good"},{"text":"E Class E 55 AMG Sedan 4D","key":"https://www.kbb.com/mercedes-benz/e-class/2001/e-55-amg-sedan-4d/?vehicleid=4714&intent=buy-used&category=sedan&modalview=false&pricetype=private-party&condition=good"},{"text":"E Class E 320 Wagon 4D","key":"https://www.kbb.com/mercedes-benz/e-class/2001/e-320-wagon-4d/?vehicleid=350162&intent=buy-used&category=wagon&modalview=false&pricetype=private-party&condition=good"},{"text":"E Class E 320 AWD Wagon 4D","key":"https://www.kbb.com/mercedes-benz/e-class/2001/e-320-awd-wagon-4d/?vehicleid=350158&intent=buy-used&category=wagon&modalview=false&pricetype=private-party&condition=good"},{"text":"M Class ML 320 Sport Utility 4D","key":"https://www.kbb.com/mercedes-benz/m-class/2001/ml-320-sport-utility-4d/?vehicleid=4734&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"M Class ML 430 Sport Utility 4D","key":"https://www.kbb.com/mercedes-benz/m-class/2001/ml-430-sport-utility-4d/?vehicleid=4728&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"M Class ML 55 AMG Sport Utility 4D","key":"https://www.kbb.com/mercedes-benz/m-class/2001/ml-55-amg-sport-utility-4d/?vehicleid=4735&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"S Class S 430 Sedan 4D","key":"https://www.kbb.com/mercedes-benz/s-class/2001/s-430-sedan-4d/?vehicleid=4712&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"S Class S 500 Sedan 4D","key":"https://www.kbb.com/mercedes-benz/s-class/2001/s-500-sedan-4d/?vehicleid=4718&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"S Class S 600 Sedan 4D","key":"https://www.kbb.com/mercedes-benz/s-class/2001/s-600-sedan-4d/?vehicleid=4717&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"S Class S 55 Sedan 4D","key":"https://www.kbb.com/mercedes-benz/s-class/2001/s-55-sedan-4d/?vehicleid=4715&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"SL Class SL 500 Roadster 2D","key":"https://www.kbb.com/mercedes-benz/sl-class/2001/sl-500-roadster-2d/?vehicleid=4720&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"SL Class SL 600 Roadster 2D","key":"https://www.kbb.com/mercedes-benz/sl-class/2001/sl-600-roadster-2d/?vehicleid=4724&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"SLK Class SLK 230 Roadster 2D","key":"https://www.kbb.com/mercedes-benz/slk-class/2001/slk-230-roadster-2d/?vehicleid=4722&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"SLK Class SLK 320 Roadster 2D","key":"https://www.kbb.com/mercedes-benz/slk-class/2001/slk-320-roadster-2d/?vehicleid=4727&intent=buy-used&modalview=false&pricetype=private-party&condition=good"}],"sources":["Mercedes Benz S430 AMG","Mercedes Benz S430 AMG $1499 OBO sedan"]}`

	ranksRequest := RanksRequest{}
	json.Unmarshal([]byte(data), &ranksRequest)
	rankedTarget := MatchSourceAndTargets(ranksRequest.Sources, ranksRequest.Targets)
	fmt.Println(rankedTarget)
}
