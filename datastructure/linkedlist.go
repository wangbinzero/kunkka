package datastructure

import "sync"

type Value interface {
}

// 节点
type LinkedNode struct {
	Data Value       //当前节点数据
	Prev *LinkedNode //前节点
	Next *LinkedNode //后节点
}

type LinkedList struct {
	mutex *sync.RWMutex //读写锁
	Size  uint          //链表长度
	Head  *LinkedNode   //头
	Tail  *LinkedNode   //尾
}

//初始化链表
func (this *LinkedList) Init() {
	this.mutex = new(sync.RWMutex)
	this.Size = 0
	this.Head = nil
	this.Tail = nil
}

//节点查询
func (this *LinkedList) Get(index uint) *LinkedNode {

	//如果当前链表长度为0 或者 查询索引大于当前链表长度则返回空
	if this.Size == 0 || index > this.Size-1 {
		return nil
	}

	//如果查询索引为0 则返回链表头
	if index == 0 {
		return this.Head
	}

	//因为所查询的节点不是头部，所以从索引为1开始
	node := this.Head
	var i uint
	for i = 1; i <= index; i++ {
		node = node.Next
	}

	return node
}

//向双向链表尾部追加节点
func (this *LinkedList) Add(node *LinkedNode) bool {
	if node == nil {
		return false
	}

	this.mutex.Lock()
	defer this.mutex.Unlock()

	//如果链表长度为0，则将节点赋值给Head Tail
	if this.Size == 0 {
		this.Head = node
		this.Tail = node
		node.Prev = nil
		node.Next = nil
	} else {
		node.Prev = this.Tail
		node.Next = nil
		this.Tail.Next = node
		this.Tail = node
	}
	this.Size++
	return true
}

//向指定位置插入节点
func (this *LinkedList) Insert(index uint, node *LinkedNode) bool {

}
