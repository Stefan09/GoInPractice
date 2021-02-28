/**
！！！代码正确性未验证

思考:
跳表本质是基于传统链表的改进
传统链表的缺点是查找定位只能顺序遍历，而跳表的改进思路也十分朴素，能否跳跃遍历
跳跃有自己的章法，一味地均匀跳跃并不是较优方式，故以空间换时间记录跳跃章法 ==> 创建多级上层索引

特点：
1. redis中应用跳表实现zset，score作为索引排序的依据，顺序是节点插入过程中保证的
2. 跳表的查询效率提升至O(logN)，且实现复杂性低于红黑树等结构
3. 检索路线呈阶梯下降状
 */

package linklist_practice

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// define node
type Node struct {
	Value    int     // 节点值
	MaxLevel int     // 节点最高层级
	Next     []*Node // 节点各层级后继节点
}

// define list
type SkipList struct {
	Head       *Node // 头结点
	LevelLimit int   // 层数上限
	Len        int   // 节点个数
}

func NewSkipList(skipListMaxIndexLevel int) *SkipList {
	head := &Node{
		Value:    math.MinInt,
		MaxLevel: skipListMaxIndexLevel,
		Next:     make([]*Node, skipListMaxIndexLevel),
	}

	return &SkipList{
		Head:       head,
		LevelLimit: skipListMaxIndexLevel,
	}
}

// find
func (s *SkipList) Find(value int) {
	iter := s.Head
	for i := s.LevelLimit - 1; i >= 0; i-- {
		for iter.Next[i] != nil && value > iter.Next[i].Value {
			iter = iter.Next[i]
		}
	}
	if iter.Next[0] != nil && value == iter.Next[0].Value {
		fmt.Println("find node:", iter.Next[0].Value)
	} else {
		fmt.Println("not found value")
	}

	return
}

// insert
func (s *SkipList) Insert(value int) {
	// 随机层数
	l := s.RandomLevel()

	// 定位插入位置 逐层操作插入
	iter := s.Head
	node := new(Node)
	node.Value = value
	node.MaxLevel = l
	node.Next = make([]*Node, l)
	for i := l - 1; i >= 0; i-- {
		for iter.Next[i] != nil && value > iter.Next[i].Value {
			iter = iter.Next[i]
		}
		node.Next[i] = iter.Next[i]
		iter.Next[i] = node
	}
	s.Len++

	return
}

// delete
func (s *SkipList) Delete(value int) {
	// 定位所有疑似删除位置
	pos := make([]*Node, s.LevelLimit)

	iter := s.Head
	for i := s.LevelLimit - 1; i >= 0; i-- {
		for iter.Next[i] != nil && value > iter.Next[i].Value {
			iter = iter.Next[i]
		}
		pos[i] = iter
	}

	if iter.Next[0].Value == value {
		for j := iter.Next[0].MaxLevel; j >= 0; j-- {
			pos[j].Next[j] = iter.Next[j]
		}
	}
	s.Len--

	return
}

// random 结构变化时选定变更层，自上而下变更
func (s *SkipList) RandomLevel() int {
	level := 1
	rand.Seed(time.Now().UnixNano())
	for {
		if level == s.LevelLimit || rand.Intn(2) == 1 {
			break
		}
		level++
	}

	return level
}
