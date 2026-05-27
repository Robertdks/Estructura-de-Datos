/*
Dado  un  arreglo  desordenado  de  números  enteros,  describa  paso  a  paso  cómo  el  algoritmo
MergeSort lo divide y lo combina.

Arreglo = [38, 27, 43, 3, 9, 82, 10]

Pasos a seguir: muestra cómo el arreglo se divide en subarreglos en cada nivel de recursión.
Indica en qué momento se empiezan a combinar los subarreglos. Explica cómo se realiza la
fusión ordenada de los elementos.
*/
package main

import "fmt"

func MergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	medio := len(a) / 2
	izquierda := MergeSort(a[:medio])
	derecha := MergeSort(a[medio:])
	return Merge(izquierda, derecha)
}
func Merge(izquierda []int, derecha []int) []int {

	resultado := []int{}

	i := 0
	j := 0

	for i < len(izquierda) && j < len(derecha) {

		if izquierda[i] < derecha[j] {

			resultado = append(resultado, izquierda[i])
			i++

		} else {

			resultado = append(resultado, derecha[j])
			j++
		}
	}

	for i < len(izquierda) {
		resultado = append(resultado, izquierda[i])
		i++
	}

	for j < len(derecha) {
		resultado = append(resultado, derecha[j])
		j++
	}

	return resultado
}
func main() {

	a := []int{38, 27, 43, 3, 9, 82, 10}

	fmt.Println(MergeSort(a))
}
