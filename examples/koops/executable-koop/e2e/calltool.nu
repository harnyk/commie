let name = 'factorial'
let params = {
    "n": 5
}

with-env {
    'commie.koop.tool.parameters': ($params | to json)
    'commie_koop_tool_parameters': ($params | to json)
} {
    ^koop-example-executable -t $name
}
