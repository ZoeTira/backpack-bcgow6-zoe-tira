/*Ejercicio 1 - Test Unitario GetAll()
Generar un Stub del Store cuya función “Read” retorne dos productos con las especificaciones que deseen. Comprobar que GetAll()
retorne la información exactamente igual a la esperada. Para esto:
Dentro de la carpeta /internal/(producto/usuario/transacción), crear un archivo repository_test.go con el test diseñado.*/

package products

import (
	"testing"

	"github.com/ZoeTira/backpack-bcgow6-zoe-tira/goWeb/claseI/domain"
	"github.com/ZoeTira/backpack-bcgow6-zoe-tira/testing/claseII/testingTM/firstAPI/internal/products"
	"github.com/stretchr/testify/assert"
)

type StubBD struct{}

func (s StubBD) Read(data interface{}) error{
	product1 := Product{1, "Tablet", 800, 3}
	product2 := Product{2, "Notebook", 1000, 7}

	ps = append(ps, product1)
	ps = append(ps, product2)

	data = &ps
	
	return nil
}

func (s StubBD) Write(data interface{}) error{
	return nil
}

func TestRead(t *testing.T){
	myStubDB := StubBD{}
	myRepo := NewRepository(myStubDB)

	expectedDate := []Product{{1, "Tablet", 800, 3}, {2, "Notebook", 1000, 7}}

	ps, err := myRepo.GetAll()
	if err != nil {
	}

	assert.Equal(t, expectedDate, ps)

}
/*Ejercicio 2 - Test Unitario UpdateName()
Diseñar Test de UpdateName, donde se valide que la respuesta retornada sea correcta para la actualización del nombre de un 
producto/usuario/transacción específico. Y además se compruebe que efectivamente se usa el método “Read” del Storage para 
buscar el producto. Para esto:
 - Crear un mock de Storage, dicho mock debe contener en su data un producto/usuario/transacción específico cuyo nombre 
 puede ser “Before Update”.

 - El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado. Puede ser a través 
 de un boolean como se observó en la clase. 

 - Para dar el test como OK debe validarse que al invocar el método del Repository UpdateName, con el id del 
 producto/usuario/transacción mockeado y con el nuevo nombre “After Update”, efectivamente haga la actualización. También debe 
 validarse que el método Read haya sido ejecutado durante el test. 
*/