package load_balance

import "math/rand"

type RandomBalance struct {
	curIndex int
	rss      []string

	// 观察者模式
	//conf LoadBalanceConf
}

func (r *RandomBalance) Add(params ...string) error {
	addr := params[0]
	r.rss = append(r.rss, addr)
	return nil
}

func (r *RandomBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}
	r.curIndex = rand.Intn(len(r.rss))
	return r.rss[r.curIndex]
}
