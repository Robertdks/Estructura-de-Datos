package main
import "fmt"

// Definicion del nodo
type Nodo struct{
	valor int
	sig *Nodo
}

// Definicion de la pila
type Pila struct{
	tope *Nodo
}

// Constructor
func NuevaPila()*Pila{
	return &Pila{tope: nil}
}

// Apilar 
fun (p *Pila) Apilar(elem int){
	nuevo := &Nodo{
		valor: elem,
		sig: p.tope
	}
	p.tope = nuevo
}

// EsVacia
func (p *Pila) EsVacia() bool {
	return p.tope == nil
}

// Tope
func (p *Pila) Tope()(int, bool){
	if p.EsVacia(){
		return 0, false
	}
	return p.tope.valor, true
}

// Desapilar 
func (p *Pila) Desapilar() (int, bool){
	if p.EsVacia{
		return 0, false
	}
	valor := p.tope.valor
	p.tope = p.tope.sig
	return valor,true
}

//Tamaño (Recorre la lista)
func (p *Pila) Tamano()int{
	contador := 0
	actual := p.tope
	for actual != nil{
		contador ++
		actual = actual.sig
	}
	return contador
}

func main(){
	p :=NuevaPila()
	
}