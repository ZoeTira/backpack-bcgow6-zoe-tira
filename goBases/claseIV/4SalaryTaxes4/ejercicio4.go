/*Bonus Track -  Impuestos de salario #4
Vamos a hacer que nuestro programa sea un poco más complejo.
1. Desarrolla las funciones necesarias para permitir a la empresa calcular:
	a. Salario mensual de un trabajador según la cantidad de horas 
	trabajadas.
		- La función recibirá las horas trabajadas en el mes y el valor de 
		la hora como argumento.
		- Dicha función deberá retornar más de un valor 
		(salario calculado y error).
		- En caso de que el salario mensual sea igual o superior a $150.000, 
		se le deberá descontar el 10% en concepto de impuesto.
		- En caso de que la cantidad de horas mensuales ingresadas sea menor 
		a 80 o un número negativo, la función debe devolver un error. 
		El mismo deberá indicar “error: el trabajador no puede haber trabajado 
		menos de 80 hs mensuales”.*/

package main

import (
    "fmt"
	"errors"
	"time"
    "os"
)

func calculateSalary(hours float64, valuePH float64)(salary float64, err error){
	
	salary = (hours * valuePH)

	if(salary >= 150000){
		salary = salary * 0.9
		return salary, nil
	}

}
func main()  {

}
/*
	b. Calcular el medio aguinaldo correspondiente al trabajador 
		- Fórmula de cálculo de aguinaldo: 
[mejor salario del semestre] / 12 * [meses trabajados en el semestre].
		- La función que realice el cálculo deberá retornar más de un valor, 
		incluyendo un error en caso de que se ingrese un número negativo.

2. Desarrolla el código necesario para cumplir con las funcionalidades 
requeridas, utilizando “errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”. 
No olvides realizar las validaciones de los retornos de error en tu función 
“main()”.

*/
