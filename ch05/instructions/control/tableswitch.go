package control

import "ch05/instructions/base"
import "ch05/rtda"

// Access jump table by index and jump
// 适用于case是连续数字的switch语句（书中说的是可以编码成索引表。看代码感觉就是连续数字，这样只需要记一个最小/最大/个数即可）
type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	// 为了保证defaultOffset的地址是4字节对齐，这里需要跳过0-3个字节,后面操作数的读取都是4个4个的。
	reader.SkipPadding()
	// default分支跳转的offset
	self.defaultOffset = reader.ReadInt32()
	// switch case的范围下限
	self.low = reader.ReadInt32()
	// switch case的范围上限
	self.high = reader.ReadInt32()
	// case的个数
	jumpOffsetsCount := self.high - self.low + 1
	// 各个case的offset
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	// pop出int值，如果在low/high区间，根据jumOffsets找到对应offset跳转即可。否则跳到default
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}
