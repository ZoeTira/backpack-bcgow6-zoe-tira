/*Ejercicio 2 - Clima
Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura y humedad y presión atmosférica de 
distintos lugares. 
Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te 
encuentres.
Imprime los valores de las variables en consola.
¿Qué tipo de dato le asignarías a las variables?
*/

package main

import "fmt"

var temp float64 = 18
var moisture float64 = 42
var pressure float64 = 1021

func main()  {
	fmt.Printf("Temperature: %v", temp)
	fmt.Println(temp)

	fmt.Println("Moisture: ")
	fmt.Println(moisture)

	fmt.Println("Pressure: ")
	fmt.Println(pressure)


}

