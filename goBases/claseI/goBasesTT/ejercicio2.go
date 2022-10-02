/*Ejercicio 2 - Préstamo

Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas 
reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, 
se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará 
interés a los que su sueldo es mejor a $100.000. 
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
Edad
Empleado?
Antiguedad
Salario
*/

package main


import "fmt"

var (
	age int = 26
	employee bool= true
	seniority int = 2
	salary float64 = 150000

)
func main()  {
	if((age > 22) && (employee) && (seniority > 1)){
		if(salary > 100000){
			fmt.Println("Podrá acceder al préstamo pagando impuestos")
		} else{
			fmt.Println("Podrá acceder al préstamo sin pagar impuestos")
		}

	} else {
		fmt.Println("No podrá acceder al préstamo")
	}


}