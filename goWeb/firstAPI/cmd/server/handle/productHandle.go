/*El paquete handler con el controlador de la entidad seleccionada.
Se debe generar la estructura request
Se debe generar la estructura del controlador que tenga como campo el servicio
Se debe generar la función que retorne el controlador
Se deben generar todos los métodos correspondientes a los endpoints
*/

package handle

import (
	"net/http"
	"strconv"

	"github.com/ZoeTira/backpack-bcgow6-zoe-tira/goWeb/firstAPI/internal/products"
	"github.com/gin-gonic/gin"
)

type request struct{
	Id   int `json:"Id"`
    Name string `json:"Name"`
	Price float64 `json:"Price"`
	Count int `json:"Count"`
}

type Product struct{
	service products.Service
}

func NewProduct(p products.Service) *Product{
	return &Product{
        service: p,
    }
}

func (prod *Product)GetAll() gin.HandlerFunc{
	return func(c *gin.Context) {
		//Valido el token
        token := c.GetHeader("token")
        if token!= "1234" {
            c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"error": "Invalid token",
			})
			return
        }

		p, err := prod.service.GetAll()
		if err!= nil {
			c.JSON(http.StatusNotFound, gin.H{
                "error": err.Error(),
            })
			return
		}

		c.JSON(http.StatusOK, p)
	}
}

func (prod *Product) Store() gin.HandlerFunc{
	return func(c *gin.Context) {
        //Valido el token
        token := c.GetHeader("token")
        if token!= "1234" {
            c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"error": "Invalid token",
			})
			return
		}

		var req request
		if err := c.BindJSON(&req); err!= nil {
			c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
			return
		}

		p, err := prod.service.Store(req.Name, req.Price, req.Count)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
                "error": err.Error(),
            })
			return
		}
		c.JSON(http.StatusOK, p)

	}
}

func (prod *Product) Update() gin.HandlerFunc{
	return func(c *gin.Context) {
        //Valido el token
        token := c.GetHeader("token")
        if token!= "1234" {
            c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"error": "Invalid token",
			})
			return
		}

		//Obtengo Id y lo paso de string a base int
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err!= nil {
			c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
			return
		}

		//Obtengo mi request
		var req request
		//La paso para leerla
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
			return
		}

		//Empiezo verificaciones
		if req.Name == ""{
			c.JSON(http.StatusBadRequest, gin.H{
                "error": "Name is required",
            })
			return
		}

		if req.Price == 0{
			c.JSON(http.StatusBadRequest, gin.H{
                "error": "Price is required",
            })
			return
		}

		if req.Count == 0{
			c.JSON(http.StatusBadRequest, gin.H{
                "error": "Count is required",
            })
			return
		}

		//Si ta todo bien, llamo al service y le paso todo
		p, err := prod.service.Update(int(id), req.Name, req.Price, req.Count)
		if err!= nil {
			c.JSON(http.StatusNotFound, gin.H{
                "error": err.Error(),
            })
		}

		c.JSON(http.StatusOK, p)
	}
}

func (prod *Product) Delete() gin.HandlerFunc{
	return func(c *gin.Context) {
        //Valido el token
        token := c.GetHeader("token")
        if token!= "1234" {
            c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"error": "Invalid token",
			})
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err!= nil {
			c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
			return
		}

		//Obtengo Id y lo paso de string a base int
		id, err = strconv.ParseInt(c.Param("id"), 10, 64)
        if err!= nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
			return
		}

		p, err := prod.service.Delete(int(id))
		if 	 err!= nil {
			c.JSON(http.StatusNotFound, gin.H{
                "error": err.Error(),
            })
			return 
		}
		c.JSON(http.StatusOK, p)

		
	}
	
}

func (prod *Product) UpdateNamePrice() gin.HandlerFunc{
	return func(c *gin.Context) {
        //Valido el token
        token := c.GetHeader("token")
        if token!= "1234" {
            c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"error": "Invalid token",
			})
			return 
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err!= nil {
			c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
			return 
		}

		var req request
		if err := c.ShouldBindJSON(&req); err!= nil {
			c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
			return
		}

		if req.Name == ""{
			c.JSON(http.StatusBadRequest, gin.H{
                "error": "Name is required",
            })
			return 
		}

		if req.Price == 0{
			c.JSON(http.StatusBadRequest, gin.H{
                "error": "Price is required",
            })
			return 
		}

		p, err := prod.service.UpdateNamePrice(int(id), req.Name, req.Price)
		if err!= nil {
			c.JSON(http.StatusNotFound, gin.H{
                "error": err.Error(),
            })
			return 
		}
		
		c.JSON(http.StatusOK, p)
	}

}
