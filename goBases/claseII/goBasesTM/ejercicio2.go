/*Ejercicio 2 - Calcular promedio
Un colegio necesita calcular el promedio (por alumno) de sus calificaciones. Se solicita generar una función 
en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los 
números ingresados sea negativo
*/

package main

import (
    "fmt"
    "errors"
)
var sum float64


func calculateAverage(values ...float64)(avg float64, err error){

    for _, value := range values{
        if(value < 0){
            err = errors.New("Negative values are not allowed")
            return
        }
        sum += value
    }

    avg = sum / float64(len(values))

    return avg, nil
}

func main() {

    average, err := calculateAverage(6, 3, 6, -1, -2)

    if(err != nil){
        fmt.Println(err.Error())
    } else {
        fmt.Println(average)
    }

}