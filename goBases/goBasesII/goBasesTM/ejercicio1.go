/*Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, 
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario. 
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 
se le descontará además un 10%.
*/
package main

import (
    "fmt"
)

func CalcularSalario(salario float64) float64{

	if(salario > 50000){
		return salario * 0.83
	} 
	if(salario > 150000){
		return salario * 0.73
	}
	return salario
}

func main() {
	result := calcularSalario(151000)
	fmt.Println(result)
}
