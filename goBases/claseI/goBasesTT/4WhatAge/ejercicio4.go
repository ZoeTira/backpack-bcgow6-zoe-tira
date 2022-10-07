/*Ejercicio 4 - Qué edad tiene...
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayuda a imprimir la edad de 
Benjamin. 

  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado también es necesario: 
 - Saber cuántos de sus empleados son mayores de 21 años.
 - Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
 - Eliminar a Pedro del mapa.
*/
package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
var cont int

func main()  {

  //Imprimiendo edad de benjamin
  fmt.Printf("Edad de Benjamín: %v \n", employees["Benjamin"])
  
  //Recorriendo el map y sumando mayores de 21
  for key, _ := range employees{
    if(employees[key] > 21){
      cont++
    }
  }

  fmt.Printf("Empleados mayores de 21: %v \n", cont)

  //Agrgando a fede
  employees["Federico"] = 25
  
  //Eliminando a pedro
  delete(employees, "Pedro")

  //Chequeo que se haya actualizado el mapa
  fmt.Println(employees)

}