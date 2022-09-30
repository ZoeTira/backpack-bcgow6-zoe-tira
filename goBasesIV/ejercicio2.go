/*Ejercicio 2 - Impuestos de salario #2

Haz lo mismo que en el ejercicio anterior pero reformulando el 
código para que, en reemplazo de “Error()”,  se implemente 
“errors.New()”.*/

package main

import (
	"fmt"
	"errors"
)

func main(){
	salary := 100000

	if (salary < 150000){
		fmt.Println(errors.New("Error: El salario ingresado no lacanza el minimo imponible"))
		return
	}
	fmt.Printf("Debe pagar impuestos")

}