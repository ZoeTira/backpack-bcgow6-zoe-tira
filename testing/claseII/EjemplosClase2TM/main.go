package main

import (
	"github.com/ZoeTira/backpack-bcgow6-zoe-tira/testing/claseII/EjemplosClase2TM/directorio"
	"fmt"
)

func main() {
	db := directorio.NewDB()
	motor := directorio.NewEngine(db)

	fmt.Printf("La versión actual es %s\n", motor.GetVersion())
}

