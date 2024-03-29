package insequence

import (
	"serverless-craigs/src/logger"
	"regexp"
	"sort"
	"strings"
)

// ResMax - The max result of insequence
type ResMax struct {
	Count  int
	Match  string
	Start  int
	End    int
	Weight float32
}

func (r ResMax) fill(n int) []ResMax {
	var s []ResMax
	for i := 1; i <= n; i++ {
		s = append(s, r)
	}
	return s
}

/**
*	b - Target string input
* match - Matching characters from source input
* start - start index of target match
* end - end index of target match
* return float - Distance from source and target match
**/
func calcDistance(b, match string, start, end int) float32 {
	logger := logger.Logger{false}
	var distance float32
	logger.Log("b = ", b, "start = ", start, "end = ", end)
	targetMatch := b[start:end+1]
	

	for i := 0; i < len(match); i++ {
		targetMatch = strings.Replace(targetMatch, string(match[i]), "", 1)
	}

	for i := 0; i < len(targetMatch); i++ {
		matchChar := targetMatch[i]
		if string(matchChar) == " " {
			distance += .5
		} else {
			distance++
		}
	}

	return distance
}

func calcWeight(count, aLen, end, start int, b, match string) float32 {
	logger := logger.Logger{false}
	// return float32(count*2) / float32(bLen+(end-start)+1)
	// targetMatchLength := (end-start)+1
	distance := calcDistance(b, match, start, end)
	
	// distance := float32(0)
	
	// if targetMatchLength > count {
	// 	distance = 1 - (float32(count) / float32(targetMatchLength))
	// }

	weight := (float32(count) / (float32(aLen) + distance))

	logger.Log("aLen = ", aLen, ", end = ", end, ", start = ", start, "weight = ", weight, "count = ", count)
	return weight
}

type byWeight []ResMax

func (a byWeight) Len() int           { return len(a) }
func (a byWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byWeight) Less(i, j int) bool { return a[j].Weight < a[i].Weight }

// Insequence - The dynamic programming algorithm for retrieving the max Weight subsequence that is increasing order aka insequence
// This is also the inverse of levenshtein
func Insequence(a string, b string) ResMax {
	var logger = logger.Logger{false}

	resMax := ResMax{Count: 0, Match: "", Start: 0, End: 0, Weight: 0}
	// Try to use a single space between words
	a = strings.ToLower(a)
	re := regexp.MustCompile(`\s`)
	a = re.ReplaceAllString(a, " ")
	re = regexp.MustCompile(`(\w)-(\w)`)
	a = re.ReplaceAllString(a, "$1 $2")

	b = strings.ToLower(b)
	re = regexp.MustCompile(`\s`)
	b = re.ReplaceAllString(b, " ")
	re = regexp.MustCompile(`(\w)-(\w)`)
	b = re.ReplaceAllString(b, "$1 $2")

	// bLen := len(b)
	aLen := len(a)

	// Create empty edit distance matrix for all possible modifications of
	// substrings of a to substrings of b.
	var distanceMatrix [][]ResMax
	for j := 1; j <= len(b)+1; j++ {
		var row = resMax.fill(len(a) + 1)
		distanceMatrix = append(distanceMatrix, row)
	}

	for j := 1; j <= len(b); j++ {
		for i := 1; i <= len(a); i++ {

			// for _, m := range distanceMatrix {
			// 	logger.Log(m)
			// }

			isMatch := a[i-1] == b[j-1]

			deletion := distanceMatrix[j][i-1] // deletion
			if deletion.Count > 0 {
				// deletion.End = deletion.End + 1
				deletion.End = j - 1
				deletion.Weight = calcWeight(deletion.Count, aLen, deletion.End, deletion.Start, b, deletion.Match)
				logger.Log("deletion = ", deletion)
			}

			insertion := distanceMatrix[j-1][i] // insertion
			if insertion.Count > 0 {
				insertion.End = j - 1
				insertion.Weight = calcWeight(insertion.Count, aLen, insertion.End, insertion.Start, b, insertion.Match)
				logger.Log("insertion = ", insertion)
			}

			substitution := distanceMatrix[j-1][i-1] // substitution
			if isMatch {
				logger.Log("before substitution = ", substitution)
				logger.Log("i = ", i, ", j = ", j)
				if substitution.Count == 0 {
					logger.Log("new start = ", j-1)
					substitution.Start = j - 1
				}
				substitution.End = j - 1
				substitution.Count++
				substitution.Match = substitution.Match + string(a[i-1])
				substitution.Weight = calcWeight(substitution.Count, aLen, substitution.End, substitution.Start, b, substitution.Match)
				logger.Log("after substitution = ", substitution)
			}

			arr := []ResMax{deletion, insertion, substitution}
			sort.Sort(byWeight(arr))

			if arr[0].Weight > resMax.Weight {
				resMax = arr[0]
			}

			distanceMatrix[j][i] = arr[0]
		}
	}

	for _, m := range distanceMatrix {
		logger.Log(m)
	}
	
	return resMax
}

// console.log(newInsequence('abc d', 'aee abc eec'));
// newInsequence('abc def ghi', 'ghi jkl abc x');
// newInsequence('CLS CLS500 MILITARY 0 DOWN NAVY FED', 'CLS Class CLS 500 Coupe 4D');
// newInsequence("Mercedes Benz GL 450 awd suv 7 passenger/ BEST OFFER .", "GLClass");
// levenshteinDistance("Using lcs, found match source =  Mercedes Benz GL 450 awd suv 7 passenger/ BEST OFFER .", "CClass");
