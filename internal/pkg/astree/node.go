package astree

type Node struct {
	_type  string
	value  int
	parent *Node
	next   []*Node
}

func NewNode(value int, parent *Node, _type string) *Node {
	return &Node{
		value:  value,
		parent: parent,
		_type:  _type,
	}
}

func (n *Node) GetValue() int {
	return n.value
}

func (n *Node) GetNext() []*Node {
	return n.next
}

func (n *Node) GetParent() *Node {
	return n.parent
}

func (n *Node) GetType() string {
	return n._type
}

func (n *Node) SetValue(value int) {
	n.value = value
}

func (n *Node) SetNext(next []*Node) {
	n.next = next
}

func (n *Node) SetParent(parent *Node) {
	n.parent = parent
}

func (n *Node) GetSiblings(parent *Node) []*Node {
	var nArr []*Node

	for _, item := range n.parent.next {
		if item != n {
			nArr = append(nArr, item)
		}
	}

	return nArr
}
