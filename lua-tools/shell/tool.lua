function execute(params)
    -- Extract command and arguments
    local command = params.command
    local args = params.command_args
    if args then
        for i, arg in ipairs(args) do
            command = command .. " " .. arg
        end
    end

    print("Executing command:", command)

    -- Redirect stderr to stdout
    command = command .. " 2>&1"

    -- Execute the command
    local handle = io.popen(command)
    local result = handle:read("*a")
    handle:close()

    return result
end
