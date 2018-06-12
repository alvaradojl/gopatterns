package main

import (
	"fmt"
)

func main() {
	fmt.Println("starting...")

	defer printTwo()
	printOne()

	fmt.Printf("\nsafe div 3,0: %d", safeDiv(3, 0))
	fmt.Printf("\nsafe div 3,2: ", safeDiv(3, 2))

	UseMaps()

	num3 := 3
	doubleNum := func() int {
		//this is known as closure
		num3 *= 2
		return num3
	}

	fmt.Printf("\ndoubleNum result: %d\n", doubleNum())

}

func safeDiv(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	solution := num1 / num2
	return solution
}

func printOne() {
	fmt.Println("printing 1...")
}

func printTwo() {
	fmt.Println("printing 2...")
}

func UseMaps() {
	presAge := make(map[string]int)
	presAge["roosevelt"] = 42
	fmt.Printf("\nmap result: %d\n", presAge["roosevelt"])
}
