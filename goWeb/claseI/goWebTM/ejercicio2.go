/*Ejercicio 2 - Hola {nombre}

Crea dentro de la carpeta go-web un archivo llamado main.go
Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
Pegale al endpoint para corroborar que la respuesta sea la correcta.
*/

package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	//Asigno un puerto, por defecto es 8080
	router := gin.Default()

	//Hago la peticion get, uso el context de gin 
	router.GET("/hello", func (c *gin.Context){
		c.JSON(200, gin.H{
            "message": "Hello Zoe!",
        })

	})

	//Corro el puerto
	router.Run()

}