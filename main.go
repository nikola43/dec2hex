package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// get arguments
	args := os.Args

	// validate number of arguments
	if len(args) < 2 {
		fmt.Println("no arguments provided")
		return
	}

	if len(args) > 2 {
		fmt.Println("too many arguments")
		return
	}

	// check if argument is a valid number
	number, err := strconv.ParseInt(args[1], 10, 0)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	// convert number to hex
	hex := fmt.Sprintf("%X", number)
	fmt.Println(hex)

}
