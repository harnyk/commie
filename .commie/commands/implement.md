# Task: Implement memory tools factory

Files:
 - EXAMPLE:
    - pkg/toolfactories/fsFactory.go
    - pkg/toolfactories/gitFactory.go
 - TARGET:
    - pkg/toolfactories/memoryFactory.go

- Examine the EXAMPLE files.
- Implement TARGET in the same way as EXAMPLE.
- Memory factory must have NewGet, NewSet, NewDel tools.
- Memory factory must be injected with MemoryRepo through constructor and pass it to the tools.