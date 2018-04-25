package vm

import (
	"github.com/goby-lang/goby/compiler/lexer"
	"github.com/goby-lang/goby/compiler/parser"
	"github.com/goby-lang/goby/vm/classes"
	"github.com/goby-lang/goby/vm/errors"
)

// Class methods --------------------------------------------------------
func builtInRipperClassMethods() []*BuiltinMethodObject {
	return []*BuiltinMethodObject{
		{
			Name: "parse",
			Fn: func(receiver Object, sourceLine int) builtinMethodBody {
				return func(t *thread, args []Object, blockFrame *normalCallFrame) Object {
					if len(args) != 1 {
						return t.vm.initErrorObject(errors.ArgumentError, sourceLine, "Expect 1 argument. got=%d", len(args))
					}

					arg, ok := args[0].(*StringObject)
					if !ok {
						return t.vm.initErrorObject(errors.TypeError, sourceLine, errors.WrongArgumentTypeFormat, classes.StringClass, arg.Class().Name)
					}

					l := lexer.New(arg.toString())
					p := parser.New(l)
					program, err := p.ParseProgram()

					if err != nil {
						return t.vm.initErrorObject(errors.TypeError, sourceLine, errors.InternalError, classes.StringClass, arg.Class().Name)
					}

					ps := program.String()

					return t.vm.initStringObject(ps)
				}
			},
		},
	}
}

// Internal functions ===================================================

// Functions for initialization -----------------------------------------

func initRipperClass(vm *VM) {
	rp := vm.initializeClass("Ripper", false)
	rp.setBuiltinMethods(builtInRipperClassMethods(), true)
	vm.objectClass.setClassConstant(rp)
}
