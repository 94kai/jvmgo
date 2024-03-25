package rtda

// 表示局部变量表的一个元素（可以是int也可以是引用，jvm中规定，一个元素存int或者引用，两个元素存long/double）
type Slot struct {
	num int32
	ref *Object
}
