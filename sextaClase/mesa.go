package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Definición de la estructura para representar un producto
type Product struct {
	Code            string
	Name            string
	CurrentPrice    float64
	CurrentQuantity int
}

// Definición de la estructura para representar una categoría con sus productos
type Category struct {
	Code     string
	Name     string
	Products []Product
}

func main() {

	//crear un archivo csv y cargarlo de esta forma
	// 001	Electrodomésticos	ListaProductos
	// 002	Muebles		ListaProductos
	// 003	Herramientas		ListaProductos
	// 004	Pinturas		ListaProductos
	// 005	Aberturas		ListaProductos
	// 006	Construcción		ListaProductos
	// 007	Automotor 		ListaProductos

	// Definir las categorías con sus productos
	categories := []Category{
		{
			Code: "001",
			Name: "Electrodomésticos",
			Products: []Product{
				{"E001", "Lavadora", 500, 10},
				{"E002", "Refrigerador", 800, 5},
				{"E003", "Lavaplatos", 300, 2},
				{"E003", "microhondas", 200, 3},
				// Agrega más productos si lo deseas
			},
		},
		{
			Code: "002",
			Name: "Muebles",
			Products: []Product{
				{"M001", "Sofá", 300, 8},
				{"M002", "Mesa de comedor", 150, 15},
				{"M002", "Mesa de noche", 100, 5},
				{"M002", "Silla", 20, 5},
				// Agrega más productos si lo deseas
			},
		},
		{
			Code: "003",
			Name: "Herramientas",
			Products: []Product{
				{"H001", "Taladro", 100, 20},
				{"H002", "Destornillador", 10, 50},
				{"H003", "Pinzas", 107, 10},
				{"H004", "Martillo", 90, 12},
				// Agrega más productos si lo deseas
			},
		},
		// Agrega más categorías si lo deseas
	}

	// 	Generar otro archivo CSV, llamado estimaciones.csv. Este tendrá los resultados de la suma de todos los productos de cada una de las categorías.
	// Imprimir todos los estimativos por consola. Ejemplo:
	// Categoría			EstimativoPorCategoría
	// Construcción				60.700
	// Pinturas 				40.500
	// Aberturas				55.300
	// TotalEstimativo 			156.500

	// Calcular los estimativos y el estimativo total
	estimates := make(map[string]float64)
	totalEstimate := 0.0

	// Recorre cada categoría en la lista de categorías
	for _, category := range categories {
		// Inicializa una variable para mantener el estimado de la categoría actual en 0
		categoryEstimate := 0.0
		// Recorre cada producto en la categoría actual
		for _, product := range category.Products {
			// Calcula el estimado para el producto actual multiplicando su precio actual por la cantidad actual
			// y luego agrega este estimado al estimado de la categoría actual
			categoryEstimate += product.CurrentPrice * float64(product.CurrentQuantity)
		}
		// Asigna el estimado de la categoría actual al mapa 'estimates' con el nombre de la categoría como clave
		estimates[category.Name] = categoryEstimate
		// Agrega el estimado de la categoría actual al estimado total
		totalEstimate += categoryEstimate
	}

	// Crear y escribir en el archivo estimaciones.csv
	estimationsFile, err := os.Create("estimaciones.csv")
	if err != nil {
		fmt.Println("Error al crear el archivo estimaciones.csv:", err)
		return
	}
	defer estimationsFile.Close()

	// Aquí se está creando un nuevo escritor de CSV utilizando el archivo estimationsFile que hemos creado previamente. estimationsWriter se utilizará para escribir los datos en formato CSV en el archivo.
	estimationsWriter := csv.NewWriter(estimationsFile)
	//La función Flush() asegura que todos los datos pendientes de escritura se escriban en el archivo antes de que el programa termine.
	defer estimationsWriter.Flush()

	estimationsWriter.Write([]string{"Categoría", "EstimativoPorCategoría"})
	for categoryName, estimate := range estimates {
		estimationsWriter.Write([]string{categoryName, strconv.FormatFloat(estimate, 'f', 2, 64)})
	}
	estimationsWriter.Write([]string{"TotalEstimativo", strconv.FormatFloat(totalEstimate, 'f', 2, 64)})

	// Imprimir los estimativos por consola
	fmt.Println("Categoría\t\t\tEstimativoPorCategoría")
	for categoryName, estimate := range estimates {
		fmt.Printf("%-30s%s\n", categoryName, fmt.Sprintf("%.2f", estimate))
	}
	fmt.Printf("%-30s%s\n", "TotalEstimativo", fmt.Sprintf("%.2f", totalEstimate))

}
