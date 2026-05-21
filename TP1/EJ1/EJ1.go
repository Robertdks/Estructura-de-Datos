package main

import "fmt"

type Pila struct {
	Elementos [100]int
	tope      int
}

// Constructor

func NuevaPila() *Pila {
	return &Pila{tope: 0}
}

// Apilar

func (p *Pila) Apilar(elem int) {
	if p.tope < len(p.Elementos) {
		p.Elementos[p.tope] = elem
		p.tope++
	} else {
		fmt.Print("La pila está llena")
	}
}

// Es vacia

func (p *Pila) EsVacia() bool {
	return p.tope == 0
}

//Tamaño

func (p *Pila) Tamano() int {
	return p.tope
}

// Desapilar

func (p *Pila) Desapilar() (int, bool) {
	if p.tope == 0 {
		return 0, false
	}
	p.tope--
	return p.Elementos[p.tope], true
}

func main() {
	p := NuevaPila()
	fmt.Println("Pila vacia?")
	if p.EsVacia() {
		fmt.Println("La pila esta vacia")
	} else {
		fmt.Println("La pila no esta vacia", "Tamaño", p.Tamano())
	}
	// Apilar elementos
	for i := 0; i < 10; i++ {
		p.Apilar((i))
	}
	fmt.Println("Tamaño: ", p.Tamano())
	for i := 0; i < 9; i++ {
		p.Desapilar()
	}
	fmt.Println("Pila vacia?")
	if p.EsVacia() {
		fmt.Println("La pila esta vacia")
	} else {
		fmt.Println("La pila no esta vacia", "Tamaño", p.Tamano())
	}
	fmt.Println("Tamaño: ", p.Tamano())
	for !p.EsVacia() {
		Elem, _ := p.Desapilar()
		fmt.Println("Desapilado: ", Elem)
	}
}
