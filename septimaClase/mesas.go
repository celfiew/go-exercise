package main

import (
	"fmt"
	"sync"
)

var (
	// El tipo sync.WaitGroup en Go es una estructura que se utiliza para coordinar y esperar a que un grupo de goroutines (hilos ligeros) finalice su ejecución antes de que el programa principal continúe. Es útil para sincronizar goroutines y asegurarse de que todas las operaciones que dependen de ellas se completen antes de seguir adelante.
	// WaitGroup se utiliza para esperar a que las goroutines terminen.
	wg sync.WaitGroup
)



// Función para imprimir los números pares recibidos a través del canal "out".
func par(out chan int, wg *sync.WaitGroup) {
	//Done(): Llamas a esta función dentro de cada goroutine después de que ha terminado su trabajo. Cada llamada a Done() disminuye el contador interno de la WaitGroup en 1, indicando que una goroutine ha finalizado su tarea.
	defer wg.Done() // Marca la finalización de esta goroutine al final.
	for num := range out {
		fmt.Printf("Par: %d\n", num)
	}
}

// Función para imprimir los números impares recibidos a través del canal "out".
func impar(out chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Marca la finalización de esta goroutine al final.
	for num := range out {
		fmt.Printf("Impar: %d\n", num)
	}
}

func main() {
	// Lista de números para evaluar.
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Crear canales para comunicarse con las goroutines par e impar.
	parCh := make(chan int)
	imparCh := make(chan int)

	// Add(n int): Esta función incrementa el contador interno de la WaitGroup en n. Debes llamar a esta función antes de iniciar las goroutines que deseas esperar. Cada vez que inicies una nueva goroutine que quieras esperar, aumentarás el contador.
	// Añadir 1 a la espera de la WaitGroup para la goroutine par.
	wg.Add(1) //Incrementar el contador en 1 antes de iniciar la goroutine par.
	// Iniciar la goroutine para manejar los números pares.
	go par(parCh, &wg) // Iniciar la goroutine par.

	// Añadir 1 a la espera de la WaitGroup para la goroutine impar.
	wg.Add(1) // Incrementar el contador en 1 antes de iniciar la goroutine impar.
	// Iniciar la goroutine para manejar los números impares.
	go impar(imparCh, &wg) // Iniciar la goroutine impar.

	// Goroutine anónima para enviar números a los canales correspondientes.
	go func() {
		for _, num := range numbers {
			if num%2 == 0 {
				parCh <- num // Enviar números pares al canal par.
			} else {
				imparCh <- num // Enviar números impares al canal impar.
			}
		}
		close(parCh)   // Cerrar el canal par después de enviar todos los números.
		close(imparCh) // Cerrar el canal impar después de enviar todos los números.
	}()

	// Wait(): Esta función bloquea la ejecución del programa principal hasta que el contador interno de la WaitGroup llegue a cero. En otras palabras, el programa se bloqueará en Wait() hasta que todas las llamadas a Done() hayan sido realizadas y el contador sea igual a cero.
	// Esperar a que todas las goroutines finalicen antes de salir de main.
	wg.Wait()
}
