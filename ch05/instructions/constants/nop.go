package constants

import "ch05/instructions/base"
import "ch05/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// 什么也不用做
}
