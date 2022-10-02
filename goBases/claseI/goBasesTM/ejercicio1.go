/*Ejercicio 1 - Imprimí tu nombre
Crea una aplicación donde tengas como variable tu nombre y dirección.
Imprime en consola el valor de cada una de las variables.
*/
package main

import "fmt"

var name string = "Zoe Tira"
var addres string = "Los jilgueros"

func main()  {
	fmt.Println("Name: " + name)
	fmt.Println("Addres: " + addres)

}