/*Servicio, debe contener la lógica de nuestra aplicación.
Se debe crear el archivo service.go.
Se debe generar la interface Service con todos sus métodos.
Se debe generar la estructura service que contenga el repositorio.
Se debe generar una función que devuelva el Servicio.
Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..).
*/

package products

type Service interface{
	GetAll() ([]Product, error)
    Store(name string, price float64, count int) (p Product, err error)
	Update(id int, name string, price float64, count int)(Product, error)
	Delete(id int) (Product, error)
	UpdateNamePrice(id int, name string, price float64)(Product, error)
}

type service struct{
	repository Repository
}

func NewService(r Repository) Service{
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err!= nil {
        return nil, err
    }
	return ps, nil
}

func (s *service) Store(name string, price float64, count int)(p Product, err error){
	lastID, err := s.repository.LastId()
	if err!= nil {
        return p, err
    }

	lastID++
	
	p, err = s.repository.Store(lastID, name, price, count)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (s *service) Update(id int, name string, price float64, count int) (Product, error){
	p, err := s.repository.Update(id, name, price, count)
    if err!= nil {
        return p, err
    }
	return p, nil
}

func (s *service) Delete(id int) (Product, error){
	p, err := s.repository.Delete(id)
	if err!= nil {
        return Product{}, err
    }
	return p, nil

}

func (s *service) UpdateNamePrice(id int, name string, price float64)(Product, error){
	p, err := s.repository.UpdateNamePrice(id, name, price)
    if err!= nil {
        return Product{}, err
    }
	return p, nil
}
