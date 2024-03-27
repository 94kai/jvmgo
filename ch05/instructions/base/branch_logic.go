package base

import "ch05/rtda"
// 跳转逻辑，字节码执行到哪里是通过PC寄存器记录的，所以跳转其实就是修改PC寄存器的值
func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread.PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
