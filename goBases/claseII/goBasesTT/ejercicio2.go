/*Ejercicio 2 - Matrix
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática 
y cuál es el valor máximo.
*/

package main

import "fmt"

type Matrix struct{
	heightMatrix int
	widthMatrix int
	values []float64
	quadratic bool
	maxValue float64

}

func (m *Matrix) setValues(values ...float64){
	numValues := m.heightMatrix * m.widthMatrix

	if(len(values) != numValues){
		fmt.Println("Error: Incorrect number of values")
        return
	}

	m.values = append(m.values, values...)

	if(m.heightMatrix == m.widthMatrix){
		m.quadratic = true
	} else {
		m.quadratic = false
	}
    
	for i, value := range values {
		if(i == 0){
			m.maxValue = value
		}
		if(value > m.maxValue){
            m.maxValue = value
        }
	}
}

func (m Matrix) printMatrix(){
	fmt.Printf("Height: %v \nWidth: %v \n", m.heightMatrix, m.widthMatrix)

	for i := 0; i < m.widthMatrix; i++ {
		fmt.Printf("  %v", m.values[i])
	}
	fmt.Println("\n")

	for i := m.heightMatrix; i < len(m.values); i++ {
		fmt.Printf("  %v", m.values[i])
	}

	fmt.Printf("\nIs quadratic? %v \nMaximun number: %v \n", m.quadratic, m.maxValue)
}


func main() {
	m := Matrix{}
	m.heightMatrix = 2
	m.widthMatrix = 2

	m.setValues(1, 2, 3, 4)
	m.printMatrix()


}