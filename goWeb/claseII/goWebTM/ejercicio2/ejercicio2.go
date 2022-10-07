/*
Ejercicio 2 - Validación de campos
Se debe implementar las validaciones de los campos al momento de enviar la petición, para
eso se deben seguir los siguientes pasos:
1. Se debe validar todos los campos enviados en la petición, todos los campos son
requeridos
2. En caso que algún campo no esté completo se debe retornar un código de error 400
con el mensaje “el campo %s es requerido”.
(En %s debe ir el nombre del campo que no está completo).
*/
package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)
//Indico la estructura de la respuesta
type request struct{
	Id int `json:"id"`
	Name string `json:"Name" binding:"required"`
	Price string `json:"Price" binding:"required"`
}
//Variable global con mis respuestas
var products []request
//Para crear el producto
func createProducts(c *gin.Context) {
	var req request

	err := c.ShouldBindJSON(&req)
//Si me da error, llamo a la funcion que valida campos
	if (err!= nil) {
		errs := validateFields(err)
		//Muestro
		c.JSON(http.StatusBadRequest, gin.H{"error": errs})
		return
	}

	//Si todo va bien
	//Obtengo ultimo id, le sumo uno, lo gurrado
	req.Id = len(products)+1

	//Agrego lo que obtube
	products = append(products, req)
	c.JSON(200, req)

}
func validateFields(err error) (errs []string) {
	//Uso el validador
	var validator validator.ValidationErrors
	if errors.As(err, &validator) {
		for _, fe := range validator {
			errs = append(errs, fmt.Sprint("El campo ", fe.Field(), " es requerido"))
		}
	}
	return errs
}
	
func main()  {
	r := gin.Default()

	//Endpoint que recibe la entidad
	r.POST("/create-products", createProducts)
	r.Run()
}