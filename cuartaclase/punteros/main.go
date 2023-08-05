package main

import "fmt"

func main() {
	var numero int = 1
	var puntero *int = &numero

	fmt.Println("Direccion de memoria de numero", &numero)

	fmt.Println("El valor del numero antes del incremento:", numero)
	incremento(puntero)
	fmt.Println("El valor del numero despues del incremento:", numero)

}

func incremento(numero *int) {
	*numero++
	fmt.Println("Direccion de memoria de numero dentro de la funcion", numero)
}
