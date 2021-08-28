package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) SetVacation(vacation bool) {
	e.vacation = vacation
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}
func (e *Employee) GetVacation() bool {
	return e.vacation
}

func main() {
	//forma 1
	e := Employee{}
	fmt.Printf("%v\n", e)
	// e.id = 1
	// e.name = "Chingon"
	// fmt.Printf("%v", e)
	// e.SetId(5)
	// e.SetName("Perej")
	// fmt.Printf("%v", e)

	// fmt.Println(e.GetId())
	// fmt.Println(e.GetName())
	//forma 2
	e2 := Employee{
		id:       1,
		name:     "Camilo",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)
	//forma 3 --> con new devuelve un apuntador
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)

}
