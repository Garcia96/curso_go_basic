package main

import (
	//"curso_go/src/mypackage"
	"fmt"
	"log"
	"strconv"
	"sync"
)

func normalfunction(message string) {
	fmt.Println(message)
}

func returnValue(a int) int {
	return a * 2
}

func dobleReturn(a int) (c, d int) {
	return a, a * 2
}

func cicloFor() {
	// For condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// for forever, nunca va a parar !!
	//counter2 := 0
	//for {
	//	fmt.Println(counter2)
	//	counter2++
	//}
}

func convertirTextoANumero() {
	value, err := strconv.Atoi("35")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(value)
}

func isPalindrome(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palíndromo")
	} else {
		fmt.Println("No es palíndromo")
	}
}

type car struct {
	brand string
	year  int
}

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "ping")
}

func (myPC *pc) duplicateRam() {
	myPC.ram = myPC.ram * 2
}

func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB de ram, %d GB de disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

type figuras2D interface {
	area() float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calculate(f figuras2D) {
	fmt.Println("Area:", f.area())
}

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {

	//declaración de variables
	base := 12 // los dos puntos indican que es la primera vez que se usa la variable
	var altura int = 14
	var area int // Zero value al declarar de esta forma la variable

	//Println
	fmt.Println("base", base)
	fmt.Println("Altura", altura)
	fmt.Println("Area", area)

	//Printf
	nombre := "Academia"
	cursos := 500
	// %s string, %d int, %v si no sabe que tipo de dato es, %T para el tipo de dato
	fmt.Printf("%s tiene mas de %d sursos \n", nombre, cursos)

	//Sprintf
	message := fmt.Sprintf("%s tiene mas de %d sursos\n", nombre, cursos)
	fmt.Printf(message)

	normalfunction("Hey world!")

	value := returnValue(2)
	fmt.Println("Retorna valor", value)

	value1, value2 := dobleReturn(2) // en caso de solo necesitar un valor seria value1. _ := dobleReturn(2)
	fmt.Println("Retorna multiples valores", value1, value2)

	cicloFor()
	convertirTextoANumero()

	// defer pone en la cola de ejecución el código que tiene al frente
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array)
	fmt.Println(len(array))
	fmt.Println(cap(array))

	// slice
	slice := []int{0, 1, 2, 3, 4}
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[1:4])
	fmt.Println(slice[2:])

	// Append
	slice = append(slice, 45)
	fmt.Println(slice)

	slice = append(slice, slice...)
	fmt.Println(slice)

	// Range
	slice2 := []string{"Que", "pasa", "carnal"}

	for i, valor := range slice2 { // reemplazar 'i' por _ para que no salgan los índices
		fmt.Println(i, valor)
	}

	for i := range slice2 { // solo imprime los índices
		fmt.Println(i)
	}

	isPalindrome("amor a roma")

	// Llave valor con maps
	m := make(map[string]int)
	m["Jose"] = 14
	m["Nico"] = 26

	for i, valor := range m {
		fmt.Println(i, valor)
	}

	value3 := m["Jose"]
	fmt.Println(value3)

	value4 := m["Joses"]
	fmt.Println(value4)

	// Structs
	myCar := car{brand: "Audi", year: 2022}
	fmt.Println(myCar)

	var noMyCar car
	noMyCar.brand = "Ferrari"

	//var oCar pk.CarPublic //Struct ubicada en otro paquete pero público
	//oCar.Brand = "Chevy"

	//pk.PrintMessage()

	// Punteros
	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPc := pc{brand: "Lenovo", disk: 500, ram: 16}
	fmt.Println(myPc)
	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)

	// Interface, se crea en este caso para obtener el area de dos structs diferentes
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	calculate(myRectangulo)
	calculate(myCuadrado)

	// Lista de interfaces
	myInterface := []interface{}{"Hola", 23, false}
	fmt.Println(myInterface...)

	// Goroutines ?
	var wg sync.WaitGroup // Puede acumular un grupo de goroutines y las puede ir liberando poco a poco
	fmt.Println("Hello")
	wg.Add(1)
	go say("World", &wg) // go, para indicarle que es una goroutine
	wg.Wait()            // esperará a que todas las goroutines del WaitGroup finalicen

}
