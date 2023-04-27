package main

import (
	"fmt"
)

func main() {
	var opcion int

	fmt.Println("Choose an operation")
	fmt.Println("1. Add")
	fmt.Println("2. Substract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")

	fmt.Print("Enter the first option: ")
	_, err := fmt.Scan(&opcion)
	if err != nil || opcion < 1 || opcion > 4 {
		fmt.Println("Error: option not valid.")
		return
	}

	var a, b float64
	fmt.Print("Enter the first number: ")
	_, err = fmt.Scan(&a)
	if err != nil {
		fmt.Println("Error: The first number is not valid.")
		return
	}

	fmt.Print("Enter the second number: ")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Error: The second number is not valid.")
		return
	}

	var resultado float64

	switch opcion {
	case 1:
		resultado = a + b
	case 2:
		resultado = a - b
	case 3:
		resultado = a * b
	case 4:
		if b == 0 {
			fmt.Println("Error: can't divide by 0.")
			return
		}
		resultado = a / b
	}

	fmt.Printf("The result of the operation is: %.2f\n", resultado)
}
