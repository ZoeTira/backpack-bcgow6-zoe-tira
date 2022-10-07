/*
Ejercicio 3 - Validar Token
Para agregar seguridad a la aplicación se debe enviar la petición con un token,
para eso se deben seguir los siguientes pasos:

1. Al momento de enviar la petición se debe validar que un token sea enviado
2. Se debe validar ese token en nuestro código (el token puede estar hardcodeado).
3. En caso que el token enviado no sea correcto debemos retornar un error 401 y un
mensaje que “no tiene permisos para realizar la petición solicitada”.
*/
package main

import (
	"fmt"
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func validateToken(c *gin.Context) (err error){
	//Obtengo el token del header
	token := c.GetHeader("token")
	//Verifico que sea correcto
	if token != "1234"{
		return errors.New("Invalid Token")
	}
	return nil
}

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
	err := validateToken(c) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	var req request

	err = c.ShouldBindJSON(&req)
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

