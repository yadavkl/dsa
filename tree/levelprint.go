package tree

type Que struct {
	Elements []*Node
	Len      int
}

func NewQue() *Que {
	return &Que{
		Elements: make([]*Node, 0),
		Len:      0,
	}
}

func (q *Que) Enq(node *Node) {
	q.Elements = append(q.Elements, node)
	q.Len++
}
func (q *Que) Deq() *Node {
	node := q.Elements[0]
	q.Elements = q.Elements[1:]
	q.Len--
	return node
}

func (q *Que) Size() int {
	return q.Len
}

func LevelPrint(root *Node, nodes *[]int) {
	que := NewQue()
	que.Enq(root)

	for que.Size() > 0 {
		node := que.Deq()
		*(nodes) = append(*(nodes), node.Val)
		if node.left != nil {
			que.Enq(node.left)
		}
		if node.right != nil {
			que.Enq(node.right)
		}
	}
}
