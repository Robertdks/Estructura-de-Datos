/*Dado un arreglo ordenado de números enteros y un número a buscar, describa paso a paso
cómo se ejecuta el algoritmo de búsqueda binaria recursivo.

Arreglo = [3, 8, 15, 19, 25, 31, 42, 57, 63, 79]     Buscar = 31

Pasos  a  seguir:  escribe  los  valores  de  inicio,  fin  y  medio  en  cada  llamada  recursiva.  Indica
qué  parte  del  arreglo  se  descarta  en  cada  paso.  Explica  qué  sucede  si  el  número  está
presente o si no se encuentra en el arreglo.
*/

package main

//import "fmt"

func BusquedaBinaria(a []int, inicio int, fin int, buscado int) int {
	if inicio > fin { // Nuestro caso base, aca vemos si aun se puede recorrer aun el arreglo porque si
		return -1 // retorna un -1 en caso de que no se encontro el elemento buscado
	}
	medio := (fin + inicio) / 2 // calculamos la mitad de nuestro arreglo
	if a[medio] == buscado {    // comparamos si la mitad del array es igual a nuestro elemento buscado en caso de que sea igual retorna el medio
		return medio
	}
	if buscado < a[medio] { // verificamos si nuestro elemento buscado se encuentra del lado derecho o izquierdo
		return BusquedaBinaria(a, inicio, medio-1, buscado) // llamamos a la funcion pero cambiamos nuestro fin  para que nuestro fin pase al medio y quedarnos con la parte izquierda del arreglo
	} else {
		return BusquedaBinaria(a, medio+1, fin, buscado) // llamamos a la funcion pero esta vez modificamos nuestro inicio,  nos quedamos con la parte derecha del array
	}
}

/*func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if BusquedaBinaria(a, 0, len(a)-1, 298) != -1 {
		fmt.Println("El numero ingresado si se encuentra en el arreglo")
	} else {
		fmt.Println("El numero ingresado no se encuentra en el arreglo")
	}

}*/
