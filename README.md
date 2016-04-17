# whereami 
## A Go library for code tracing

whereami has one exported function: WhereAmI() which acts as a wrapper for runtime.Caller().  The WhereAmi() function returns a formatted string indicating the current source file name, function name, and line number.

The WhereAmI() function expects either one or zero integer parameters.  In either case, the WhereAmI() function receives the parameters as a slice of ints.  If that slice is nil, no parameters have been specified.  If a parameter is specified, it will appear as entry number zero in the slice and will be passed on to runtime.Caller().

The example programs include:

src/whereami_example1.go - This is a program that shows simple, inline use of the WhereAmI() function.  The output is as follows:

File: whereami_example1.go  Function: main.main Line: 15

src/whereami_example2.go - This is a program showing where the code is executing in main() and a factorial function called fact().

src/whereami_example3.go - This program demonstrates a runtime issue if you embed the call to WhereAmI() in another function that you invoke.

src/whereami_example4.go - This program fixes the issue with whereami_example3.go by specifying a depth of 2 when invoking WhereAmI() from the trace() function.


