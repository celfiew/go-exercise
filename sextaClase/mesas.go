// package main

// import (
// 	"fmt"
// 	"os"
// )

// // Definición de la estructura para un producto.
// type Producto struct {
// 	Codigo         string
// 	Nombre         string
// 	PrecioActual   float64
// 	CantidadActual int
// }

// // Definición de la estructura para una categoría, que incluye una lista de productos.
// type Categoria struct {
// 	Codigo    string
// 	Nombre    string
// 	Productos []Producto
// }

// // Función para generar datos de ejemplo con categorías y productos.
// func generarDatosEjemplo() []Categoria {
// 	categorias := []Categoria{
// 		{
// 			Codigo: "001",
// 			Nombre: "Electrodomésticos",
// 			Productos: []Producto{
// 				{Codigo: "E001", Nombre: "Lavadora", PrecioActual: 499.99, CantidadActual: 10},
// 				{Codigo: "E002", Nombre: "Refrigerador", PrecioActual: 799.99, CantidadActual: 15},
// 			},
// 		},
// 		{
// 			Codigo: "002",
// 			Nombre: "Muebles",
// 			Productos: []Producto{
// 				{Codigo: "M001", Nombre: "Sofá", PrecioActual: 299.99, CantidadActual: 5},
// 				{Codigo: "M002", Nombre: "Mesa de comedor", PrecioActual: 199.99, CantidadActual: 8},
// 			},
// 		},
// 		{
// 			Codigo: "003",
// 			Nombre: "Herramientas",
// 			Productos: []Producto{
// 				{Codigo: "H001", Nombre: "Taladro", PrecioActual: 79.99, CantidadActual: 20},
// 				{Codigo: "H002", Nombre: "Destornillador", PrecioActual: 9.99, CantidadActual: 50},
// 			},
// 		},
// 	}
// 	return categorias
// }

// func guardarEnCSV(categorias []Categoria) {
// 	// Crear o abrir el archivo "categorias.csv" en modo de escritura.
// 	file, err := os.Create("categorias.csv")
// 	if err != nil {
// 		fmt.Println("Error al crear el archivo:", err)
// 		return
// 	}
// 	// Asegurarse de que el archivo se cierre al final de la función.
// 	defer file.Close()

// 	// Iterar a través de cada categoría en la lista de categorías.
// 	for _, categoria := range categorias {
// 		// Crear una línea de texto en el formato: "Codigo    Nombre    ListaProductos".
// 		//  \t se utiliza para separar los valores de código y nombre con tabulaciones, creando una estructura de columnas en el archivo CSV.
// 		// \n se utiliza al final de cada línea generada para indicar que se debe cambiar a la siguiente línea
// 		line := fmt.Sprintf("%s\t%s\tListaProductos\n", categoria.Codigo, categoria.Nombre)
// 		fmt.Println(line)
// 		// Escribir la línea en el archivo.
// 		file.WriteString(line)
// 		// Iterar a través de cada producto en la lista de productos de la categoría actual.
// 		for _, producto := range categoria.Productos {
// 			productLine := fmt.Sprintf("%s\t%s\t%.2f\t%d\n", producto.Codigo, producto.Nombre, producto.PrecioActual, producto.CantidadActual)
// 			file.WriteString(productLine)
// 		}
// 	}
// 	fmt.Println("Datos guardados en categorias.csv")
// }

// func main() {
// 	datosEjemplo := generarDatosEjemplo()
// 	guardarEnCSV(datosEjemplo)
// }
