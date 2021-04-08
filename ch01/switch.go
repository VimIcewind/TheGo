package main

import "fmt"

func main() {
	var heads, tails int
	switch coinflip() {
	case "heads":
		heads++
	case "tails":
		tails++
	default:
		fmt.Println("landed on edge!")
	}
	fmt.Printf("heads = %d\t tails = %d\n", heads, tails)
	fmt.Printf("heads's sign is %d\ntails's sign is %d\n", Signum(heads), Signum(tails))
}

func coinflip() string {
	return "heads"
}

func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}
