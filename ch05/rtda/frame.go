package rtda

type Frame struct {
	lower        *Frame        // 上一个栈帧
	localVars    LocalVars     // 局部变量表指针
	operandStack *OperandStack // 操作数栈指针
	Thread       *Thread
	nextPC       int
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		Thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
