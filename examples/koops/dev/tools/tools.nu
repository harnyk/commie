def params [] {
    $env.KOOP_TOOL_PARAMETERS | from json
}


export def "main ping" [] {
    try {
        let host: string = params | get host
        {} | insert res (^ping $host -c 4 out+err>| lines | last 2)      
            | to json
    } catch {|e|
        {} | insert err $e.msg | to json
    }

}

export def "main sysinfo" [] {
    try {
        uname | wrap res | to json
    } catch {|e|
        $e.msg | wrap error | to json
    }
}

export def "main time" [] {
    {}
    | insert utc   ( date now | date to-timezone utc)
    | insert local ( date now )
    | to json
}

export def main [] {}