package luatool

import (
	"fmt"

	"github.com/harnyk/gena"
	lua "github.com/yuin/gopher-lua"
)

type LuaHandler struct {
	Module string
}

func (h LuaHandler) FunctionName() string {
	return "execute"
}

func (h LuaHandler) Execute(params gena.H) (any, error) {
	state := lua.NewState()
	defer state.Close()

	state.SetGlobal("print", state.NewFunction(func(L *lua.LState) int {
		args := L.ToString(1)
		fmt.Println("[Lua print]:", args)
		return 0
	}))

	err := state.DoFile(h.Module)
	if err != nil {
		return nil, fmt.Errorf("error loading Lua file: %w", err)
	}

	paramTable := convertMapToLuaTable(state, params)

	result := state.GetGlobal(h.FunctionName())
	if fn, ok := result.(*lua.LFunction); ok {
		err = state.CallByParam(lua.P{
			Fn:      fn,
			NRet:    1, // Expect one return value
			Protect: true,
		}, paramTable)
		if err != nil {
			return nil, fmt.Errorf("error calling Lua function: %w", err)
		}

		resultValue := state.Get(-1)
		state.Pop(1) // Remove the result from the stack

		if resultStr, ok := resultValue.(lua.LString); ok {
			return string(resultStr), nil
		} else {
			return nil, fmt.Errorf("unexpected return type from Lua function: %v", resultValue.Type())
		}
	}

	return nil, fmt.Errorf("function %s not found in Lua module", h.FunctionName())
}
