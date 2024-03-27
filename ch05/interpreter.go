package main

import (
	"ch05/instructions"
	"ch05/instructions/base"
	"fmt"
)
import "ch05/classfile"
import "ch05/rtda"

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code // 其他代码
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, bytecode)
}
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		// decode
		reader.Reset(bytecode, pc)
		// 读取操作码，顺便pc会++
		opcode := reader.ReadUint8()
		// 根据操作码构造出指令对象
		inst := instructions.NewInstruction(opcode)
		// 获取该指令对应的操作数（会根据操作数的大小对pc进行+n）
		inst.FetchOperands(reader)
		// 修改frame的pc
		frame.SetNextPC(reader.PC())
		// 执行指令
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
