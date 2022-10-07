/*Ejercicio 2 - Leer archivo
La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por pantalla mostrando los valores tabulados, 
con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio 
se debe visualizar el total (Sumando precio por cantidad)
Ejemplo:

ID                            Precio  Cantidad
111223                      30012.00         1
444321                    1000000.00         4
434321                         50.50         1
                          4030062.50
*/

package main

import(
	"fmt"
	"os"
	"encoding/csv"

)
const FilePath = "./products.csv"

func formatDate(p []string){
	fmt.Printf(p[1])
}

func main() {
    file, err := os.Open(FilePath)

	if err!= nil {
        fmt.Println(err)
    }

	reader := csv.NewReader(file)
	products, _ := reader.ReadAll()

	for _, value := range products {
		fmt.Println(value)
        
		
	}

	

}