/*El main del programa.
Se debe importar e inyectar el repositorio, servicio y handler
Se debe implementar el router para los diferentes endpoints
*/

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ZoeTira/backpack-bcgow6-zoe-tira/goWeb/firstAPI/cmd/server/handle"
	"github.com/ZoeTira/backpack-bcgow6-zoe-tira/goWeb/firstAPI/internal/products"
)

func main()  {
	repo := products.NewRepository()
	service := products.NewService(repo)
	controller := handle.NewProduct(service)


	r := gin.Default()
    pr := r.Group("/products")
	{
		pr.POST("/create", controller.Store())
		pr.GET("/list", controller.GetAll())
		pr.PUT("/update/:id", controller.Update())
		pr.DELETE("/delete/:id", controller.Delete())
		pr.PATCH("/updateNamePrice/:id", controller.UpdateNamePrice())
	}

	r.Run()
}