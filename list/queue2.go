package list

type (
	//Queue2 队列
	Queue2 struct {
		top    *node
		rear   *node
		length int
	}
	//双向链表节点
	node struct {
		pre   *node
		next  *node
		value interface{}
	}
)

// Create a new Queue2
func New() *Queue2 {
	return &Queue2{nil, nil, 0}
}

//获取队列长度
func (q *Queue2) Len() int {
	return q.length
}

//返回true队列不为空
func (q *Queue2) Any() bool {
	return q.length > 0
}

//返回队列顶端元素
func (q *Queue2) Peek() interface{} {
	if q.top == nil {
		return nil
	}
	return q.top.value
}

//入队操作
func (q *Queue2) Push(v interface{}) {
	n := &node{nil, nil, v}
	if q.length == 0 {
		q.top = n
		q.rear = q.top
	} else {
		n.pre = q.rear
		q.rear.next = n
		q.rear = n
	}
	q.length++
}

//出队操作
func (q *Queue2) Pop() interface{} {
	if q.length == 0 {
		return nil
	}
	n := q.top
	if q.top.next == nil {
		q.top = nil
	} else {
		q.top = q.top.next
		q.top.pre.next = nil
		q.top.pre = nil
	}
	q.length--
	return n.value
}
