/*Ejercicio 3 - A qué mes corresponde

Realizar una aplicación que contenga una variable con el número del mes. 
Según el número, imprimir el mes que corresponda en texto. 
¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
Ej: 7, Julio
*/
package main

import (
    "fmt"
)
var monthNumber int
var months = [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

func main()  {
	monthNumber = 1

	fmt.Printf("%v - %v", monthNumber, months[monthNumber-1])

	
		
}