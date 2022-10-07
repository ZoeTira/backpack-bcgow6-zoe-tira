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

func calculateMinimum(values ...float64) (min float64){

	for i, value := range values {
		if(i == 0){
			min = value
		} else{
			if(value < min){
			min = value
			}
		} 
	}
	return min
}

func calcularApromedio(values ...float64) (average float64){
	sum := 0.0

	for _, value := range values {
			sum += value
	}

	average = sum / float64(len(values))

	return average
}

func calcularMaximun(values ...float64) (max float64){

	for _, value := range values {
		if (value > max){
			max = value
		}
	}

	return max
	
}

func operation(operation string, values ...float64)(result float64, err error){
	switch operation {
	case "minimum":
		result = calculateMinimum(values...)
	
	case "maximum":
		result = calcularMaximun(values...)
	
	case "average":
		result = calcularApromedio(values...)
	
	default:
		err = errors.New("Operation not fount")
		return
		
	}
	return result, nil

} 
  
func main(){
	vare := "resto"
	minFunc, err := operation(minimum, 2.0, 3.4, 3.5, 4, 10, 2, 4, 5)
	averageFunc, err2 := operation(average, 4, 5)
	maxFunc, err3 := operation(maximum, 2, 3, 3, 4, 1, 2, 4, 5)
	notFound, err4 := operation(vare, 3, 5, 6, 3, 6)

	if(err != nil){
		fmt.Println(err.Error())
	}
	if(err2 != nil){
		fmt.Println(err2.Error())
	}
	if(err3 != nil){
		fmt.Println(err3.Error())
	}
	if(err4 != nil){
		fmt.Println(err4.Error())
	}

	fmt.Println(minFunc)
    fmt.Println(averageFunc)
	fmt.Println(maxFunc)
	fmt.Println(notFound)
}
