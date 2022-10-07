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

func calculateSalary(minutes int, category string) (salary float64){

	switch category {
		case "A":{
			salary = ((float64(minutes) * 3000.0)/60)
			salary *= 1.5
		}

		case "B":{
			salary = ((float64(minutes) * 1500.0)/60)
			salary *= 1.2
		}

		case "C":
			salary = ((float64(minutes) * 1000.0)/60)
	}

	return salary

}

func main() {
	result := calculateSalary(60, "B")

	fmt.Println(result)
}