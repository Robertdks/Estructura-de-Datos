/*
Escriba un programa que lea una secuencia de caracteres e informe si la misma es palíndro-
mo. Utilice en la solución la implementación de alguno de los TADs especificados.
*/
package main

import "fmt"

type Nodo struct {
	valor     byte
	Siguiente *Nodo
}
type Pila struct {
	tope   *Nodo
	tamano int
}

func NuevoNodo(elem byte) *Nodo {
	Nuevo := &Nodo{}
	Nuevo.Siguiente = nil
	Nuevo.valor = elem
	return Nuevo
}
func NuevaPila() *Pila {
	return &Pila{tope: nil}
}
func (P *Pila) EsVacia() bool {
	if P.tamano == 0 {
		return true
	}
	return false
}
func (P *Pila) Apilar(elem byte) {

	Nuevo := NuevoNodo(elem)
	Nuevo.Siguiente = P.tope
	P.tope = Nuevo
	P.tamano++
}
func (P *Pila) Desapilar() bool {
	if P.EsVacia() {
		fmt.Println("La pila se encuentra vacia!")
		return false
	}
	P.tope = P.tope.Siguiente
	P.tamano--
	return true
}
func (P *Pila) Mostrar() {
	Aux := P.tope
	for Aux != nil {
		fmt.Printf("%c\n", Aux.valor)
		Aux = Aux.Siguiente
	}
}
func (P *Pila) Palindromo(tex string) bool {
	Aux := P.tope
	i := 0
	for Aux != nil {
		if Aux.valor != tex[i] {
			return false
		}
		i++
		Aux = Aux.Siguiente
	}
	return true
}
func main() {
	Pila := NuevaPila()
	var palabra string
	fmt.Println("Ingrese una palabra: ")
	fmt.Scan(&palabra)
	for i := 0; i < len(palabra); i++ {
		Pila.Apilar(palabra[i])
	}
	Pila.Mostrar()
	if Pila.Palindromo(palabra) {
		fmt.Println("La palabra es palindromo")
	} else {
		fmt.Println("La palabra no es palindromo")
	}
}
