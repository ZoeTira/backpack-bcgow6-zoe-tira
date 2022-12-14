/*Ejercicio 3 - Listar Entidad

Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
1. Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
2. Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
3. Genera un handler para el endpoint llamado “GetAll”.
4. Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/ZoeTira/backpack-bcgow6-zoe-tira/goWeb/claseI/domain"
	"github.com/gin-gonic/gin"
) /*
type Product struct{
	Id int `json:"Id"`
    Name string `json:"Name"`
	Colour string `json:"Colour"`
	Price float32 `json:"Price"`
    Stock int `json:"Stock"`
	Code string `json:"Code"`
	Published bool `json:"Published"`
	CreationDate string `json:"creationDate"`
}*/
var products []domain.Product

func main() {
	data, err := os.ReadFile("../../../jsons/products.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	
	//fmt.Println(string(data))
	p:=domain.Product{}
	//data, lo paso a algo que pueda leer
	data = []byte(data)
	//Lo formateo
	err = json.Unmarshal(data, &p)
    if err!= nil {
		fmt.Println(err.Error())
	}

	fmt.Println(p)

	products = append(products, p)
	//Asigno un puerto, por defecto es 8080
	router := gin.Default()

	//Hago la peticion get, uso el context de gin
	router.GET("/products", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"products": products,
		})

	})

	//Corro el puerto
	router.Run()
}
