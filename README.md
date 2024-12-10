# GO_tutorial

make() => allocate memory and INIT, and non-zeroed storage

new() => allocate memory and no INIT, zeroed storage

Both use memory address

Garbage collection happens automatically for out of scope or nil variables

You can control GC value even.

Concurrency V/S Parallelism

GOroutines managed by Go runtime, and flexible stack - 2kb; faster

Do not communicate by sharing the memory instead share memory by communicating