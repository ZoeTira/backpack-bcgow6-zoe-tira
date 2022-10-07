/*Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

package main

import (
	"fmt"
	"os"
)

type Product struct{
	Id int
	Name string
    Price float64
	Quantity int
}

func (p Product) formatProductData() string {
	return fmt.Sprintf("ID: %d - Name: %s - Price: %f - Quantity: %d\n", p.Id, p.Name, p.Price, p.Quantity)
}

func main() {
    p1 := Product{
		Id: 1,
        Name: "Tesla",
		Price: 100.0,
        Quantity: 1,
	}
    p2 := Product{
		Id: 2,
		Name: "BMW",
        Price: 200.0,
		Quantity: 2,
	}

	os.Create("products.csv")

	f, err1 := os.OpenFile("products.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if(err1 != nil){
		fmt.Println(err1)
        return
	}
	_, err := f.Write([]byte(p1.formatProductData()))
	if(err != nil){
		fmt.Println(err)
	}

	_, err2 := f.Write([]byte(p2.formatProductData()))
	if(err2 != nil){
		fmt.Println(err2)
	}

}