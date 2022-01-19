/** 位图bitmap
目的压缩存储空间以单字节代表元素；快速定位查找，多应用于大数据量去重等场景
golang中无原生bitmap，故需要实现基本操作：元素定位（位于哪个字节、偏移量多少）进而通过位运算实现add/delete
tips：golang中很多数据结构都没有像java一样封装实现，设计者可能是觉得基础构建好之后其余的结构都可以实现，就像bitmap可以用[]byte

定位思路借鉴操作系统中内存分页式管理（做除法&取余）

思考：有符号数的处理
*/
package search_sort_practice

type BitMap []byte

func NewBitMap(l int64) BitMap {
	return make(BitMap, l)
}

// 定位
func (b *BitMap) Pos(v uint64) (index, offset uint64) {
	index = v / 8
	offset = v & 0x07
	return
}

// 添加
func (b *BitMap) Add(v uint64) {
	index, offset := b.Pos(v)
	(*b)[index] |= 1 << offset
}

// 删除
func (b *BitMap) Del(v uint64) {
	index, offset := b.Pos(v)
	(*b)[index] &= ^(1 << offset)
}
