/*Repositorio, debe tener el acceso a la variable guardada en memoria.
1. Se debe crear el archivo repository.go
2. Se debe crear la estructura de la entidad
3. Se deben crear las variables globales donde guardar las entidades
4. Se debe generar la interface Repository con todos sus métodos
5. Se debe generar la estructura repository
6. Se debe generar una función que devuelva el Repositorio
7. Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..)
*/

package products

import "fmt"

type Product struct{
	Id   int `json:"Id"`
    Name string `json:"Name"`
	Price float64 `json:"Price"`
	Count int `json:"Count"`
}

var ps []Product
var lastId int

type Repository interface{
	GetAll() ([]Product, error)
    Store(id int, name string, price float64, count int) (p Product, err error)
	LastId() (lastId int, err error)
	Update(id int, name string, price float64, count int)(Product, error) 
	Delete(id int) (Product, error)
	UpdateNamePrice(id int, name string, price float64)(Product, error)
}

type repository struct{}

func NewRepository() Repository{
	return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (r *repository) Store(id int, name string, price float64, count int)(p Product, err error){
	p = Product{id, name, price, count}
	
	ps = append(ps, p)
    lastId = id

	return p, nil
}

func (r *repository) LastId()(int, error){
	return lastId, nil
	
}

func (r *repository) Update(id int, name string, price float64, count int)(Product, error){
	update := false
	var index int

	for i, p := range ps {
        if p.Id == id {
            ps[i].Name = name
            ps[i].Price = price
            ps[i].Count = count
            update = true
			index = i
		}
	}
	
	if !update{
		return Product{}, fmt.Errorf("Product %d not found", id)
		
	}
	return ps[index], nil
}

func (r *repository) Delete(id int)(Product, error){
	for i, p := range ps {
        if p.Id == id {
			ps = append(ps[:i], ps[i+1:]...)
			return p, nil
		}
	}
	
	return Product{}, fmt.Errorf("Product &d not found", id)
}

func (r *repository) UpdateNamePrice(id int, name string, price float64)(Product, error){
	update := false
    var index int

	for i, p := range ps {
        if p.Id == id {
			ps[i].Name = name
            ps[i].Price = price
			update = true
            index = i
		}
	}

	if !update{
		return Product{}, fmt.Errorf("Product %d not found", id)
	}
	return ps[index], nil
}