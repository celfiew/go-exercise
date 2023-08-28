package main

import (
	"fmt"
)

func main() {
	// 	La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla.
	//  A - Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
	//  B - Luego, imprimir cada una de las letras.
	var palabra string = "diccionario"

	// var cantidad int = len(palabra) to declarate the type of the variable, improve the performance

	cantidad := len(palabra)

	// fmt.Println("Cantidad de letras que tiene la palabtra es:", cantidad)
	fmt.Printf("Cantidad de letras que tiene la palabtra %s es: %d\n", palabra, cantidad)

	// for i := 0; i < len(palabra); i++ {
	// 	fmt.Println(string(palabra[i]))
	// }

	for _, valor := range palabra {
		fmt.Println(string(valor))
	}

	for clave, valor := range palabra {
		fmt.Printf("clave: %d - Valor: %s\n", clave, string(valor))
	}

}
