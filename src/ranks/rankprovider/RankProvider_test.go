package rankprovider

import (
	"testing"
	"fmt"
)

func TestRankProvider(t *testing.T) {
	res := Get("abcd", "abc")
	fmt.Println(res)
	res = Get("BMW 528i V6 65,000 miles No accidents, Excellent Condition, Luxurious! sedan",
"X5 xDrive35i Sport Activity Sport Utility 4D")
	fmt.Println(res)
}