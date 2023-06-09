package object

import (
	"fmt"
	"strconv"
)

var Builtins = []struct {
	Name    string
	Builtin *Builtin
}{
	{"len", &Builtin{
		Fn: func(env interface{}, args ...Object) Object {
			if len(args) != 1 {
				return NewError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *Array:
				return &Integer{Value: int64(len(arg.Elements))}
			case *String:
				return &Integer{Value: int64(len(arg.Value))}
			default:
				return NewError("argument to 'len' not supported, got %s", args[0].Type())
			}
		}},
	},
	{"puts", &Builtin{
		Fn: func(env interface{}, args ...Object) Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return nil
		}},
	},
	{"first", &Builtin{
		Fn: func(env interface{}, args ...Object) Object {
			if len(args) != 1 {
				return NewError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return NewError("argument to 'first' must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}
			return nil
		}},
	},
	{"last", &Builtin{
		Fn: func(env interface{}, args ...Object) Object {
			if len(args) != 1 {
				return NewError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return NewError("argument to 'last' must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}
			return nil
		}},
	},
	{"rest", &Builtin{
		Fn: func(env interface{}, args ...Object) Object {
			if len(args) != 1 {
				return NewError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return NewError("argument to 'rest' must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]Object, length-1)
				copy(newElements, arr.Elements[1:length])
				return &Array{Elements: newElements}
			}
			return nil
		}},
	},
	{"push", &Builtin{
		Fn: func(env interface{}, args ...Object) Object {
			if len(args) != 2 {
				return NewError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return NewError("argument to 'push' must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*Array)
			length := len(arr.Elements)
			newElements := make([]Object, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]
			return &Array{Elements: newElements}
		}},
	},
	{"int", &Builtin{
		Fn: func(env interface{}, args ...Object) Object {
			if len(args) != 1 {
				return NewError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != STRING_OBJ {
				return NewError("argument to 'int' must be STRING, got %s", args[0].Type())
			}
			str := args[0].(*String)
			if i, err := strconv.Atoi(str.Value); err == nil {
				return &Integer{Value: int64(i)}
			}
			return NewError("Cannot Convert to int from string(%s)", str)
		}},
	},
	{"string", &Builtin{
		Fn: func(env interface{}, args ...Object) Object {
			if len(args) != 1 {
				return NewError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != INTEGER_OBJ {
				return NewError("argument to 'string' must be Integer, got %s", args[0].Type())
			}
			i := args[0].(*Integer)
			str := fmt.Sprint(i.Value)
			return &String{Value: str}
		}},
	},
}

func GetBuiltinByName(name string) *Builtin {
	for _, def := range Builtins {
		if def.Name == name {
			return def.Builtin
		}
	}
	return nil
}

func SetBuiltinFunction(name string, builtIn *Builtin) *Builtin {
	newBuiltin := struct {
		Name    string
		Builtin *Builtin
	}{
		Name: name, Builtin: builtIn,
	}
	Builtins = append(Builtins, newBuiltin)
	return newBuiltin.Builtin
}

func NewError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}
