/*Ejercicio 1 - Letras de una palabra
La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por 
separado para deletrearla. 
Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
Luego imprimí cada una de las letras.
*/
package main

import "fmt"

var world string

func main()  {
	world := "Hello"
	fmt.Println(len(world))

	for i := 0; i < len(world); i++ {
		fmt.Println(string(world[i]))
	}
}