package mypackage

import "fmt"

type car struct {
	brand string
	year  int
}

// CarPublic car con acceso público
type CarPublic struct {
	Brand string
	Year  int
}

func PrintMessage() {
	fmt.Println("Función pública")
}
