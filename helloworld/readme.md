# Go Cli

* go build (compile a bunch of go source code files)
* go run (compiles and executes one or two files)
* go fmt (formats all the code in each file in the current directory)
* go install( compiles and installs a package )
* go get ( download the raw source code of someone else's package )
* go test (runs any tests associated with the current project)


## What does 'package main' mean?
Package == Project == Workspace.

Types of packages..
Executable: generates a file theat we can run;
Reusable: Code used as 'helpers'. Good place to put reusable logic

## What's that 'func' ?
Tells go we're about to declare a function, set the name of the function end list of arguments
to pass the function
``` func main (){

}
```

# References
https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc
https://github.com/golang-standards/project-layout
