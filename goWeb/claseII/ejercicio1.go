/*Ejercicio 1 - Crear Entidad
Se debe implementar la funcionalidad para crear la entidad. para eso se deben seguir los
siguientes pasos:
1. Crea un endpoint mediante POST el cual reciba la entidad.
2. Se debe tener un array de la entidad en memoria (a nivel global), en el cual se
deberán ir guardando todas las peticiones que se vayan realizando.
3. Al momento de realizar la petición se debe generar el ID. Para generar el ID se debe
buscar el ID del último registro generado, incrementarlo en 1 y asignarlo a nuestro
nuevo registro (sin tener una variable de último ID a nivel global).*/

package main

import (
    "encoding/json"
    "fmt"
	
    "github.com/gin-gonic/gin"
)

type request struct{
	Id int `json:"id"`
	Name string `json:"Name"`
	Price string `json:"Price"`
}

var products []request


func main()  {
	r := gin.Default()

	//Endpoint que recibe la entidad
	r.POST("/products", func(c *gin.Context) {
		var req request
		err := c.ShouldBindJSON(&req)

        if (err!= nil) {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
		//Obtengo ultimo id, le sumo uno, lo gurrado
		req.Id = len(products)+1

		//Agrego lo que obtube
		products = append(products, req)
		c.JSON(200, req)

	})
	r.Run()
}