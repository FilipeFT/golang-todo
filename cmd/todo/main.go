package main

import (
	"fmt"
	"github.com/FilipeFT/golang-todo"
	"os"
	"strings"
)

// Hardcoding the filename
const todoFileName = ".todo.json"

func main() {

	// Define an items list
	l := &todo.List{}

	// Use the Get method to read to do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments provided
	switch {

	case len(os.Args) == 1:

		for _, item := range *l {
			fmt.Println(item.Task)
		}

	default:

		item := strings.Join(os.Args[1:], " ")

		l.Add(item)

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	}
}
