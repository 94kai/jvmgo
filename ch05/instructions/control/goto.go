package control

import "ch05/instructions/base"
import "ch05/rtda"

// 继承自BranchInstruction,所有有一个操作数用来记录offset，直接跳转到操作数对应的offset
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
