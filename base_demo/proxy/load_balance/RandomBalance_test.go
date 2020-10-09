package load_balance

import (
	"fmt"
	"testing"
)

func TestRandomBalance_Add(t *testing.T) {
	//rb := &RandomBalance{}
	rb := &RoundRobinBalance{}
	rb.Add("127.0.0.1:2023")
	rb.Add("127.0.0.1:2024")
	rb.Add("127.0.0.1:2025")
	rb.Add("127.0.0.1:2026")
	rb.Add("127.0.0.1:2027")
	rb.Add("127.0.0.1:2028")
	rb.Add("127.0.0.1:2029")
	rb.Add("127.0.0.1:2030")

	for i := 0; i < 100; i++ {
		fmt.Println(rb.Next())
	}
}
