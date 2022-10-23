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

package products

import (
	"fmt"
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

type MockDB struct {
	BeforeUpdate Product
	UpdateNameWasCalled bool
	Products []Product
}

func (m *MockDB) Read(data interface{}) error{
	m.UpdateNameWasCalled = true
//Busco los datos en mi stub bd
	productsJSON, _ := json.Marshal(m.Products)

	return json.Unmarshal(productsJSON, &data)
	
}
	
func (m *MockDB) Write(data interface{}) error{
	newData, err := json.Marshal(data)
    if err!= nil {
        return err
    }
	return json.Unmarshal(newData, &m.Products)
}

func TestUpdateName(t *testing.T){
	p1 := Product{1, "Tablet", 800, 3}
	afterUpdate := Product{1, "Casa", 900, 3}

	myMockDB := MockDB{p1, false, []Product{p1}}
	myRepo := NewRepository(&myMockDB)
//	fmt.Println(myMockDB.Products)
	psw, err := myRepo.UpdateNamePrice(1, "Casa", 900)
	if err != nil {
		fmt.Println(err.Error())
	}
	/*fmt.Println(myMockDB.Products)
fmt.Println(myMockDB.UpdateNameWasCalled)*/
	assert.Equal(t, afterUpdate.Name, psw.Name)
	assert.True(t, myMockDB.UpdateNameWasCalled)
	
}
