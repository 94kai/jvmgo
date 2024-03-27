package stack

import "ch05/instructions/base"
import "ch05/rtda"
// 从操作数栈中弹出变量
type POP struct{ base.NoOperandsInstruction }
type POP2 struct{ base.NoOperandsInstruction }

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}
func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
