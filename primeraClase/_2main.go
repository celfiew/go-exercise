package main

import "fmt"

func main() {

	// Realizar una aplicación que contenga una variable con el número del mes.
	// Según el número, imprimir el mes que corresponda en texto.
	// ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?

	mes := 15

	mapaMeses := map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}

	mesLetra, ok := mapaMeses[mes]

	if !ok {
		fmt.Println("Mes no existe")
	} else {
		fmt.Println(mesLetra)
	}

}
