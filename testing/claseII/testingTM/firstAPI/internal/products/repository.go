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

import (
	"fmt"

	"github.com/ZoeTira/backpack-bcgow6-zoe-tira/testing/claseII/testingTM/firstAPI/pkg/store"
)

type Product struct {
	Id    int     `json:"Id"`
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
	Count int     `json:"Count"`
}

var ps []Product

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name string, price float64, count int) (Product, error)
	LastId() (lastId int, err error)
	Update(id int, name string, price float64, count int) (Product, error)
	Delete(id int) (Product, error)
	UpdateNamePrice(id int, name string, price float64) (Product, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

/*Voejo metodo get all
func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}*/

/*Viejo metodo store
func (r *repository) Store(id int, name string, price float64, count int)(p Product, err error){
	p = Product{id, name, price, count}

	ps = append(ps, p)
    lastId = id

	return p, nil
}*/

/*Viejo metodo last id
func (r *repository) LastId()(int, error){
	return lastId, nil
}*/

func (r *repository) Update(id int, name string, price float64, count int) (Product, error) {
	update := false
	var index int

	var ps []Product

	//Desde el repositorio llamo a la bd y leo, le paso la variable donde
	//guardar los datos
	err := r.db.Read(&ps)
	if err != nil {
		return Product{}, err
	}

	for i, p := range ps {
		if p.Id == id {
			ps[i].Name = name
			ps[i].Price = price
			ps[i].Count = count
			update = true
			index = i
		}
	}

	if !update {
		return Product{}, fmt.Errorf("Product %d not found", id)

	}
	err = r.db.Write(ps)
	if err != nil {
		return Product{}, err
	}
	return ps[index], nil
}

func (r *repository) Delete(id int) (Product, error) {
	//Creo la variable donde voy a guardar lo que leo
	var ps []Product

	//Desde el repositorio llamo a la bd y leo, le paso la variable donde
	//guardar los datos
	err := r.db.Read(&ps)
	if err != nil {
		return Product{}, err
	}

	for i, p := range ps {
		if p.Id == id {
			ps = append(ps[:i], ps[i+1:]...)
			//escribo el coso de nuevo
			err = r.db.Write(ps)
			if err != nil {
				return Product{}, err
			}
			return p, nil
		}
	}

	return Product{}, fmt.Errorf("Product &d not found", id)
}

func (r *repository) UpdateNamePrice(id int, name string, price float64) (Product, error) {
	update := false
	var index int
	//Creo la variable donde voy a guardar lo que leo
	var ps []Product

	//Desde el repositorio llamo a la bd y leo, le paso la variable donde
	//guardar los datos
	err := r.db.Read(&ps)
	if err != nil {
		return Product{}, err
	}

	for i, p := range ps {
		if p.Id == id {
			ps[i].Name = name
			ps[i].Price = price
			update = true
			index = i
		}
	}

	if !update {
		return Product{}, fmt.Errorf("Product %d not found", id)
	}
	//Llamo al metodo escribir y le paso lo que tengo que escribir
	err = r.db.Write(ps)
	if err != nil {
		return Product{}, err
	}
	return ps[index], nil
}

func (r *repository) Store(id int, name string, price float64, count int) (Product, error) {
	//Creo la variable donde voy a guardar lo que leo
	var ps []Product

	//Desde el repositorio llamo a la bd y leo, le paso la variable donde
	//guardar los datos
	err := r.db.Read(&ps)
	if err != nil {
		return Product{}, err
	}

	p := Product{id, name, price, count}
	ps = append(ps, p)

	//Llamo al metodo escribir y le paso lo que tengo que escribir
	err = r.db.Write(ps)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) GetAll() ([]Product, error) {
	var ps []Product
	//Leo t le paso donde guardar
	err := r.db.Read(&ps)
	if err != nil {
		return []Product{}, err
	}
	return ps, nil
}

func (r *repository) LastId() (lastId int, err error) {
	var ps []Product

	//Leo y traigo todo en el array
	err = r.db.Read(&ps)
	if err != nil {
		return 0, err
	}
	//Si no hay ningun producto, retorno 0 como id
	if len(ps) == 0 {
		return 0, nil
	}
	//Si trae mas de 0, voy al ultimo objeto y le saco el id
	return ps[len(ps)-1].Id, nil
}
