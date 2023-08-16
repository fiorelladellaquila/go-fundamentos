package main

import (
	"fmt"
	"os"
	"time"
)

// output
// impresion de devolucion
// impresion de prestamos x2

func main() {

	var input string
	mensajePrestamo := make(chan string)
	mensajeDevolucion := make(chan string)

	go prestamo(mensajePrestamo)

	go devolucion(mensajeDevolucion)

	go printPrestamo(mensajePrestamo)

	go printDevolucion(mensajeDevolucion)

	fmt.Scan(&input)

	if input != "" {
		fmt.Println("Saliendo del programa")
		os.Exit(0)
	}

}

func prestamo(mensajePrestamo chan string) {
	defer close(mensajePrestamo)
	for {
		time.Sleep(time.Second / 2)
		mensajePrestamo <- "prestamo"
	}
}

func devolucion(mensajeDevolucion chan string) {
	defer close(mensajeDevolucion)
	for {
		time.Sleep(time.Second)
		mensajeDevolucion <- "devolucion"
	}
}

func printPrestamo(mensajePrestamo chan string) {
	for prestamo := range mensajePrestamo {
		fmt.Println(prestamo)
	}
}

func printDevolucion(mensajeDevolucion chan string) {
	for devolucion := range mensajeDevolucion {
		fmt.Println(devolucion)
	}
}
