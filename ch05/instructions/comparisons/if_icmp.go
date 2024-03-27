package comparisons

import "ch05/instructions/base"
import "ch05/rtda"

// Branch if int comparison succeeds
// 操作数栈顶pop出两个元素进行比较，相同就跳转
type IF_ICMPEQ struct{ base.BranchInstruction }

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 == val2 {
		base.Branch(frame, self.Offset)
	}
}

// 操作数栈顶pop出两个元素进行比较，不同就跳转
type IF_ICMPNE struct{ base.BranchInstruction }

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}

// 操作数栈顶pop出两个元素进行比较，第一个小于第二个就跳转
type IF_ICMPLT struct{ base.BranchInstruction }

func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		base.Branch(frame, self.Offset)
	}
}

// ...
type IF_ICMPLE struct{ base.BranchInstruction }

func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		base.Branch(frame, self.Offset)
	}
}

// ...
type IF_ICMPGT struct{ base.BranchInstruction }

func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		base.Branch(frame, self.Offset)
	}
}

// ...
type IF_ICMPGE struct{ base.BranchInstruction }

func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 >= val2 {
		base.Branch(frame, self.Offset)
	}
}

func _icmpPop(frame *rtda.Frame) (val1, val2 int32) {
	stack := frame.OperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return
}
