# Slides for daily session

## Day 1

- Intro
- Topics to be covered
- Why Go?
- History
- Go playground
- Hello world example
- Hello web example

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

## Day 4

### Scopes

- Local variable vs global variable

### Exporting / exposing / public members from package

- Just start the name with a Upper case!!
- Variables / functiontions with upper case starting will exported outside the package!

### Functions

- Go supports first class functions
- different ways of creating and using the functions in examples

### Flow control statements

- if / else
- for

### Assignment-3

- Print fibanacci series -> 0, 1, 1, 2, 3, 5, 8, 11 so on!
- Input is the number of elements

## Day 5

### Use Case Playing cards

- Deck of cards functionality
  - New Deck
  - Print Deck
  - Shuffle
  - Deal
  - Save to file
  - Read from file

### Slices / Arrays

- Array is a data structure to maintain a list of similar types
- Length is fixed for arrays in Go
- Index starts from 0
- how to create an array?
- Slices are reference data types which points to underlying array
- Array is a value data type
- Slice is Reference data type

### Assignment-4

- Shuffle the deck of cards

## Day 6

### Type conversions

- For converting compatible types you just use `new_type(old_value)`

### Functions / Methods

- Functions are called with package name / directly inside the same package
- Methods have a reciever arg, can be only called with the receiver

### Use case - Write to file, Print

### Assignment-5

- Read the deck from file

## Day 7

### Structs

- What is Struct?
- Definining struct
- Declaration / Initialization
- Embdedded structs
- functions / methods (Reciver arg)

### Pass by value

- Go -> calling a function or method -> args -> pass by value

### Value Types / Reference Types

| Value Type       | Referece Type |
| ---------------- | ------------- |
| int              | slices        |
| float            | maps          |
| string           | channels      |
| bool             | pointers      |
| structs          | functions     |
| arrays           |      |

## Day 8

### Maps

- Map is a data structure -> A list of key / value pairs
- All keys should be of same type
- All the values should be of same type

| Map                 | Struct               |
|-------------------- |----------------------|
| All keys - same type| different types      |
| related properties  | represent something    |
| All values - same type|can be different |
| keys - can be added | Define them at compile time |
| keys are indexed   | keys not indexed |
| Reference Type     | Value Type |

### Pointers

- User `&` to access the pointer (address)
- Use `*` to access the value from pointer
- User `*` infornt of a type (ex: person) to indiate it's a pointer type
- Pointer indirection

### Assignment-6

- Do the fib serires by using recursion
