# Estructura de Datos

Repositorio de trabajos prácticos de la materia **Estructura de Datos**, correspondiente a 2do año, primer cuatrimestre, Ingeniería en Informática - UNNOBA.

El proyecto reúne ejercicios teóricos y prácticos hechos en **Go**, trabajando con análisis de algoritmos, TADs, estructuras lineales y recursión.

## Contenido general

- **TP1 - Análisis de algoritmos**: primeras implementaciones en Go, pilas, factorial y Fibonacci.
- **TP2 - Estructuras lineales**: especificación de TADs, pilas, colas, listas, cola doblemente terminada y ejercicios de aplicación.
- **TP3 - Recursión**: funciones recursivas e iterativas, búsqueda binaria y MergeSort.
- **TADsGo.zip**: material de apoyo con TADs en Go.

## Estructura del repositorio

```text
.
|-- README.md
|-- TADsGo.zip
|-- TP1
|   |-- EnunciadoTP1 - Análisis de algoritmos.pdf
|   |-- EJ1
|   |   `-- EJ1.go
|   |-- EJ2
|   |   `-- EJ2.go
|   |-- EJ3
|   |   `-- EJ3.go
|   `-- EJ4
|       `-- EJ4.go
|-- TP2
|   |-- EnunciadoTP2 - Estructuras Lineales.pdf
|   |-- EJ1
|   |-- EJ2
|   |-- EJ3
|   |-- EJ4
|   |   `-- EJ4.go
|   |-- EJ5
|   |-- EJ6
|   |   `-- EJ6.go
|   |-- EJ7
|   |   `-- EJ7.go
|   `-- EJ8
|       `-- EJ8.go
`-- TP3
    |-- EnunciadoTP3 - Recursion.pdf
    |-- EJE.go
    |-- EJ4.go
    `-- EJ5.go
```

## TP1 - Análisis de algoritmos

En este trabajo práctico se comenzaron a implementar estructuras y funciones básicas en Go.

| Ejercicio | Archivo | Descripción |
| --- | --- | --- |
| EJ1 | `TP1/EJ1/EJ1.go` | Implementación de una pila usando un arreglo fijo de enteros. Incluye operaciones para apilar, desapilar, consultar si está vacía y obtener el tamaño. |
| EJ2 | `TP1/EJ2/EJ2.go` | Borrador de pila dinámica implementada con nodos enlazados. Incluye estructura de nodo, pila, apilar, desapilar, tope y tamaño. |
| EJ3 | `TP1/EJ3/EJ3.go` | Función recursiva para calcular el factorial de un número ingresado por teclado. |
| EJ4 | `TP1/EJ4/EJ4.go` | Función recursiva para calcular un término de la sucesión de Fibonacci. |

## TP2 - Estructuras lineales

En este trabajo se trabajaron TADs y estructuras lineales, combinando respuestas teóricas con implementaciones.

| Ejercicio | Archivo | Descripción |
| --- | --- | --- |
| EJ1 | `TP2/EJ1` | Análisis de cola, lista enlazada, lista doblemente enlazada y cola doblemente terminada. Se describen características, operaciones válidas, operaciones no válidas y operaciones accesorias. |
| EJ2 | `TP2/EJ2` | Revisión y corrección de una especificación del TAD Pila. Incluye errores detectados, versión corregida, axiomas y justificación. |
| EJ3 | `TP2/EJ3` | Especificación inicial del TAD lista enlazada usando la plantilla de clase. |
| EJ4 | `TP2/EJ4/EJ4.go` | Implementación de una cola dinámica con nodos. Incluye encolar, desencolar, frente, tamaño, vaciar, copiar y mostrar. |
| EJ5 | `TP2/EJ5` | Respuesta teórica sobre las condiciones necesarias para que la operación longitud sea de orden constante en pila o cola. |
| EJ6 | `TP2/EJ6/EJ6.go` | Programa que usa una pila para imprimir números en orden inverso. |
| EJ7 | `TP2/EJ7/EJ7.go` | Programa que verifica si una palabra es palíndromo usando una pila. |
| EJ8 | `TP2/EJ8/EJ8.go` | Implementación de una cola doblemente terminada con inserción y eliminación por ambos extremos. |

## TP3 - Recursión

En este trabajo se implementaron funciones recursivas e iterativas, además de algoritmos clásicos basados en división de problemas.

| Ejercicio | Archivo | Descripción |
| --- | --- | --- |
| EJE | `TP3/EJE.go` | Implementaciones recursivas e iterativas de cuenta regresiva, suma de enteros, conteo de vocales y consonantes, búsqueda del mayor elemento e inversión de arreglos. |
| EJ4 | `TP3/EJ4.go` | Implementación de búsqueda binaria recursiva sobre un arreglo ordenado. |
| EJ5 | `TP3/EJ5.go` | Implementación de MergeSort con función auxiliar para fusionar subarreglos ordenados. |

## Temas trabajados

- Tipos abstractos de datos.
- Pilas estáticas y dinámicas.
- Colas dinámicas.
- Listas enlazadas.
- Colas doblemente terminadas.
- Recursión e iteración.
- Factorial y Fibonacci.
- Palíndromos usando pila.
- Búsqueda binaria recursiva.
- Ordenamiento MergeSort.
- Análisis de operaciones y complejidad.

## Cómo ejecutar los ejercicios

Los ejercicios están escritos en Go y pueden ejecutarse de forma individual desde la raíz del repositorio.

Ejemplos:

```bash
go run ./TP1/EJ1/EJ1.go
go run ./TP1/EJ3/EJ3.go
go run ./TP1/EJ4/EJ4.go
go run ./TP2/EJ4/EJ4.go
go run ./TP2/EJ6/EJ6.go
go run ./TP2/EJ7/EJ7.go
go run ./TP2/EJ8/EJ8.go
go run ./TP3/EJ5.go
```

Algunos programas piden datos por teclado, como factorial, Fibonacci y palíndromo.

## Estado actual

- El repositorio contiene el desarrollo de los tres primeros trabajos prácticos.
- Hay ejercicios teóricos guardados como archivos sin extensión dentro de `TP2`.
- Algunas implementaciones funcionan como programas completos con `main`.
- Otros archivos contienen funciones o borradores para completar, revisar o integrar más adelante.

## Próximos pasos sugeridos

- Revisar y completar los ejercicios que quedaron como borrador.
- Aplicar `gofmt` a los archivos `.go` antes de entregar o pushear.
- Probar cada ejercicio con `go run`.
- Agregar mensajes de commit claros por cada avance importante.

## Comandos útiles para Git

```bash
git status
git add README.md TP1 TP2 TP3 TADsGo.zip
git commit -m "Agrega trabajos practicos de estructura de datos"
git push
```
