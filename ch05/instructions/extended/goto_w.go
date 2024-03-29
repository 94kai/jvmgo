package extended

import "ch05/instructions/base"
import "ch05/rtda"

// Branch always (wide index)
// 对goto指令offset的扩展，用4字节表示offset,goto使用2字节来记录offset
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}
