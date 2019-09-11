package rbTree

import "fmt"

type Node struct {
	Black bool
	Left, Right, Parent *Node
	Data float64
}

func NewRoot(data float64) *Node {
	return &Node{
		Black:  true,
		Data:   data,
	}
}

func (n *Node) isRoot() bool {
	if n == nil {
		// edge leaf node
		return false
	}
	return n.Parent == nil
}

func (n *Node) Insert(data float64) (err error) {
	if data == n.Data {
		return fmt.Errorf("duplicate data: %v", data)
	}

	newNode := &Node{
		Black: false,
		Data: data,
	}

	parent := n
	var child *Node
	for {
		if data < parent.Data {
			child = parent.Left
			if child == nil {
				parent.Left = newNode
				newNode.Parent = parent.Left
			}
		} else if data > parent.Data {
			child = parent.Right
			if child == nil {
				parent.Right = newNode
				newNode.Parent = parent.Right
			}
		} else {
			return fmt.Errorf("duplicate data: %v", data)
		}

		if child != nil {
			parent = child
		} else {
			newNode.insertBalance()
			return
		}
	}
}

func (n *Node) insertBalance() {
	// n一定不是根结点
	noGrandParent := n.Parent.isRoot()
	grandParent := n.Parent.Parent
	switch {
	// 父节点是root，肯定不是下面三种需要修复平衡的情况
	case noGrandParent:
		return

	// 插入修复情况1：如果当前结点的父结点是红色且祖父结点的另一个子结点（叔叔结点）是红色
	case !(grandParent.Left.Black || grandParent.Right.Black):
		n.insertBalance1()

	// 插入修复情况2：当前节点的父节点是红色,叔叔节点是黑色，当前节点是其父节点的右子
	case (grandParent.Left.Black || grandParent.Right.Black) && (n.Parent.Right == n):
		n.insertBalance2()

	// 插入修复情况3：当前节点的父节点是红色,叔叔节点是黑色，当前节点是其父节点的左子
	case (grandParent.Left.Black || grandParent.Right.Black) && (n.Parent.Left == n):
		n.insertBalance3()
	}
}

// 插入修复情况1：如果当前结点的父结点是红色且祖父结点的另一个子结点（叔叔结点）是红色
func (n *Node) insertBalance1() {
	grand := n.Parent.Parent
	grand.Black = false
	grand.Left.Black = true
	grand.Right.Black = true
	grand.insertBalance()
}

// 插入修复情况2：当前节点的父节点是红色,叔叔节点是黑色，当前节点是其父节点的右子
func (n *Node) insertBalance2() {
	parent := n.Parent
	grand := parent.Parent

	if grand.Left == n.Parent {
		grand.Left = n
	} else {
		grand.Right = n
	}

	parent.Right = n.Left
	n.Left.Parent = parent.Right
	n.Left = parent
	parent.Parent = n

	parent.insertBalance()
}

// 插入修复情况3：当前节点的父节点是红色,叔叔节点是黑色，当前节点是其父节点的左子
func (n *Node) insertBalance3() {
	parent := n.Parent
	grand := parent.Parent

	parent.Black = true
	grand.Black = false

	grand.Left = parent.Right
	parent.Right.Parent = grand
	parent.Right = grand
	grand.Parent = parent
	parent.Parent = nil
}

func (n *Node) Delete(data float64) (err error) {
	return
}

func (n *Node) Find(data float64) (err error) {
	return
}
