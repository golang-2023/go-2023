package main

import (
	"fmt"
	parser "github.com/golang-2023/go-2023/assignment_2/parser"
	sorter "github.com/golang-2023/go-2023/assignment_2/sorter"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Invalid input. Please provide a type ('-int', '-float', '-string', or '-mix') and then your elements.")
		return
	}

	sortType := args[0]
	switch sortType {
	case "-int":
		int, err := parser.IntParser(args[1:])
		if err != nil {
			fmt.Println("Error while parsing integers:", err)
			return
		}
		sorter.SortInt(int)
		fmt.Println(int)
	case "-float":
		float, err := parser.FloatParser(args[1:])
		if err != nil {
			fmt.Println("Error while parsing floats:", err)
			return
		}
		sorter.SortFloat(float)
		fmt.Println(float)
	case "-string":
		string := args[1:]
		sorter.SortString(string)
		fmt.Println(string)
	case "-mix":
		mixed, err := parser.MixedParser(args[1:])
		if err != nil {
			fmt.Println("Error while parsing mixed values:", err)
			return
		}
		sorting.SortMixedElements(mixed)
		fmt.Println(mixed)
	default:
		fmt.Println("Invalid type. Please provide a valid type: '-int', '-float', '-string', or '-mix'")
	}
}
