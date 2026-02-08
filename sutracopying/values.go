package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("go" + "lang")
	fmt.Print("1+1 =", 1+1)

	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Printf("the values is %d", 42)
	fmt.Printf("the values of pi is approximately: %.2f", 3.14159)
	fmt.Printf("%s has %d years of experience.", "john", 5)
	fmt.Printf("the statement is %t", true)
	fmt.Printf("Hello\n%s!", "World")

	fmt.Fprintln(os.Stdout, "the values is", 43)

	values := fmt.Sprintln("the statement is", true)
	fmt.Println(values)
}
