package main

import "fmt"

// Product es la estructura que representa un producto.
type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

// Slice global de Product llamado Products.
var Products []Product

// Save agrega el producto actual al slice Products.
func (p *Product) Save() {
	Products = append(Products, *p)
}

// GetAll imprime todos los productos guardados en el slice Products.
func GetAll() {
	fmt.Println("Lista de productos:")
	for _, p := range Products {
		fmt.Printf("ID: %d, Nombre: %s, Precio: %.2f, Descripción: %s, Categoría: %s\n", p.ID, p.Name, p.Price, p.Description, p.Category)
	}
}

// getById retorna el producto correspondiente al ID pasado como parámetro.
func getById(id int) *Product {
	for i := range Products {
		if Products[i].ID == id {
			return &Products[i]
		}
	}
	return nil
}

func main() {
	// Creamos algunos productos y los guardamos usando el método Save.
	product1 := Product{ID: 1, Name: "Producto 1", Price: 10.99, Description: "Este es el producto 1", Category: "Categoría 1"}
	product2 := Product{ID: 2, Name: "Producto 2", Price: 20.50, Description: "Este es el producto 2", Category: "Categoría 2"}

	product1.Save()
	product2.Save()

	// Imprimimos todos los productos guardados usando el método GetAll.
	GetAll()

	// Probamos la función getById para obtener un producto por ID.
	id := 1
	product := getById(id)
	if product != nil {
		fmt.Printf("Producto encontrado - ID: %d, Nombre: %s, Precio: %.2f, Descripción: %s, Categoría: %s\n", product.ID, product.Name, product.Price, product.Description, product.Category)
	} else {
		fmt.Printf("Producto con ID %d no encontrado.\n", id)
	}
}
