package main

import (
	"fmt"
)

// En la función main, definir una variable salary y asignarle un valor de tipo “int”. Luego, crear un error personalizado con un struct que implemente Error() con el mensaje “Error: el salario ingresado no alcanza el mínimo imponible” y lanzarlo en caso de que salary sea menor a 150.000. De lo contrario, imprimir por consola el mensaje “Debe pagar impuesto”.
// CustomError es una estructura que implementa la interfaz Error
type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

func main() {
	salary := 170000 // Cambia este valor según tus necesidades

	if salary < 150000 {
		err := &CustomError{"Error: el salario ingresado no alcanza el mínimo imponible"}
		panic(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

	// salary := 120000 // Cambia este valor según tus necesidades

	// if salary < 150000 {
	// 	err := fmt.Errorf("Error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	// 	panic(err)
	// } else {
	// 	fmt.Println("Debe pagar impuesto")
	// }
}
