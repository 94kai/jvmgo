package stack

import "ch05/instructions/base"
import "ch05/rtda"

// Swap the top two operand stack values
// 交换操作数栈顶顶两个元素的值
type SWAP struct{ base.NoOperandsInstruction }

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
