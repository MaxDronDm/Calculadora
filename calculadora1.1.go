package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Mostrar opciones de operaciones
	fmt.Println("Choose an operation:")
	fmt.Println("1. Add")
	fmt.Println("2. Substract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")

	// Leer entrada del usuario
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the desired option: ")
	opcionStr, _ := reader.ReadString('\n')
	opcion, _ := strconv.Atoi(opcionStr[:len(opcionStr)-1])

	// Leer los dos números
	fmt.Print("Enter the first number: ")
	num1Str, _ := reader.ReadString('\n')
	num1, _ := strconv.ParseFloat(num1Str[:len(num1Str)-1], 64)
	fmt.Print("Enter the second number: ")
	num2Str, _ := reader.ReadString('\n')
	num2, _ := strconv.ParseFloat(num2Str[:len(num2Str)-1], 64)

	// Realizar la operación seleccionada
	var resultado float64
	var operacion string
	switch opcion {
	case 1:
		resultado = num1 + num2
		operacion = "+"
	case 2:
		resultado = num1 - num2
		operacion = "-"
	case 3:
		resultado = num1 * num2
		operacion = "*"
	case 4:
		resultado = num1 / num2
		operacion = "/"
	default:
		fmt.Println("Opción inválida")
		return
	}

	// Imprimir el resultado
	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operacion, num2, resultado)
}
