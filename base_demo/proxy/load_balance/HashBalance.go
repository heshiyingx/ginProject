package load_balance

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type Hash func(data []byte) uint32
type Uint32Slice []uint32

func (s Uint32Slice) Len() int {
	return len(s)
}

func (s Uint32Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Uint32Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ConsistentHashBalance struct {
	mux      sync.RWMutex
	hash     Hash
	replicas int               // 复制因子
	keys     Uint32Slice       // 已排序的节点hash切片
	hashMap  map[uint32]string // 节点hash和key的map,key是hash值，值是节点的key

	// 观察主体
	//conf LoadBalanceConf
}

func NewConsistentHashBalance(replicas int, fn Hash) *ConsistentHashBalance {
	if fn == nil {
		fn = crc32.ChecksumIEEE
	}
	b := &ConsistentHashBalance{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[uint32]string),
	}
	return b
}

func (c *ConsistentHashBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}

	addr := params[0]
	c.mux.Lock()
	defer c.mux.Unlock()

	for i := 0; i < c.replicas; i++ {
		hash := c.hash([]byte(strconv.Itoa(i) + addr))
		c.keys = append(c.keys, hash)
		c.hashMap[hash] = addr
	}
	sort.Sort(c.keys)
	return nil
}

func (c *ConsistentHashBalance) Get(key string) (string, error) {
	if c.isEmpty() {
		return "", errors.New("node is empty")
	}

	hash := c.hash([]byte(key))
	// 通过二分查找获取最优节点，第一个服务器hash值大于 数据hash值的就是最优 服务器节点
	idx := sort.Search(len(c.keys), func(i int) bool { return c.keys[i] >= hash })
	if idx == len(c.keys) {
		idx = 0
	}
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.hashMap[c.keys[idx]], nil
}

func (c *ConsistentHashBalance) isEmpty() bool {
	return false
}
