package main

import "fmt"

func main() {
	// hello.PrintHello()
	// values.PrintValues()
	// variables.PrintVariables()
	//constants.PrintConstants()
	// fors.PrintFor()
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
