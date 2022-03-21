package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

//how is code run in Go projects?
// go run filename.go

//what does the 'package main' do?
// package == project == workspace
// main == executable == program
// anything else == reusable

//what does 'import "fmt"' do?
// standard library package, used to print out lots of info in a formatted way
// by default, go program has access to no other packages, so must import them manually
// golang.org/pkg lists all packages in standard go library

// what is the 'func' ?
// func == function, () is a list of arguments to pass. inside {} is the body of the function where it does stuff.

// how is main.go organized?
// always do package declaration at very top
// underneath, list all other imported packages
// then, declare functions and tell go to do stuff.