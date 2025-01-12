# Task: Rewrite tools to use injected commandRunner

Files:
 - EXAMPLE:
    - pkg/tools/git/add.go
    - pkg/tools/git/commit.go
 - TARGET:
    - pkg/tools/git/push.go
 

Act step by step:
 - Examine the EXAMPLE files (in parallel)
 - Refactor the TARGET so that it would use the injected commandRunner in the same way as it is used in the EXAMPLE. Constructors must accept the commandRunner pointer.
 - Print me the code (only changes) for review
 - Once I agree, dump **the whole** changed code into the TARGET