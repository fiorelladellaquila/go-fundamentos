package main

import "fmt"

type Libro struct {
	Nombre          string
	CantidadPaginas int
	Autor           Autor
}

type Autor struct {
	Nombre string
}

func (a *Autor) SetNombre(nombre string) {
	a.Nombre = nombre
}

func main() {
	autor := Autor{
		Nombre: "Juan Perez",
	}

	libro := Libro{
		Nombre:          "Go programacion",
		CantidadPaginas: 100,
		Autor:           autor,
	}

	fmt.Println(libro)
	libro.Autor.SetNombre("Jorge Gonzalez")
	fmt.Println(libro)
}
