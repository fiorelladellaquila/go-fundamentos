package main

import (
	"fmt"
	"log"
)

func main() {

	// Camino feliz
	var nombre interface{} = "Juan"

	nombreString, ok := nombre.(string)

	if !ok {
		log.Println("No se pudo convertir el nombre") // log del error en consola.
	} else {
		fmt.Println(nombreString)
	}

	// Camino fail. La app entra en panico
	var nombre2 interface{} = false

	nombreFallo, ok := nombre2.(string) // No es el tipo de dato que estamos esperando.

	if !ok {
		panic("No se pudo convertir el nombre") // entra en panico la app.
	} else {
		fmt.Println(nombreFallo)
	}

}
