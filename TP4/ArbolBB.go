/*
Dada la estructura interna propuesta en teoría para un ABB en Go:

	type ArbolBB struct {
	    valor int
	    hijo_izq, hijo_der *ArbolBB
	}

Implemente  el  método  buscar  de  forma  iterativa  y  recursiva  donde  dicho
método reciba un valor del tipo int y retorna el nodo (ArbolBB) que lo contiene o
nil si no está en el árbol.
*/
package main

type Nodo struct {
	valor     int
	derecho   *Nodo
	izqueirdo *Nodo
}

func CrearArbol() *Nodo {
	Nuevo := *&Nodo{derecho: nil, izqueirdo: nil}
	return &Nuevo
}
func (ArbolBB *Nodo) Vacia() bool {
	if ArbolBB != nil {
		return false
	}
	return true
}
func (Arbol *Nodo) InsertarNodo(elemento int) *Nodo {
	if Arbol.Vacia() {
		Arbol = CrearArbol()
		Arbol.valor = elemento
	}
	if Arbol.valor < elemento {
		Arbol.derecho.InsertarNodo(elemento)
	} else {
		Arbol.izqueirdo.InsertarNodo(elemento)
	}
	return Arbol
}
