package main

import "fmt"

type Nodo struct {
	valor     int
	siguiente *Nodo
	anterior  *Nodo
}
type ColaDoble struct {
	frente *Nodo
	final  *Nodo
	tamano int
}

func NuevoNodo(elem int) *Nodo {
	Nuevo := &Nodo{}
	Nuevo.anterior = nil
	Nuevo.siguiente = nil
	Nuevo.valor = elem
	return Nuevo
}
func NuevaCola() *ColaDoble {
	return &ColaDoble{frente: nil, final: nil, tamano: 0}
}
func (C *ColaDoble) EsVacia() bool {
	if C.tamano == 0 {
		return true
	}
	return false
}
func (C *ColaDoble) MostrarCola() {
	Aux := C.frente
	for Aux != nil {
		fmt.Println("->", Aux.valor)
		Aux = Aux.siguiente
	}
}
func (C *ColaDoble) EnconlarFrente(elem int) {
	Nueva := NuevoNodo(elem)
	if C.EsVacia() {
		C.final = Nueva
		C.frente = Nueva
		C.tamano++
	} else {
		Nueva.siguiente = C.frente
		Nueva.anterior = nil
		C.frente.anterior = Nueva
		C.frente = Nueva
		C.tamano++
	}
}
func (C *ColaDoble) EncolarFinal(elem int) {
	Nueva := NuevoNodo(elem)
	if C.EsVacia() {
		C.frente = Nueva
		C.final = Nueva
		C.tamano++
	} else {
		Nueva.anterior = C.final
		C.final.siguiente = Nueva
		C.final = Nueva
		C.tamano++
	}
}
func (C *ColaDoble) DesecoalarFrente() bool {
	if C.EsVacia() {
		return false
	}
	if C.tamano == 1 {
		C.final = nil
		C.frente = nil
		C.tamano--
		return true
	}
	C.frente = C.frente.siguiente
	C.frente.anterior = nil
	C.tamano--
	return true
}
func (C *ColaDoble) DesencolarFinal() bool {
	if C.EsVacia() {
		return false
	}
	if C.tamano == 1 {
		C.final = nil
		C.frente = nil
		C.tamano--
		return true
	}
	C.final = C.final.anterior
	C.final.siguiente = nil
	C.tamano--
	return true
}
func main() {
	Cola := NuevaCola()
	for i := 0; i < 10; i++ {
		Cola.EncolarFinal(i)
	}
	for i := 0; i < 10; i++ {
		Cola.EnconlarFrente(i)
	}
	Cola.MostrarCola()
	Cola.DesecoalarFrente()
	Cola.DesencolarFinal()
	Cola.MostrarCola()
}
