package tokenmatch

import (
    "strings"
)


// TokenMatch - Token string which was matched, and it's corresponding weight and insequence match
type TokenMatch struct {
    Token string
    Weight float32
    Match string
}

// TokenList is a list of TokenMatch
type TokenList struct {
    TokenMatches []TokenMatch
}

func (t TokenList) JoinMatch() string {
    var arr []string
    for _, m := range t.TokenMatches {
        arr = append(arr, m.Match)
    }

    return strings.Join(arr, " ")
}

func (t TokenList) AverageWeight() float32 {
    var sum float32
    for _,m := range t.TokenMatches {
        if m.Weight > .6 {
            sum += m.Weight
        }
    }
    return sum / float32(len(t.TokenMatches));
}
