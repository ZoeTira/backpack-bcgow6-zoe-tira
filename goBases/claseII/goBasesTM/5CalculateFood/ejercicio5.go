/*Ejercicio 5 - Calcular cantidad de alimento
Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. 
Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

perro necesitan 10 kg de alimento
gato 5 kg
Hamster 250 gramos.
Tarántula 150 gramos.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función
y un mensaje (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.


ejemplo:
const (
   dog = "dog"
   cat = "cat"
)
 
...
 
animalDog, msg := Animal(dog)
animalCat, msg := Animal(cat)
 
...
var amount float64
amount+= animaldog(5)
amount+= animalCat(8)*/

package main

import (
   "fmt"
)

const (
   tarantula = "Tarantula"
   hamster = "Hamster"
   dog = "Dog"
   cat = "Cat"
)

func amountDogFood(quantity int) (result float64){
   return float64(quantity) * 10.0
}
func amountCatFood(quantity int) (result float64){
   return float64(quantity) * 5
}
func amountHamsterFood(quantity int) (result float64){
   return float64(quantity) * 0.150
}
func amountTarantulaFood(quantity int) (result float64){
   return float64(quantity) * 0.150
}

func Animal(animal string, quantity int) (result float64, msg string) {
   switch animal {
      case dog:
         result = amountDogFood(quantity)
         
      case cat:
         result = amountCatFood(quantity)

      case hamster:
         result = amountHamsterFood(quantity)

      case tarantula:
         result = amountTarantulaFood(quantity)

      default:
         msg = fmt.Sprintf("Invalid animal: %s", animal)
         return

   }
   return result, msg
}

func main()  {
   foodDog, msg := Animal(dog, 3)
   
   fmt.Println(foodDog, msg)
}


