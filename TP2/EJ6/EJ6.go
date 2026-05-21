/*
Escriba un programa que lea 10 números y los imprima en orden inverso. ¿Qué estructura
de datos usó? Justifique la elección.
*/
package main

import "fmt"

type Nodo struct {
	valor     int
	siguiente *Nodo
}
type Pila struct {
	tope   *Nodo
	tamano int
}

// Constructor
func NuevaPila() *Pila {
	return &Pila{tope: nil}
}

//Es vacia
func (P *Pila) EsVacia() bool {
	if P.tamano == 0 {
		return true
	}
	return false
}

// Apilar
func (P *Pila) Apilar(elem int) {
	Nuevo := &Nodo{}
	Nuevo.valor = elem
	Nuevo.siguiente = P.tope
	P.tope = Nuevo
	P.tamano++
}

//Desapilar
func (P *Pila) Desapilar() bool {
	if P.EsVacia() {
		return false
	}
	P.tope = P.tope.siguiente
	P.tamano--
	return true
}

//Mostrar
func (P *Pila) Mostrar() {
	if P.EsVacia() {
		fmt.Println("La pila se encuentra vacia!")
		return
	}
	Aux := P.tope
	for Aux != nil {
		fmt.Println(" ", Aux.valor)
		Aux = Aux.siguiente
	}
}
func main() {
	Pila := NuevaPila()
	for i := 0; i < 10; i++ {
		Pila.Apilar(i + 1)
	}
	Pila.Mostrar()
}
