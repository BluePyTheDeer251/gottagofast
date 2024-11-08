package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World!")
	
	// Variables
	fmt.Println("1+1: ", 1+1)
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)
	var myFloat = 2.51
	fmt.Println(myFloat)

	// Constants
	const mySum = 4+4
        const myStupidHardMulti = 6*3*25*666*2
	// const myImpossibleEquation = 5.522*10^17372
	// fmt.Println(myImpossibleEquation)
	fmt.Println(mySum)
	fmt.Println(myStupidHardMulti)
	fmt.Println(math.Sin(myFloat))
	// I have no idea on what a sin (in math) is.

	// For loops
	for C := 0; C >=251; C++ {
		fmt.Println(C)
	}

	for A := range 6 {
		if A%2 == 0 {
			continue
		}
	        fmt.Println(A)
		}
        // If/Else
	if 251%2 == 0 {
		fmt.Println("251 is even somehow")
	} else {
		fmt.Println("You get the thing, 251 is odd.")
	if 251%2 == 0 || 45%2 == 0 {
                fmt.Println("Either 251 or 45 are even.         Boo. Math.")
        }
	if myPoopyNum := 8; myPoopyNum < 0 {
		fmt.Println("myPoopyNum is somehow negative (bruh)")
	} else if myPoopyNum < 10 {
		fmt.Println("Man, myPoopyNum is a 1-digit number, bro, it's too easy.")
	} else {
		fmt.Println("myPoopyNum has more than 1 digit, which probably isn't even true,")
	// Reading Go by Example is fun
  }
}
}
