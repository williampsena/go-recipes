package main

import (
	"log"
	"os"
	"strconv"
)

// This function is a minimal implementation of the calculator
func calculate(args []string) float64 {
	if len(args) < 3 {
		panic("invalid arguments")
	}

	x, err := strconv.Atoi(args[0])

	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(args[2])

	if err != nil {
		panic(err)
	}

	var r float64

	switch args[1] {
	case "+":
		r = float64(x + y)
	case "-":
		r = float64(x - y)
	case "x":
		r = float64(x * y)
	case "/":
		r = float64(x / y)
	default:
		log.Fatal("invalid operation")
	}

	return r
}

func main() {
	args := os.Args[1:]

	r := calculate(args)

	log.Printf("🟰  %.2f\n", r)
}
