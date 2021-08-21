package main

import "fmt"

func main() {
	fmt.Println("Zurna" + "8" + "Inches")
	fmt.Println("45+21", 45+21)

	fmt.Println("7/3 output will be int", 7/3)       // int
	fmt.Println("7/3 output will be float", 7.0/3.0) // float

	fmt.Println(true)          // just true
	fmt.Println(false)         // just false
	fmt.Println(!false)        // oppposite of false
	fmt.Println(!true)         // opposite of false
	fmt.Println(true && false) // true and false
	fmt.Println(true || false) // true or false
	fmt.Println("testing if true is working with normal println or not", true)
}
