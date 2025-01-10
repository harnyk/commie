# Task: group the tools into a single package - filesystem

Files:
 - EXAMPLE:
    - pkg/tools/git/{tool_name}.go
 - SOURCE:
    - pkg/tools/{tool_name}/{tool_name}.go, where {tool_name} is:
        - dump
        - list
        - ls
        - pwd
        - realpath
        - rename
        - rm
 - TARGET:
    - pkg/tools/filesystem/{tool_name}.go

Task:
 - Deeply analyze the entire content of EXAMPLE, which is a package containing tools for the git operations
 - Pay attention on the naming convention
 - Now analyze the content of SOURCE, which a collection of packages containing tools for the file operations
 - Create a single package TARGET, which will contain all the tools for the file operations, organized in the way shown in EXAMPLE