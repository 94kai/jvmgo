package base

import "ch05/rtda"

// 每个指令都是操作码+[操作数]
type Instruction interface {
	// 提取操作数
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

// 不带操作数的指令，要嘛具体的index等已经隐含在操作码中了，如iload_1。要嘛就是该指令不需要操作数如pop
type NoOperandsInstruction struct{}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

// 带有2个字节操作数的指令，跳转指令。跳转的offset用2个字节表示
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

// 带操作数的指令，操作数占1个字节
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

// 带操作数的指令，操作数占2个字节
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
