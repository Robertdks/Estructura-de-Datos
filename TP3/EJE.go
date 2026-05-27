package main

import (
	"fmt"
	"strings"
)

/*Implemente en GO las siguientes funciones en forma recursiva e iterativa. Calcule para cada
caso el orden de complejidad. Justifique.*/
/*Cuenta regresiva: recibe un número entero e imprime todos los números comprendidos
entre el mismo y 0.*/
//Forma recursiva
func CuentaRegresiva(n int) int {
	if n == 0 {
		return n
	}
	fmt.Println(n)
	return CuentaRegresiva(n - 1)
}

// Forma Iterativa
func CuentaRegresivaI(n int) {
	for i := n; i > 0; i-- {
		fmt.Println(i)
	}
}

/*
	Suma  de  enteros:  permite  sumar  todos  los  números  enteros  comprendidos  entre  un

parámetro de inicio y uno de fin
*/
//Forma Recursiva
func Suma(n1 int, n2 int) int {
	if n1 == n2 {
		return n1
	}
	return n1 + Suma(n1+1, n2)
}

// Forma Iterativa
func SumaI(n1 int, n2 int) int {
	Suma := 0
	for i := n1; i <= n2; i++ {
		Suma += i
	}
	return Suma
}

/*
Vocales y consonantes: devuelve la cantidad de vocales y de consonantes que contiene
una cadena.
*/
//Forma recursiva
func VocalesYConsonantes(texto string, n1 int, n2 int) (int, int) {
	if len(texto) == 0 {
		return n1, n2
	}
	texto = strings.ToLower(texto)
	switch texto[0] {
	case 'a', 'e', 'i', 'o', 'u':
		return VocalesYConsonantes(texto[1:], n1+1, n2)
	case ' ':
		return VocalesYConsonantes(texto[1:], n1, n2)
	default:
		return VocalesYConsonantes(texto[1:], n1, n2+1)
	}
}

// Forma Iterativa
func VocalesYConsonantesI(texto string) (int, int) {
	Vocales := "aeiou"
	n1 := 0
	n2 := len(texto)
	Vocales = strings.ToLower(Vocales)
	for i := 0; i < len(texto); i++ {
		for j := 0; j < len(Vocales); j++ {
			if texto[i] == Vocales[j] {
				n1++
			}
		}
	}
	n2 = n2 - n1
	return n1, n2
}

// Mayor elemento: dado un arreglo de enteros, devolver el mayor elemento
// Iterativa
func MayorI(a []int, long int) int {
	nmax := 0
	for i := 0; i < long; i++ {
		if a[i] > nmax {
			nmax = a[i]
		}
	}
	return nmax
}

// Recursiva
func Mayor(numeros []int, mayor int) int {
	if len(numeros) == 0 {
		return mayor
	}
	if numeros[0] > mayor {
		return Mayor(numeros[1:], numeros[0])
	}
	return Mayor(numeros[1:], mayor)
}

// nvertir: dado un arreglo de enteros, invertirlo
// Iterativa
func InvertirArrayI(a []int, b []int) {
	for i := 0; i < len(a); i++ {
		b[i] = a[len(a)-1-i]
	}
}
func InvertirArray(a []int, b []int, i int) []int {
	if len(a) == 0 {
		return b
	}
	b[len(a)-1] = a[i]
	return InvertirArray(a[1:], b[:], i)
}
