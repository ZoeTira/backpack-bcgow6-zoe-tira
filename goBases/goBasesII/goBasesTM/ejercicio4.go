/*Ejercicio 4 - Calcular estadísticas
Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos 
de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que 
devuelva otra función ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de 
enteros y devuelva el cálculo que se indicó en la función anterior
Ejemplo:

*/
package main

import (
    "fmt"
	"errors"
)
const(
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
 )
func calcularMinimo(num []float64) float64{
	var numeroMenor float64
	for _, i := range num {
		if i > numeroMenor{
			numeroMenor = i
		}
	}
	return numeroMenor

}
func calcularApromedio(num []float64) float64{
	var sum float64
	for _, i := range num {
			sum += i
	}
	return sum/float64(len(num))

}
func calcularMaximo(num []float64) float64{
	var numeroMayor float64
	for _, i := range num {
		if i < numeroMayor{
			numeroMayor = i
		}
	}
	return numeroMayor
	
}

func operation(val string) (func(num[]float64), error){ 
	switch val {
        case minimum:
			return calcularMinimo

		case average:
			return calcularAVG

		case maximum:
			return calcularMaximo
		}
		return 0, error

}
  
func main(){
	minFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)
	
	

	minValue := minFunc(2.0, 3.4, 3.5, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(minValue)
    fmt.Println(averageValue)
	fmt.Println(maxValue)
}
