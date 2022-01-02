package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter a starting number: ")

	var numberStr string
	fmt.Scanln(&numberStr)
	number, err := strconv.Atoi(numberStr)

	if err != nil {
		panic(err)
	}

	if number < 1 {
		fmt.Println("Number must be greater than zero.")
		os.Exit(3)
	}

	i := 1

	for {
		if (number % 2) == 0 {
			number /= 2
		} else {
			number = (number * 3) + 1
		}

		fmt.Printf("[Step %v] %v\n", i, number)

		if number == 1 {
			break
		} else {
			i++
		}
	}
}
