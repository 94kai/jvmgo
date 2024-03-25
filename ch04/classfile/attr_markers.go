package classfile

type DeprecatedAttribute struct{ MarkerAttribute }
type SyntheticAttribute struct{ MarkerAttribute }
type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// Deprecated和Synthetic都属于标记，数据长度是0 
}
