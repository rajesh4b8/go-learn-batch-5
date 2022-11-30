# Slides for daily session

## Day 1

- Intro

## Day 2

- Install Go
  - Windows by using Choco
  - Mac by Brew
- Install Visual Studio Code

### Understanding the basic program

- The three main areas to fous
  - package main-> every file needs to define a package, package main is mandatory for any program to run
  - import -> to get and use other packages, only needed if you are using other packages
  - function main()-> you need a main() function to run the program

### Understanding packages

- main package - Executable package, every program needs a main package if we have to run the program
- Reusable package - These packages will be imported to other package to run - Reusable code / share code accorss different modules

### Go CLI

- Go is a tool for managing Go source code.
- Command line interface, you get the cli when you install go.
- You use this cli to work with Go
  - `go build` - This will compiles and generates the binary. You need a main package to generate binary
  - `go run` - This will build and run the binary
  - `go fmt` - Formats the code
  - `go install` - To install a go library / tool
  - `go get` - To get a package to our system
  - `go test` - To run the unit tests

### Assignment-1

- Install Go and VsCode
- Run the helloworld program successfully

