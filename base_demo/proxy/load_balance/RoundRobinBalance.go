package load_balance

import "errors"

type RoundRobinBalance struct {
	curIndex int
	rss      []string

	// 观察者模式
	//conf LoadBalanceConf
}

func (r *RoundRobinBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	r.rss = append(r.rss, addr)
	return nil
}

func (r *RoundRobinBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}

	lens := len(r.rss)
	r.curIndex = (r.curIndex + 1) % lens
	if r.curIndex >= lens {
		r.curIndex = 0
	}

	curAddr := r.rss[r.curIndex]
	return curAddr
}
