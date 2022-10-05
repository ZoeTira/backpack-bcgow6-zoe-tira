/*Ejercicio 1 - Estructura un JSON
Según la temática elegida, genera un JSON que cumpla con las siguientes claves según la temática.
Los productos varían por id, nombre, color, precio, stock, código (alfanumérico), publicado (si-no), fecha de creación.
Los usuarios varían por id, nombre, apellido, email, edad, altura, activo (si-no), fecha de creación.
Las transacciones: id, código de transacción (alfanumérico), moneda, monto, emisor (string), receptor (string), fecha de
transacción.
1. Dentro de la carpeta go-web crea un archivo temática.json, el nombre tiene que ser el tema elegido, ej: products.json.
2. Dentro del mismo escribí un JSON que permita tener un array de productos, usuarios o transacciones con todas sus variantes.
*/

package e1

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct{
	Id int `json:"Id"`
    Name string `json:"Name"`
	Colour string `json:"Colour"`
	Price float32 `json:"Price"`
    Stock int `json:"Stock"`
	Code string `json:"Code"`
	Published bool `json:"Published"`
	CreationDate string `json:"creationDate"`
}/*
type user struct{
	Id int `json:"Id"`
    Name string `json:"Name"`
	Surname string `json:"Surname"`
    Email string `json:"Email"`
    Age int `json:"Age"`
	Height string `json:"Height"`
	Active bool `json:"Active"`
    creationDate string `json:"creationDate"`
}
type transaction struct{
	Id int `json:"Id"`
    TransactionCode string `json:"TransactionCode"`
	Currency string `json:"Currency"`
	Amount float32 `json:"Amount"`
	Transmitter string `json:"Transmitter"`
	Receiver string `json:"Receiver"`
	TransactionDate string `json:"TransactionDate"`
}*/

func writejson() {
     p := Product{
		Id: 1,
        Name: "Nootbook",
		Colour: "Green",
        Price: 10.00,
		Stock: 10,
        Code: "123ABC",
		Published: true,
        CreationDate: "2020-01-01",
	 }



	 jsonData, err := json.Marshal(p)

	 if(err != nil){
		fmt.Println(err.Error())
	 }

	 err1 := os.WriteFile("../jsons/products.json", []byte(jsonData), 0644)
	 if(err1 != nil){
		fmt.Println(err1.Error())
	 }
	
}