/*Ejercicio 1 - Impuestos de salario #1
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el salario 
ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor a 150.000. Caso contrario, 
imprime por consola el mensaje “Debe pagar impuesto”.*/

package main

import (
	"fmt"
	"os"
)

//Creo mi error personalizado de tipo struct
type myError struct{
	msg string
}
//Hago que mi error implemente el metodo Error()	
func (e *myError) Error() string{
	return fmt.Sprintf(e.msg)
}

//Creo la funcion para "llenar" mi error personalizado
func payTaxes(salary int)(message string, err error){
	if(salary < 150000){
		//Ignoso el primer mensaje y le seteo al error el msg
		return "Error", &myError{
			msg: "Error: El salario ingresado no lacanza el minimo imponible",
		}
	}
	return "Debe pagar impuesto", nil
}

func main(){
	salary := 200000

	response, err := payTaxes(salary)

	//Si el error no es nulo, o sea, existe
	if(err != nil){ 
		fmt.Println(err) //Imprimo el mensaje de error
		os.Exit(1) //Salgo del programa
	}
	fmt.Print(response)


}