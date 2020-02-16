package red_black_tree

import "fmt"

//参考https://www.cnblogs.com/ailumiyana/p/10963658.html

const (
	RED   = true
	BLACK = false
)

type Item interface {
	Less(then Item) bool
}

type Int int

func (x Int) Less(than Item) bool {
	fmt.Println(x, " ", than.(Int))
	return x < than.(Int)
}

type Uint32 uint32

func (x Uint32) Less(than Item) bool {
	fmt.Println(x, " ", than.(Uint32))
	return x < than.(Uint32)
}

type String string

func (x String) Less(than Item) bool {
	fmt.Println(x, " ", than.(String))
	return x < than.(String)
}

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	color  bool
	Item
}

type Rbtree struct {
	NIL   *Node
	root  *Node
	count uint64
}

func New() *Rbtree {
	node := &Node{nil, nil, nil, BLACK, nil}
	return &Rbtree{
		NIL:   node,
		root:  node,
		count: 0,
	}
}

func less(x, y Item) bool {
	return x.Less(y)
}

//          |                                  |
//          X                                  Y
//         / \         left rotate            / \
//        α  Y       ------------->         X   γ
//           / \                            / \
//          β  γ                            α  β
func (rbt *Rbtree) LeftRotate(no *Node) {
	if no.Right == rbt.NIL {
		return
	}
	rchild := no.Right
	no.Right = rchild.Left
	if rchild.Left != rbt.NIL {
		rchild.Left.Parent = no
	}
	rchild.Parent = no.Parent
	if no.Parent == rbt.NIL {
		rbt.root = rchild
	} else if no == no.Parent.Left {
		no.Parent.Left = rchild
	} else {
		no.Parent.Right = rchild
	}
	rchild.Left = no
	no.Parent = rchild
}

//          |                                  |
//          X                                  Y
//         / \         right rotate           / \
//        Y   γ      ------------->         α  X
//       / \                                    / \
//      α  β                                    β  γ
func (rbt *Rbtree) RightRotate(no *Node) {
	if no.Left == rbt.NIL {
		return
	}
	lchild := no.Left
	no.Left = lchild.Right
	if lchild.Right != rbt.NIL {
		lchild.Right.Parent = no
	}
	lchild.Parent = no.Parent
	if no.Parent == rbt.NIL {
		rbt.root = lchild
	} else if no == no.Parent.Left {
		no.Parent.Left = lchild
	} else {
		no.Parent.Right = lchild
	}
	lchild.Right = no
	no.Parent = lchild
}

func (rbt *Rbtree) Insert(no *Node) {
	x := rbt.root
	var y = rbt.NIL
	for x != rbt.NIL {
		y = x
		if less(no.Item, x.Item) {
			x = x.Left
		} else if less(x.Item, no.Item) {
			x = x.Right
		} else {
			fmt.Println("that node already exist")
		}
	}
	no.Parent = y
	if y == rbt.NIL {
		rbt.root = no
	} else if less(no.Item, y.Item) {
		y.Left = no
	} else {
		y.Right = no
	}
	rbt.count++
	rbt.insertFixup(no)
}

func (rbt *Rbtree) insertFixup(no *Node) {
	for no.Parent.color == RED {
		if no.Parent == no.Parent.Parent.Left {
			y := no.Parent.Parent.Right
			if y.color == RED {
				//
				// 情形 4
				fmt.Println("TRACE Do Case 4 :", no.Item)
				no.Parent.color = BLACK
				y.color = BLACK
				no.Parent.Parent.color = RED
				no = no.Parent.Parent //循环向上自平衡.
			} else {
				if no == no.Parent.Right {
					//
					// 情形 5 : 反向情形
					// 直接左旋转 , 然后进行情形3(变色->右旋)
					fmt.Println("TRACE Do Case 5 :", no.Item)
					if no == no.Parent.Right {
						no = no.Parent
						rbt.LeftRotate(no)
					}
				}
				fmt.Println("TRACE Do Case 6 :", no.Item)
				no.Parent.color = BLACK
				no.Parent.Parent.color = RED
				rbt.RightRotate(no.Parent.Parent)
			}
		} else { //为父父节点右孩子情形，和左孩子一样，改下转向而已.
			y := no.Parent.Parent.Left
			if y.color == RED {
				no.Parent.color = BLACK
				y.color = BLACK
				no.Parent.Parent.color = RED
				no = no.Parent.Parent
			} else {
				if no == no.Parent.Left {
					no = no.Parent
					rbt.RightRotate(no)
				}
				no.Parent.color = BLACK
				no.Parent.Parent.color = RED
				rbt.LeftRotate(no.Parent.Parent)
			}
		}
	}
	rbt.root.color = BLACK
}
