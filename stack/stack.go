package stack

type Stack struct {
	data []interface{}
	Top  int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]interface{}, 0),
		Top:  0,
	}
}

func (st *Stack) Push(v interface{}) {
	st.data = append(st.data, v)
	st.Top++
}

func (st *Stack) Pop() interface{} {
	val := st.data[st.Top-1]
	st.data = st.data[:st.Top-1]
	st.Top--
	return val
}

func (st *Stack) IsEmpty() bool {
	if st.Top <= 0 {
		return true
	}
	return false
}
