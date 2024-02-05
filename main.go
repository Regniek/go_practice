package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	// de esta manera puedo declarar variables las dos son equivalentes
	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Println(y)

	// de esta manera go maneja las exepciones lo que hacemos es sacara un err de la declaracion del valor y luego con un condicional
	// pasarlo al log debido
	myValue, err := strconv.ParseInt("NaN", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println((myValue))
	}

	// los map son estructuras de llave valor
	// en el map debes definir que tipo de dato son tu llaves e igual que tipo de dato es el valor
	m := make(map[string]int)
	m["key"] = 6

	fmt.Println(m["key"])
	fmt.Println(m)
	// slice de enteros es algo similar a un arreglo se puede utilizar como una lista
	s := []int{1, 2, 3}
	// range normalmente s eutiliza para hacer iteraciones por ejemplo como en este slice
	// range devuelve index indice y value valor
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	// con append agregas cualquier elemento al final del slice

	s = append(s, 16)

	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	c := make(chan int)
	go doSomething(c)
	<-c

	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)

}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
