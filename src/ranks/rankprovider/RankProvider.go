package rankprovider

import (
    // "fmt"
	"serverless-craigs/src/model/tokenmatch"
	"strings"
	"serverless-craigs/src/insequence"
    "regexp"
)


func toShortStr(s string) string {
	s = strings.ToLower(s)
	re := regexp.MustCompile(`\s|-`)
	s = re.ReplaceAllString(s, " ")
	return s
}

func getTokens(s string) []string {
    // Instead of replacing hyphen with empty string, we could use a space, but if we do that then we wont be able to remove the 
    // matched word/s with the indexes we return.
	s = strings.ToLower(s)
	re := regexp.MustCompile(`-`)
	s = re.ReplaceAllString(s, "")
	return strings.Fields(s)
}

func findSourceTokens(source, target string) tokenmatch.TokenList {
    sourceTokens := getTokens(source)
    shortTarget := toShortStr(target);
    var res = _findSourceTokens(sourceTokens, shortTarget);
    return tokenmatch.TokenList{TokenMatches: res}
}

/**
 * Iterates over sourceTokens and finds the insequence match using
 * the whole target instead of removing from the target. For each
 * insequence we check if it's more than 30% of the original token
 * then we annotate the token with the match.
 */
func _findSourceTokens(sourceTokens []string, shortTarget string) []tokenmatch.TokenMatch {

    var res []tokenmatch.TokenMatch
    for _, sourceToken := range sourceTokens {
        subMatch := insequence.Insequence(sourceToken, shortTarget)
        // Idk why we weight this against only the sourceToken length, it should be the target length. 
        // weight := float32(subMatch.Count) / float32(len(sourceToken))
        // fmt.Println("shortTarget = ", shortTarget)
        // fmt.Println(subMatch)
        weight := float32(subMatch.Count * 2) / float32(subMatch.End - subMatch.Start + 1 + len(sourceToken))
        // fmt.Println(weight)
        res = append(res, tokenmatch.TokenMatch{Token: sourceToken, Weight: weight, Match: subMatch.Match})
    }

    return res;
}

type Tokens struct {
	Source tokenmatch.TokenList
	Target tokenmatch.TokenList
}

func Get(source, target string) float32 {
    var weight float32
    tokens := _triWayTokenMerge(source, target);
    weight = (tokens.Source.AverageWeight() + tokens.Target.AverageWeight()) / 2;
    // fmt.Println(tokens.Source.AverageWeight(), tokens.Target.AverageWeight())
    // fmt.Println(tokens.Source)
    // fmt.Println(tokens.Target)
    // logger.log();
    // logger.log("source = ", source, ", target = ", target);
    // logger.log("sourceTokens = " + sourceTokens);
    // logger.log("targetTokens = " + targetTokens);
    // logger.log("weight = ", weight);
    // logger.log("triWayTokenMerge Time: ", new Date() - startTime);
    // return {weight, sourceTokenMatch, targetTokenMatch};
    // return {weight, sourceTokens};
    return weight
}

func _reduceTokens(source, target, sourceWords, targetWords string) Tokens {

    sourceTokens := findSourceTokens(source, targetWords);
    targetTokens := findSourceTokens(target, sourceWords);

    return Tokens{Source: sourceTokens, Target: targetTokens};
}

type TokenMap map[string]bool

/**
 * Repeatedly finds tokens and merges them until the merges match ie. the tokens are unchanged.
 * Need a way to remove noisy tokens. When tokens are found, but later they do not persist after reducing, 
 * then we know those are noisy tokens. Save the tokens that are found eg. the TokenMatch object.
 * Could continuously save the previous TokenMatch objects, then take the set difference of the previous and 
 * remove the resulting indexes from the string. Then, rerun until there's no more set difference.
 * @param {Source string} source 
 * @param {Target string} target 
 */
func _triWayTokenMerge(source, target string) Tokens {
    // logger.log();
    // logger.log("_triWayTokenMerge");
    // logger.log("source = ", source, ", target = ", target);
    prevSourceTokens := make(TokenMap)
    tokens := _reduceTokens(source, target, source, target)
    sourceJoined := tokens.Source.JoinMatch()
    _,inMap := prevSourceTokens[sourceJoined]

    for !inMap {
        prevSourceTokens[sourceJoined] = true
        tokens = _reduceTokens(source, target, sourceJoined, tokens.Target.JoinMatch())
        sourceJoined = tokens.Source.JoinMatch()
        _,inMap = prevSourceTokens[sourceJoined]
    }
    
    return tokens
}
