/*1a. Cola
  - Caracteristicas:
      1. Cantidad de elementos: Dinamica (Puede crecer o decreser).
      2. Tipo de datos: homogenea (Todos los elementos son del mismo tipo).
      3. Forma de acceso: Restringido, solo se accede al frente o al final (cola o cabeza).
      4. Organizacion: sigue el principio FIFO (Fist In, First Out)
      5. Encapsulamiento: No se conoce la implementacion interna, solo sus operaciones.
  - Operaciones que tienen sentido:
      1. Encolar(e) -> Agrega un elemento al Final.
      2. Desencolar() -> Elimina el elemento de frente.
      3. Frente() -> Devuelve el primer elemento.
      4. Es_vacia() -> Indica si esta vacia.
      5. Tamano() -> Devuelve la cantidad de elementos.
  - Operaciones que no tienen Sentido:
      1. Insertar en una posicion intermedia.
      2. Eliminar un elemento del medio
      3. Acceder directamente a cualquier elemento que no sea el frente.
      -> Porque rompe el comportamiento FIFO
  - Operaciones con sentido no Indispensables
      1. Vaciar() -> Elimina todos los elementos.
      2. Copiar() -> Genera una copia de la cola
      3. Mostrar() -> Recorre los elementos*/

package main

import "fmt"

type Nodo struct {
	valor     int
	Siguiente *Nodo
}
type Cola struct {
	tope   *Nodo
	final  *Nodo
	tamano int
}

func NuevaCola() *Cola {
	return &Cola{tope: nil, final: nil}
}
func (C *Cola) EsVacia() bool {
	if C.tamano == 0 {
		return true
	}
	return false
}
func (C *Cola) Encolar(elemento int) {
	Nuevo := &Nodo{}
	Nuevo.valor = elemento
	Nuevo.Siguiente = nil
	if C.tope == nil {
		C.tope = Nuevo
		C.tamano++
		C.final = C.tope
		return
	}
	C.final.Siguiente = Nuevo
	C.tamano++
	C.final = Nuevo
}
func (C *Cola) Desencolar() bool {
	if C.EsVacia() {
		fmt.Println("La cola se encuentra vacia!")
		return false
	}
	C.tope = C.tope.Siguiente
	C.tamano--
	return true
}
func (C *Cola) Frente() int {
	if C.EsVacia() {
		fmt.Println("La cola se encuentra vacia")
		return 0
	}
	return C.tope.valor
}
func (C *Cola) Tamano() int {
	return C.tamano
}
func (C *Cola) Vaciar() {
	C.tope = nil
	C.final = nil
	C.tamano = 0
}
func (C *Cola) Copiar() *Cola {
	if C.EsVacia() {
		fmt.Println("La cola se encuentra vacia")
		return NuevaCola()
	}
	Aux := C.tope
	Nuevo := NuevaCola()
	for Aux != nil {
		Nuevo.Encolar(Aux.valor)
		Aux = Aux.Siguiente
	}
	return Nuevo
}
func (C *Cola) Mostrar() {
	if C.EsVacia() {
		fmt.Println("La cola se encuentra vacia!")
		return
	}
	Aux := C.tope
	for Aux != nil {
		fmt.Println(" ", Aux.valor)
		Aux = Aux.Siguiente
	}
}
func main() {
	Cola := NuevaCola()
	for i := 0; i < 10; i++ {
		Cola.Encolar(i)
	}
	Cola.Mostrar()
	Cola.Desencolar()
	Cola.Mostrar()
}
