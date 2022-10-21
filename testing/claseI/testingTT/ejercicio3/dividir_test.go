package ejercicio3

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T){
	num := 4
	den := 0

	_, err := Dividir(num, den)

	assert.Equal(t, "The denominator cannot be 0", err.Error(), "debe devolver error")
}