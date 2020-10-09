package load_balance

import (
	"fmt"
	"sort"
	"testing"
)

func TestWeightRoundRobinBalance(t *testing.T) {
	idx := sort.Search(5, func(i int) bool {
		return false
	})
	t.Log(idx)
	return
	r := &WeightRoundRobinBalance{rss: make([]*WeightNode, 0, 10)}

	r.Add("1", 5)
	r.Add("2", 1)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Next())
	}
}
