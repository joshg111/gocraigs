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
	var data = `{"targets":[{"text":"Equinox LS Sport Utility 4D","key":"https://www.kbb.com/chevrolet/equinox/2019/lt-sport-utility-4d/?vehicleid=436730&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"Impala LS Sedan 4D","key":"https://www.kbb.com/chevrolet/impala/2019/ls-sedan-4d/?vehicleid=436764&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"Impala LT Sedan 4D","key":"https://www.kbb.com/chevrolet/impala/2019/lt-sedan-4d/?vehicleid=436737&intent=buy-used&modalview=false&pricetype=private-party&condition=good"},{"text":"Impala Premier Sedan 4D","key":"https://www.kbb.com/chevrolet/impala/2019/premier-sedan-4d/?vehicleid=436770&intent=buy-used&modalview=false&pricetype=private-party&condition=good"}],"sources":["CHEVROLET EQUINOX LT AWD","CHEVY EQUINOX LT AWD 4DR SUV/// FACTORY WARRANTY SUV"]}`

	ranksRequest := RanksRequest{}
	json.Unmarshal([]byte(data), &ranksRequest)
	rankedTarget := MatchSourceAndTargets(ranksRequest.Sources, ranksRequest.Targets)
	fmt.Println(rankedTarget)
}
