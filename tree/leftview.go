package tree

/*
	Level terversal for left view of the tree
	1. Insert levele nodes in tree
	2. print 1st node of level and insert childs of the nodes
*/
func LeftView(root *Node, nodes *[]int) {
	que := NewQue()
	que.Enq(root)
	for que.Size() > 0 {
		size := que.Size()
		for i := 0; i < size; i++ {
			node := que.Deq()
			if i == 0 {
				*(nodes) = append(*(nodes), node.Val)
			}
			if node.left != nil {
				que.Enq(node.left)
			}
			if node.right != nil {
				que.Enq(node.right)
			}
		}

	}
}
