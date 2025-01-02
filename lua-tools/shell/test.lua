require("lua-tools.shell.tool")

local params = {
    binary = "ping",
    args = { "8.8.8.8", "-c", "4" }
}

local result = execute(params)
print(result)