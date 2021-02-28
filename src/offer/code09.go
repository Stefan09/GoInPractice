package offer

import "container/list"

/**
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
*/

type CQueue struct {
	stack1 *list.List
	stack2 *list.List
}

func Constructor() CQueue {
	inst := CQueue{}
	inst.stack1 = list.New()
	inst.stack2 = list.New()
	return inst
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.PushFront(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack1.Len()+this.stack2.Len() == 0 {
		return -1
	}
	if this.stack2.Len() != 0 {
		res := this.stack2.Remove(this.stack2.Front()).(int)
		return res
	}
	for this.stack1.Len() != 0 {
		iter := this.stack1.Remove(this.stack1.Front()).(int)
		this.stack2.PushFront(iter)
	}
	res := this.stack2.Remove(this.stack2.Front()).(int)
	return res
}
