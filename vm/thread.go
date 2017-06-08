package vm

import "fmt"

type thread struct {
	// a stack that holds call frames
	callFrameStack *callFrameStack
	// call frame pointer
	cfp int
	// data stack
	stack *stack
	// stack pointer
	sp int

	vm *VM
}

func (t *thread) getBlock(name string, filename filename) *instructionSet {
	return t.vm.getBlock(name, filename)
}

func (t *thread) getMethodIS(name string, filename filename) (*instructionSet, bool) {
	return t.vm.getMethodIS(name, filename)
}

func (t *thread) getClassIS(name string, filename filename) *instructionSet {
	return t.vm.getClassIS(name, filename)
}

func (t *thread) startFromTopFrame() {
	cf := t.callFrameStack.top()
	t.evalCallFrame(cf)
}

func (t *thread) evalCallFrame(cf *callFrame) {
	for cf.pc < len(cf.instructionSet.instructions) {
		i := cf.instructionSet.instructions[cf.pc]
		t.execInstruction(cf, i)
		if msg, yes := t.hasError(); yes {
			fmt.Println(msg)
			return
		}
	}
}

func (t *thread) hasError() (string, bool) {
	var hasError bool
	var msg string
	if t.stack.top() != nil {
		if err, ok := t.stack.top().Target.(*Error); ok {
			hasError = true
			msg = err.Message
		}
	}
	return msg, hasError
}

func (t *thread) execInstruction(cf *callFrame, i *instruction) {
	cf.pc++

	defer func() {
		if p := recover(); p != nil {
			if stackTrace == 0 {
				fmt.Printf("Internal Error: %s\n", p)
			}
			stackTrace++
			panic(p)
		}
	}()

	i.action.operation(t, cf, i.Params...)
}

func (t *thread) builtInMethodYield(blockFrame *callFrame, args ...Object) *Pointer {
	c := newCallFrame(blockFrame.instructionSet)
	c.blockFrame = blockFrame
	c.ep = blockFrame.ep
	c.self = blockFrame.self

	for i := 0; i < len(args); i++ {
		c.insertLCL(i, 0, args[i])
	}

	t.callFrameStack.push(c)
	t.startFromTopFrame()

	return t.stack.top()
}

// TODO: Use this method to replace unnecessary panics
func (t *thread) returnError(msg string) {
	err := &Error{Message: msg}
	t.stack.push(&Pointer{err})
	panic(errorMessage(msg))
}

func (t *thread) UndefinedMethodError(methodName string, receiver Object) {
	err := initializeError(UndefinedMethodErrorClass, "Undefined Method '%+v' for %+v", methodName, receiver.Inspect())
	t.stack.push(&Pointer{err})
}