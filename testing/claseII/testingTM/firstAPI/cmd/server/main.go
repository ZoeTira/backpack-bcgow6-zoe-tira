/*El main del programa.
Se debe importar e inyectar el repositorio, servicio y handler
Se debe implementar el router para los diferentes endpoints
*/

package main

import (
	"github.com/ZoeTira/backpack-bcgow6-zoe-tira/testing/claseII/testingTM/firstAPI/cmd/server/handle"
	"github.com/ZoeTira/backpack-bcgow6-zoe-tira/testing/claseII/testingTM/firstAPI/internal/products"
	"github.com/ZoeTira/backpack-bcgow6-zoe-tira/testing/claseII/testingTM/firstAPI/pkg/store"
	"github.com/gin-gonic/gin"

	"log"

	"github.com/joho/godotenv"
)

func main()  {
	//Leo mi .env
	err := godotenv.Load()
	db := store.New(store.FileType, "./products.json")
	if err!= nil {
        log.Fatal(err.Error())
    }

	repo := products.NewRepository(db)
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