package main

import (
	"errors"
	"fmt"
)

var (
	ErrValorNegativo = errors.New("el valor no puede ser negativo")
	ErrSalarioMinimo = errors.New("el salario esta por debajo del minimo")
)

func main() {

	salario, err := getSalario(1000)

	if err != nil {
		// Compara si es el error esperado.
		if errors.Is(err, ErrSalarioMinimo) {
			fmt.Println("El error fue el esperado (err salario minimo)")
		} else {
			fmt.Println(err)
		}

		// Tambien se puede hacer con un switch.
		// switch err {
		// case ErrSalarioMinimo:
		// 	fmt.Println("El error fue el esperado (err salario minimo)")
		// default:
		// 	fmt.Println(err)
		// }

	} else {
		fmt.Println("El salario es:", salario)
	}

}

func getSalario(valor int) (int, error) {
	if valor <= 10000 {
		return 0, ErrSalarioMinimo
	}

	if valor < 0 {
		return 0, ErrValorNegativo
	}

	return valor, nil
}
