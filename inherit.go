package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

// en GO no existe la herencia existe la composicion sobre herencia
// lo que se hace en la erencia es pasar el struct del Employee y
// de Person dentro del nuevo struct eso ya nos permitira halar algo similar a la herencia
type FullTimeEmployee struct {
	Person
	Employee
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

// las interfaces no se implementan de manera explicita en Go
// forma implicita para agregar interfaces en ciertas clases o structs
func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Time Employee"
}

// func GetMessage(p Person) {
// 	fmt.Printf("%s with age %d\n", p.name, p.age)
// }

// se genera un interfaz para reemplazar el metodo anterior y evitar repetir codigo
// la interface es simila r afirmamr un contrato donde te comprometes a implementar todol lo que esta en ella

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Juanito"
	ftEmployee.id = 5
	ftEmployee.age = 28
	fmt.Printf("%v\n", ftEmployee)
	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
	//GetMessage(ftEmployee.Person)
}
