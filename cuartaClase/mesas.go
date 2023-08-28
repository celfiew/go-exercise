package main

import "fmt"

const (
	P = "Pequeño"
	M = "Mediano"
	G = "Grande"
)

// Define the product interfaz including the price method.
type Producto interface {
	Precio() float64
}

//Now we create three types of structures that implements an interface to represent the three type of products: Small, medium and Big.

// Small product structure
type Pequeno struct {
	precio float64
}

// Price method implementation (Small product)
// Smaall: only has the product cost
func (p Pequeno) Precio() float64 {
	return p.precio
}

// Medium product structure
type Mediano struct {
	precio float64
}

// Price method implementation (Medium product)
// Medium: the product price + 3% more or keep it in the store.
func (m Mediano) Precio() float64 {
	return m.precio + m.precio*0.03
}

// Big product structure
type Grande struct {
	precio float64
}

// Price method implementation (Big product)
// Big: The product price + 6% more or keep it in the store + 2500
func (g Grande) Precio() float64 {
	return g.precio + g.precio*0.06 + 2500.0
}

// Factory function that create and return a product in fuction of type and price
func CrearProducto(tipo string, valor float64) Producto {
	switch tipo {
	case P:
		return Pequeno{precio: valor}
	case M:
		return Mediano{precio: valor}
	case G:
		return Grande{precio: valor}
	default:
		panic("Tipo de producto no válido")
	}
}

func main() {
	// Use example
	CanastaDeproductos := []Producto{
		CrearProducto("Pequeño", 100.0),
		CrearProducto("Mediano", 200.0),
		CrearProducto("Grande", 300.0),
	}

	fmt.Println(CanastaDeproductos)

	productoPequeno := CrearProducto("Grande", 50.0)

	fmt.Println("El valor del producto grande es :", productoPequeno.Precio())

	// variable to sum the total price
	sumaPrecios := 0.0

	// Calculate the total price for each product and we print them.
	for _, producto := range CanastaDeproductos {
		precioTotal := producto.Precio()
		sumaPrecios += precioTotal
	}
	fmt.Printf("Precio total de la canasta de productos: %.2f\n", sumaPrecios)
}
