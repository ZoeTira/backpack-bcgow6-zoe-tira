/*Ejercicio 2 - Get one endpoint
Generar un nuevo endpoint que nos permita traer un solo resultado del array de la temática.
Utilizando path parameters el endpoint debería ser /temática/:id (recuerda que siempre tiene que ser en plural la temática).
Una vez recibido el id devuelve la posición correspondiente.
1. Genera una nueva ruta.
2. Genera un handler para la ruta creada.
3. Dentro del handler busca el item que necesitas.
4. Devuelve el item según el id.
Si no encontraste ningún elemento con ese id devolver como código de respuesta 404.
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ZoeTira/backpack-bcgow6-zoe-tira/goWeb/claseI/domain"
	"github.com/gin-gonic/gin"
)

var products []*domain.Product

func readProducts() []*domain.Product {
	data, err := os.ReadFile("../../../jsons/products.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	//fmt.Println(string(data))
	p := domain.Product{}

	err = json.Unmarshal(data, &p)
	if err != nil {
		fmt.Println(err.Error())
	}

	products = append(products, &p)
	//fmt.Println(products)
	return products
}

func productsHandler(c *gin.Context) {
	productsList := readProducts()

	id, _ := strconv.Atoi(c.Param("Id"))

	for _, p := range productsList {
		if p.Id == int(id) {
			c.JSON(http.StatusOK, p)
			break
		}
	}
//	c.String(200, "p %s", id) Es el ejemplo.... Y NO ANDA :,(
	c.JSON(http.StatusNotFound, nil)
//c.JSON(http.StatusOK, productsList) ni esto me muestra. aiiiish
}

func main() {

	r := gin.Default()
	r.GET("/products/:Id", productsHandler)

	r.Run()
}
