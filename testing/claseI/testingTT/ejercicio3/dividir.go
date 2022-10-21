/*
Ejercicio 3 - Test Unitario Método Dividir
Para el Método Dividir, visto en la clase:
//Funcino que recibe dos enteros (numerador y denominador) y retorna la division resultante

	func Dividir(num, den int) int{
		return num /den
	}

Cambiar el método para que no sólo retorne un entero sino también un error. Incorporar una validación en la que si el denominador
es igual a 0,  retorna un error cuyo mensaje sea “El denominador no puede ser 0”. Diseñar un test unitario que valide el error
cuando se invoca con 0 en el denominador.
Dentro de la carpeta go-testing crear un archivo dividir.go con la función a probar.
Dentro de la carpeta go-testing crear un archivo dividir test.go con el test diseñado.
*/
package ejercicio3

import "fmt"

func Dividir(num, den int) (int, error){
	if(den == 0){
		return 0, fmt.Errorf("The denominator cannot be 0")
	}
	return num /den, nil
}