/*Ejercicio 2 - Ecommerce
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. Para ello requieren que 
tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.
*/

package main

import "fmt"

type user struct{
	Name string
	Surname string
    Email string
	Products []Product
}

type Product struct{
    Name string
    Price float64
	Quantity int
}

func (user *user) addProduct(p *Product){
	user.Products = append(user.Products, *p)
}

func (user *user) deleteProduct(p Product){
	for i, v := range user.Products {		
        if v.Name == p.Name {

            user.Products = append(user.Products[:i], user.Products[i+1:]...)
		}
	}
}

func newProduct(name string, price float64, quantity int) Product{
	return Product{
		Name: name,
    	Price: price,
    	Quantity: quantity,
	}
}

func main() {
	user1 := user{
		Name: "John",
        Surname: "Doe",
        Email: "john.doe@gmail.com",
	}

	p1 := newProduct("Producto 1", 10.5, 2)
	p2 := newProduct("Producto 2", 11.3, 4)
	
	user1.addProduct(&p1)
	user1.addProduct(&p2)
    
	fmt.Println(user1.Products)

	user1.deleteProduct(p1)
	
	fmt.Println(user1.Products)

}