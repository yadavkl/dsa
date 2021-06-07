package tree

func Serialise(root *Node) []int {
	arr := make([]int, 0)
	InOrderSerialise(root, &arr)
	return arr
}

func InOrderSerialise(root *Node, arr *[]int) {
	if root == nil {
		*arr = append(*arr, -1)
		return
	}
	*arr = append(*arr, root.Val)
	InOrderSerialise(root.left, arr)
	InOrderSerialise(root.right, arr)
}

func Deserialise(arr []int) *Node {
	index := 0
	return InOrderDeserialise(arr, &index)
}

func InOrderDeserialise(arr []int, index *int) *Node {
	if *index == len(arr) {
		return nil
	}
	val := arr[*index]
	(*index)++
	if val == -1 {
		return nil
	}
	root := NewNode()
	root.Val = val
	root.left = InOrderDeserialise(arr, index)
	root.right = InOrderDeserialise(arr, index)
	return root
}
