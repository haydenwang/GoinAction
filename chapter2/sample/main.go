package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	fmt.Println("here is my first init??")
	log.SetOutput(os.Stdout)
}

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	fmt.Println("here is my init.......")
	log.SetOutput(os.Stdout)
}

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	fmt.Println("here is my another init.......")
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	fmt.Println("....start from hayden................")

	// Perform the search for the specified term.
	search.Run("president")
	fmt.Print("here is the end....................")
}
