package list

import "fmt"

type List struct {
	len       int64
	firstNode *node
}

func NewList() *List {
	l := &List{
		len:       0,
		firstNode: nil,
	}
	return l
}

type node struct {
	data     int64
	nextNode *node
}

func (l *List) Add(data int64) (index int64) {
	if l.firstNode == nil {
		n := &node{}
		n.data = data
		l.firstNode = n
		l.len++
		return l.len - 1
	}
	nn := l.firstNode
	for {
		if nn.nextNode == nil {
			break
		}
		nn = nn.nextNode
	}
	n := &node{}
	n.data = data
	nn.nextNode = n
	l.len++
	return l.len - 1
}

func (l *List) Delete(index int64) (ok bool) {
	if l.firstNode == nil {
		return false
	}
	if index >= l.len {
		return false
	}
	nn := l.firstNode
	if index == 0 {
		l.firstNode = nn.nextNode
		l.len--
		return true
	}
	for i := int64(0); i < index-1; i++ {
		nn = nn.nextNode
	}
	delNode := nn.nextNode
	nn.nextNode = delNode.nextNode
	delNode.nextNode = nil
	l.len--
	return true
}

func (l *List) Print() {
	fmt.Println("List length", l.len)
	if l.firstNode == nil {
		return
	}
	nn := l.firstNode

	for {
		if nn.nextNode == nil {
			fmt.Println(nn.data)
			break
		}
		fmt.Println(nn.data)
		nn = nn.nextNode
	}
}
func (l *List) SortIncrease() {
	for j := int64(0); j < l.len; j++ {
		nn := l.firstNode
		for {
			if nn.nextNode == nil {
				break
			}
			value := nn.data
			prevData := &nn.data
			nn = nn.nextNode
			if nn.data < value {
				a := nn.data
				nn.data = value
				*prevData = a
			}
		}
	}
}

func (l *List) SortDecrease() {
	for j := int64(0); j < l.len; j++ {
		nn := l.firstNode
		for {
			if nn.nextNode == nil {
				break
			}
			value := nn.data
			prevData := &nn.data
			nn = nn.nextNode
			if nn.data > value {
				a := nn.data
				nn.data = value
				*prevData = a
			}
		}
	}
}

func (l *List) SortIncreaseNode() {
	var (
		a *node
		b *node
		c *node
	)

	for j := int64(0); j < l.len; j++ {
		nn := l.firstNode
		a = nn
		b = nn
		c = nn

		for {
			if c.nextNode == nil {
				break
			}
			b = nn.nextNode
			if b.nextNode == nil {
				if a.data > b.data {
					l.firstNode = b
					a.nextNode = nil
					b.nextNode = a
				}
				break
			}
			c = b.nextNode
			if (a.data > b.data) && (l.firstNode == a) {
				l.firstNode = b
				a.nextNode = b.nextNode
				b.nextNode = a
			}
		}
	}
}
