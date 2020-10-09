package load_balance

// WeightNode 权重模型
type WeightNode struct {
	addr            string
	weight          int // 权重
	currentWeight   int // 节点当前权重
	effectiveWeight int // 有效权重
}
type WeightRoundRobinBalance struct {
	curIndex int
	rss      []*WeightNode
	rsw      []int
}

/**
实现思路(4,3,2)
1. currentWeight = currentWeight+effectiveWeight
2. 选中最大的currentWeight节点为选中节点
3.currentWeight = currentWeight-totalWeight
*/
func (r *WeightRoundRobinBalance) Add(addr string, weight int) {
	r.rss = append(r.rss, &WeightNode{
		addr:            addr,
		weight:          weight,
		currentWeight:   0,
		effectiveWeight: weight,
	})
}
func (r *WeightRoundRobinBalance) Next() string {
	total := 0
	var best *WeightNode

	for i := 0; i < len(r.rss); i++ {
		node := r.rss[i]
		total += node.effectiveWeight
		node.currentWeight += node.effectiveWeight

		//	有效权重默认与权重相同，通讯异常-1，通讯成功+1,直到恢复到weight大小
		if node.effectiveWeight < node.weight {
			node.effectiveWeight++
		}

		if best == nil || node.currentWeight > best.currentWeight {
			best = node
		}
	}

	if best == nil {
		return ""
	}

	best.currentWeight -= total
	return best.addr
}
