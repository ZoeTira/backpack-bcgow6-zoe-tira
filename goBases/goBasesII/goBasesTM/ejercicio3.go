/*Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y 
la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y 
que devuelva su salario.
*/
package main

import (
    "fmt"
)

func calcularSalario(min int, cat string) float64{

	switch cat {
		case "A":{
			salario := ((float64(min) * 3000.0)/60)
			salario = salario * 1.5
			return salario
		}
		case "B":{
			salario := ((float64(min) * 1500.0)/60)
			salario = salario * 1.2
			return salario
		}
		case "C":
			return ((float64(min) * 1000.0)/60)
	}
	return 0

}
func main() {
	result := calcularSalario(60, "B")
	fmt.Println(result)
}