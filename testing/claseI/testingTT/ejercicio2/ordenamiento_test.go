package ejercicio2

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T){
	numberList := []int{2, 6, 3, 8, 19, 4, 99, 5}

	sort.Ints(numberList)

	listInOrder := Order(numberList)

	assert.Equal(t, numberList, listInOrder, "deben ser iguales")




}