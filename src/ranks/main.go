package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"serverless-craigs/src/ranks/rankprovider"
	"sync"
)

func inSequenceCount(s Source, t Target, wg *sync.WaitGroup, targetChannel chan<- RankedTarget) {

	targetChannel <- RankedTarget{Target: t, Rank: rankprovider.Get(string(s), t.Text)}
	wg.Done()
}

func MatchSourceAndTargets(sources []Source, targets []Target) RankedTarget {
	var wg sync.WaitGroup
	combos := len(targets) * len(sources)
	targetChannel := make(chan RankedTarget, combos)

	for _, t := range targets {
		for _, s := range sources {
			wg.Add(1)
			go inSequenceCount(s, t, &wg, targetChannel)
		}
	}
	wg.Wait()
	close(targetChannel)

	rankedMap := combineRankedTargets(targetChannel)
	rankedTargets := consolodateRankedTargets(rankedMap)
	return findHighestRank(rankedTargets)
}

// Map of RankedTarget to their combined ranks
func combineRankedTargets(targetChannel <-chan RankedTarget) TargetRankMap {

	rankedMap := make(TargetRankMap)
	countMap := make(map[Target]int)
	for receivedTarget := range targetChannel {
		rankedMap[receivedTarget.Target] += receivedTarget.Rank
		countMap[receivedTarget.Target]++
	}

	for k, v := range countMap {
		rankedMap[k] /= float32(v)
	}

	return rankedMap
}

// Consolidate RankedTargets into unique RankedTargets after combining ranks
func consolodateRankedTargets(rankedMap TargetRankMap) []RankedTarget {

	var rankedTargets []RankedTarget
	for t := range rankedMap {
		rankedTargets = append(rankedTargets, RankedTarget{Target: t, Rank: rankedMap[t]})
	}
	return rankedTargets
}

// Find highest rank
func findHighestRank(rankedTargets []RankedTarget) RankedTarget {

	var high RankedTarget
	for _, rankedTarget := range rankedTargets {
		if rankedTarget.Rank >= high.Rank {
			high = rankedTarget
		}
	}
	return high
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Target is the kbb related info for the Text we're matching, and the associated meta Data
type Target struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

// RankedTarget wraps a Target with a ranking
type RankedTarget struct {
	Target Target  `json:"target"`
	Rank   float32 `json:"rank"`
}

// TargetRankMap for combining ranks per target
type TargetRankMap map[Target]float32

// Source - The craigslist related strings eg. title and description
type Source string

// RanksRequest is the input type for Handler
type RanksRequest struct {
	Sources []Source `json:"sources"`
	Targets []Target `json:"targets"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {

	ranksRequest := RanksRequest{}
	json.Unmarshal([]byte(request.Body), &ranksRequest)
	rankedTarget := MatchSourceAndTargets(ranksRequest.Sources, ranksRequest.Targets)

	fmt.Println(ranksRequest)
	encoded, _ := json.Marshal(&rankedTarget)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(encoded),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "ranks-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
	// r := events.APIGatewayProxyRequest{}
	// fmt.Println(Handler(r))
	// fmt.Println("Done");
}
