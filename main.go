package main

import "fmt"

func FizzBuzz(n int) string {
	fizz := (n%3 == 0)
	buzz := (n%5 == 0)

	if fizz && buzz {
		return "FizzBuzz"
	} else if fizz {
		return "Fizz"
	} else if buzz {
		return "Buzz"
	} else {
		return fmt.Sprintf("%d", n)
	}
}

func main() {
	for i := 1; i < 100; i++ {
		fmt.Print(FizzBuzz(i), " ")
	}
}
