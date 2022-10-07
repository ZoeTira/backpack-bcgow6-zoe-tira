/*Ejercicio 1 - Red social
Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura. 
Para optimizar y ahorrar memoria requieren que la estructura de usuarios ocupe el mismo lugar en memoria para el main del programa y para 
las funciones.
La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
Y deben implementarse las funciones:
Cambiar nombre: me permite cambiar el nombre y apellido.
Cambiar edad: me permite cambiar la edad.
Cambiar correo: me permite cambiar el correo.
Cambiar contraseña: me permite cambiar la contraseña.
*/

package main

import "fmt"

type User struct{
	Name string `json:"Name"`
	Surname string `json:"Surname"`
	Age  int    `json:"Age"`
	Email string `json:"Email"`
	Password string `json:"Password"`
}

func (user *User) setName(name, surname string){
	user.Name = name
	user.Surname = surname
}

func (user *User) setAge(age int){
	user.Age = age
}

func (user *User) setEmail(email string){
	user.Email = email
}

func (user *User) setPassword(password string){
    user.Password = password
}

func main() {
	u1 := &User{}

	u1.setName("Zoé", "Tira")
	u1.setAge(26)
	u1.setEmail("zoe@gmail.com")
	u1.setPassword("123456")

	fmt.Println(*u1)

}