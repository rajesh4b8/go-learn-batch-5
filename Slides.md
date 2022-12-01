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

## Day 3

### Git quick walk through

### Buit in types

- bool, string
- int int8 int16 int32 int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte // alias for uint8
- float32, float64
- complex64, complex128

### Variables

- `var` is the keyword
- Varaibles can only be declared once
- `:=` use this to declare and iniatialize
- Go is a type strcit lang -> you can't change the type after decliaration

### Zero values

- Variables declared without an explicit initial value are assigned with zero values
- `0` for numeric types
- `false` for boolean type
- `""` (empty string) for string type

| Type          | Zero Value  |
| -----------   | ----------- |
| Integer       | 0           |
| Floating point| 0.0         |
| Boolean       | false       |
| String        | ""          |
| Pointer       | nil         |
| Interface     | nil         |
| Slice         | nil         |
| Map           | nil         |
| Channel       | nil         |
| Function      | nil         |

### Constants

- used with keyword `const`


### Assignment-2

- Watch the recording if any doubts with git concepts
- Create a github account
- install git in your system
- Create SSH pub / private keys

```ssh-keygen -t ed25519 -C "your_email@example.com"```

- add the pub key to github.com at settings
- cheat sheets for reference
[Git cheatsheet 1](https://training.github.com/downloads/github-git-cheat-sheet)

[Git cheatsheet 2](https://about.gitlab.com/images/press/git-cheat-sheet.pdf)