package control

import "ch05/instructions/base"
import "ch05/rtda"

// 对于case可以建立索引的switch，用tableswitch，否则用LOOKUP_SWITCH
type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	// 类似于一个map，存放元素个数为case数*2,安装kv的形式记录case和offset
	matchOffsets []int32
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < self.npairs*2; i += 2 {
		if self.matchOffsets[i] == key {
			offset := self.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(self.defaultOffset))
}
