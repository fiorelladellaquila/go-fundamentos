package main

import "fmt"

// Interfaz que describre los metodos que debe implementar un animal.
type Animal interface {
	Saludar() string
}

// Estructura Gato representa el modelo de un gato.
type Gato struct {
	saludo string
}

// Saludar es el metodo implementado por el gato.
func (g *Gato) Saludar() string {
	return g.saludo
}

// Estructura Perro representa el modelo de un perro.
type Perro struct {
	saludo string
}

// Saludar es el metodo implementado por el perro.
func (p *Perro) Saludar() string {
	return p.saludo
}

func (p *Perro) CuatraPatas() {
	fmt.Println("El perro tiene cuatro patas")
}

// Metodo factory que crea animales que cumplan con la implementacion de la interfaz animal.
func NewAnimal(nombre string) Animal {
	switch nombre {
	case "gato":
		return &Gato{saludo: "Miau"}
	case "perro":
		return &Perro{saludo: "Guau"}
	default:
		return nil
	}
}

func main() {
	gato := NewAnimal("gato")
	fmt.Println("El gato saluda:", gato.Saludar())
	perro := NewAnimal("perro")
	fmt.Println("El perro saluda:", perro.Saludar())
}
