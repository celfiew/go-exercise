package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)

func par(out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range out {
		fmt.Printf("Par: %d\n", num)
	}
}

func impar(out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range out {
		fmt.Printf("Impar: %d\n", num)
	}
}

func main() {
	// Lista de números para evaluar.
	listaNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Crear canales para comunicarse con las goroutines par e impar.
	parChanel := make(chan int)
	imparChanel := make(chan int)

	wg.Add(1)
	go par(parChanel, &wg)

	wg.Add(1)
	go impar(imparChanel, &wg)

	go func() {
		for _, num := range listaNumbers {
			if num%2 == 0 {
				parChanel <- num
			} else {
				imparChanel <- num
			}
		}
		close(parChanel)
		close(imparChanel)
	}()

	wg.Wait()
}
