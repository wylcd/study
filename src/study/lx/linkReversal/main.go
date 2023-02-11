package main

import "fmt"

// 链表反转
type linkNode struct {
	Sort int64
	ln   *linkNode
}

var lb = &linkNode{
	1, &linkNode{
		2, &linkNode{
			3, &linkNode{
				4, &linkNode{
					5, nil,
				},
			},
		},
	},
}

// 迭代
func linkList1() *linkNode {
	a := &linkNode{}
	for lb != nil {
		next := lb.ln // 把下一个节点缓存起来。保存此时的下一个节点。（1保存2）
		lb.ln = a     // 当前节点的下一个节点为上一个节点。将此时的下一个节点（1中的2）反转存储为上一个节点（1中的0）。
		a = lb        // 将倒叙的节点存储在lb中。将此时的节点存储起来留着下一个反转节点使用（将2的下一个节点修改为1）。
		lb = next     // 重置cur。使用下一个节点转换（使用2）。
	}
	return a
}

// 递归
func linkList2(head *linkNode) *linkNode {
	// 此时要从1获取到5。
	if head == nil || head.ln == nil {
		return head
	}
	newHead := linkList2(head.ln) // 5
	head.ln.ln = head             // 5->4
	head.ln = nil                 // 4->nil
	return newHead
}

func main() {
	// 反转链表。
	a := linkList1()
	//a := linkList2(lb)
	for a != nil {
		fmt.Print(a.Sort, "->")
		a = a.ln
	}
}
