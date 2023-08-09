package main

import "fmt"

func main() {

	resultado := division(10, 0)
	fmt.Println("El resultado es:", resultado)

}

func division(a, b int) int {

	// Defer de funcion anonima para recuperar (recover) el panic
	defer func() {
		err := recover()
		if err != nil {
			// La app entro en panico pero la ejecucion termino de la forma normal
			fmt.Println("La aplicacion entro en panic")
		}
	}()
	return a / b
}
