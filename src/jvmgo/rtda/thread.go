package rtda

import rtc "jvmgo/rtda/class"

/*
JVM
  Thread
    pc
    Stack
      Frame
        LocalVars
        OperandStack
*/
type Thread struct {
    pc      int
    stack   *Stack
    // todo
}

// getters & setters
func (self *Thread) PC() (int) {
    return self.pc
}
func (self *Thread) SetPC(pc int) {
    self.pc = pc
}

func (self *Thread) IsStackEmpty() (bool) {
    return self.stack.isEmpty()
}

func (self *Thread) CurrentFrame() (*Frame) {
    return self.stack.top()
}
func (self *Thread) TopFrame() (*Frame) {
    return self.stack.top()
}

func (self *Thread) PushFrame(frame *Frame) {
    self.stack.push(frame)
}
func (self *Thread) PopFrame() (*Frame) {
    top := self.stack.pop()
    if top.onPopAction != nil {
        // todo
        top.onPopAction()
    }
    return top
}

func (self *Thread) NewFrame(method *rtc.Method) (*Frame) {
    return newFrame(self, method)
}

func NewThread(maxStackSize int) (*Thread) {
    stack := newStack(maxStackSize)
    return &Thread{0, stack}
}
