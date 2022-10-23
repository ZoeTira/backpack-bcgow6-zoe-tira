/*Ejercicio 1 - Test Unitario GetAll()
Generar un Stub del Store cuya función “Read” retorne dos productos con las especificaciones que deseen. Comprobar que GetAll()
retorne la información exactamente igual a la esperada. Para esto:
Dentro de la carpeta /internal/(producto/usuario/transacción), crear un archivo repository_test.go con el test diseñado.*/

package products

import (
	"fmt"
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)
//Esta seria mi bd
type StubBD struct{
	Products []Product

}

func (s StubBD) Read(data interface{}) error{
//Busco los datos en mi stub bd
	productsJSON, _ := json.Marshal(s.Products)

	return json.Unmarshal(productsJSON, &data)

}

func (s StubBD) Write(data interface{}) error{
	return nil
}

func TestRead(t *testing.T){
	expectedDate := []Product{{1, "Tablet", 800, 3}, {2, "Notebook", 1000, 7}}

	myStubDB := StubBD{expectedDate}
	myRepo := NewRepository(&myStubDB)

	psw, err := myRepo.GetAll()
	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Equal(t, expectedDate, psw)

}

