package	rtda

type Frame struct{
	lower			*Frame
	localVars		LocalVars
	operandStack	*OperandStack
	thread			*Thread
	nextPC			int  //next instrcution
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame{
	return &Frame{
		thread:			thread,
		localVars:		newLocalVars(maxLocals),
		operandStack:	newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars{
	return self.localVars
}


func (self *Frame) OperandStack() *OperandStack{
	return self.operandStack
}

func (self *Frame) SetNextPC(pc int){
	self.nextPC = pc
}


func (self *Frame) Thread() *Thread{
	return self.thread
}

func (self *Frame) NextPC() int{
	return self.nextPC
}
