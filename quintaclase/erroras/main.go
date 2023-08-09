package main

import (
	"errors"
	"fmt"
)

// ErrorPersonalizado es una estructura de error personalizado que cumple
// con la interfaz de error.
type ErrorPersonalizado struct {
	Mensaje string
	Codigo  int
}

func (e *ErrorPersonalizado) Error() string {
	return fmt.Sprintf("Mensaje: %s - Codigo: %d", e.Mensaje, e.Codigo)
}

func main() {

	salario, err := getSalario(1000)

	if err != nil {
		var errorP *ErrorPersonalizado
		// Error as itera sobre los posibles errores esperados
		// hasta encontrar el esperado (target)
		if errors.As(err, &errorP) {
			fmt.Println("Encontramos el error personalizado")
		}

	} else {
		fmt.Println("El salario es:", salario)
	}

}

func getSalario(valor int) (int, error) {
	if valor <= 10000 {
		// retornamos el error personalizado.
		return 0, &ErrorPersonalizado{Mensaje: "El salario es el minio", Codigo: 500}
	}

	return valor, nil
}
