package luatool

import (
	lua "github.com/yuin/gopher-lua"
)

func convertMapToLuaTable(L *lua.LState, m map[string]interface{}) *lua.LTable {
	// Create a new Lua table
	table := L.NewTable()

	for key, value := range m {
		// Convert the value to a Lua value
		luaValue := convertValueToLua(L, value)
		// Set the key-value pair in the Lua table
		table.RawSetString(key, luaValue)
	}

	return table
}

// convertValueToLua converts a Go value to a Lua-compatible value
func convertValueToLua(L *lua.LState, value interface{}) lua.LValue {
	switch v := value.(type) {
	case string:
		return lua.LString(v)
	case int:
		return lua.LNumber(v)
	case int64:
		return lua.LNumber(v)
	case float64:
		return lua.LNumber(v)
	case bool:
		if v {
			return lua.LTrue
		}
		return lua.LFalse
	case map[string]interface{}:
		return convertMapToLuaTable(L, v)
	case []interface{}:
		// Convert slices to Lua tables
		luaTable := L.NewTable()
		for _, item := range v {
			luaTable.Append(convertValueToLua(L, item))
		}
		return luaTable
	default:
		return lua.LNil
	}
}
