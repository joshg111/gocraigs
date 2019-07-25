package main


import (
	"testing"
	"fmt"
)

func TestRankProvider(t *testing.T) {

	sources := []Source{
		Source("abcd efg"),
	}
	targets := []Target{
		Target{Text: "aaaa", Data: "1"},
		Target{Text: "abcd", Data: "2"},
		Target{Text: "abce", Data: "3"},
		Target{Text: "abcefg", Data: "4"},
		Target{Text: "abcd efg", Data: "5"},
		// The following also has a rank of 1 due to a limitation having to weighting subMatch against their source word,
		// instead of weighing against the target string. This is acceptable so far.
		// Target{Text: "abcdefg", Data: "6"},
		Target{Text: "acd fg", Data: "7"},
	}

	res := MatchSourceAndTargets(sources, targets)
	fmt.Println(res)
}