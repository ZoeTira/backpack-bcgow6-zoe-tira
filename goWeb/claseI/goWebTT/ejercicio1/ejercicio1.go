/*Ejercicio 1 - Filtremos nuestro endpoint
Según la temática elegida, necesitamos agregarles filtros a nuestro
endpoint, el mismo se tiene que poder filtrar por todos los campos.
Dentro del handler del endpoint, recibí del contexto los valores a
filtrar.
Luego genera la lógica de filtrado de nuestro array.
Devolver por el endpoint el array filtrado.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
}

func filterProductsHandler(c *gin.Context){
	products := reedProducts()
	var pp Product

	if c.ShouldBindQuery(&pp) == nil { // Setea las variables obtenidas de c.Query("nombredelavariable")
		log.Println(pp.Id)
		log.Println(pp.Name)
		log.Println(pp.Colour)
		log.Println(pp.Price)
		log.Println(pp.Stock)
		log.Println(pp.Code)
		log.Println(pp.Published)
		log.Println(pp.CreationDate)
	}

	var filter []*Product
	for _, p := range *products { // filtrado por todos los campos
		if pp.Id == p.Id && pp.Name == p.Name && pp.Colour == p.Colour && pp.Price == p.Price && pp.Stock == p.Stock && pp.Code == p.Code && pp.Published == p.Published && pp.CreationDate == p.CreationDate {
			filter = append(filter, &p)
		}
	}

	c.JSON(http.StatusOK, filter) // devovemos el array filtrado

}
func reedProducts() (products *[]Product){
	data, err := os.ReadFile("../../../jsons/products.json")
	if err != nil {
		log.Println(err.Error())
	}
	
	//fmt.Println(string(data))

	err = json.Unmarshal(data, &products)
    if err!= nil {
		log.Println(err.Error())
	}
	fmt.Println(products)
	return products

}
func main() {
	r := gin.Default()
	r.GET("/productsFilter", filterProductsHandler)

	r.Run()

}