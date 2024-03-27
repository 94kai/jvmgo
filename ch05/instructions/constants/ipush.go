package constants

import "ch05/instructions/base"
import "ch05/rtda"

// 从操作数中提取出一个byte或short，扩展成int，push到操作数栈中
type BIPUSH struct{ val int8 }  // Push byte
type SIPUSH struct{ val int16 } // Push short

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
