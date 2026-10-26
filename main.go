// Paquete main define el punto de entrada del programa.
package main

import (
	"fmt"
)

// Ejercicio 1: Intercambio de valores.
// swap recibe dos punteros a enteros e intercambia sus valores en memoria[cite: 2].
func swap(a *int, b *int) {
	// Se almacena el valor apuntado por 'a' en una variable temporal (aux).
	aux := *a
	// Se sobrescribe el valor en la dirección de 'a' con el valor de 'b'.
	*a = *b
	// Se asigna el valor temporal a la dirección de 'b'.
	*b = aux
}

// Ejercicio 2: Estadísticas de un arreglo.
// stats calcula e imprime la suma, mínimo, máximo, cantidad de negativos
// y el recorrido inverso de un slice de enteros[cite: 3].
func stats(values []int) {
	// Validación de seguridad para evitar pánico en arreglos vacíos.
	if len(values) == 0 {
		return
	}

	// Inicialización de variables locales utilizando el primer elemento.
	suma := 0
	minimo := values[0]
	maximo := values[0]
	negativos := 0

	// Recorrido principal para calcular las métricas.
	for i := 0; i < len(values); i++ {
		val := values[i]
		suma += val

		if val < minimo {
			minimo = val
		}
		if val > maximo {
			maximo = val
		}
		if val < 0 {
			negativos++
		}
	}

	fmt.Printf("Suma: %d, Mínimo: %d, Máximo: %d, Negativos: %d\n", suma, minimo, maximo, negativos)

	// Recorrido inverso sin modificar el slice original[cite: 3].
	fmt.Print("Recorrido inverso: ")
	for i := len(values) - 1; i >= 0; i-- {
		fmt.Printf("%d ", values[i])
	}
	fmt.Println()
}

// Ejercicio 3: Copia independiente.
// copyArray crea y retorna un nuevo slice con los mismos elementos que el original.
func copyArray(values []int) []int {
	// Se reserva memoria para el nuevo slice utilizando make[cite: 3].
	other := make([]int, len(values))

	// Se copian los valores uno a uno mediante un ciclo propio[cite: 3].
	for i := 0; i < len(values); i++ {
		other[i] = values[i]
	}

	return other
}

// Ejercicio 4: Ordenar una copia.
// sortAscending implementa el algoritmo de ordenamiento de burbuja.
// Modifica el slice recibido directamente[cite: 4].
func sortAscending(values []int) {
	n := len(values)
	// Ciclo externo para recorrer todos los elementos.
	for i := 0; i < n-1; i++ {
		// Ciclo interno para comparar elementos adyacentes.
		for j := 0; j < n-i-1; j++ {
			// Si el elemento actual es mayor al siguiente, se intercambian.
			if values[j] > values[j+1] {
				temp := values[j]
				values[j] = values[j+1]
				values[j+1] = temp
			}
		}
	}
}

// Ejercicio 5: Dos valores que suman un objetivo.
// findPair busca dos valores en un slice ordenado que sumen el target[cite: 4].
func findPair(values []int, target int) bool {
	if len(values) < 2 {
		return false
	}

	// Se utilizan dos punteros (índices), uno en cada extremo del slice[cite: 4].
	left := 0
	right := len(values) - 1

	for left < right {
		suma := values[left] + values[right]

		if suma == target {
			fmt.Printf("Un par válido es (%d, %d)\n", values[left], values[right])
			return true
		} else if suma < target {
			// Si la suma es menor, se necesita un valor mayor (avanzar a la derecha).
			left++
		} else {
			// Si la suma es mayor, se necesita un valor menor (retroceder a la izquierda).
			right--
		}
	}

	fmt.Println("No existe")
	return false
}

// Ejercicio 6: Problema integrador 3Sum.
// threeSum encuentra todas las ternas únicas que sumen cero.
// Retorna la cantidad de ternas encontradas[cite: 5].
func threeSum(values []int) int {
	if len(values) < 3 {
		return 0
	}

	// Se crea y ordena una copia para no alterar el slice original[cite: 5].
	sorted := copyArray(values)
	sortAscending(sorted)

	count := 0

	// Se fija un elemento y se busca un par que sume el inverso de este elemento.
	for i := 0; i < len(sorted)-2; i++ {
		// Condición para evitar procesar el mismo valor inicial y generar ternas duplicadas.
		if i > 0 && sorted[i] == sorted[i-1] {
			continue
		}

		left := i + 1
		right := len(sorted) - 1

		for left < right {
			suma := sorted[i] + sorted[left] + sorted[right]

			if suma == 0 {
				fmt.Printf("(%d, %d, %d) ", sorted[i], sorted[left], sorted[right])
				count++

				// Saltamos elementos duplicados para el segundo y tercer valor de la terna.
				for left < right && sorted[left] == sorted[left+1] {
					left++
				}
				for left < right && sorted[right] == sorted[right-1] {
					right--
				}

				// Avanzamos los índices tras procesar la terna válida.
				left++
				right--
			} else if suma < 0 {
				left++
			} else {
				right--
			}
		}
	}

	if count == 0 {
		fmt.Print("Ninguna")
	}
	fmt.Println()
	return count
}

// Ejercicio 7: Máximos desde ambos extremos.
// buildLeftMax construye un slice con el máximo histórico de izquierda a derecha[cite: 5].
func buildLeftMax(height []int, leftMax []int) {
	if len(height) == 0 {
		return
	}
	max := height[0]
	for i := 0; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
		leftMax[i] = max
	}
}

// buildRightMax construye un slice con el máximo histórico de derecha a izquierda[cite: 5].
func buildRightMax(height []int, rightMax []int) {
	if len(height) == 0 {
		return
	}
	max := height[len(height)-1]
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > max {
			max = height[i]
		}
		rightMax[i] = max
	}
}

// Ejercicio 8: Trapping Rain Water.
// trappedWater calcula la cantidad de agua que queda atrapada entre las columnas.
// Retorna el total como un int64[cite: 5].
func trappedWater(height []int) int64 {
	n := len(height)
	if n == 0 {
		return 0
	}

	// Creación de slices auxiliares utilizando make[cite: 5].
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	// Llenado de los máximos utilizando las funciones del Ejercicio 7[cite: 5].
	buildLeftMax(height, leftMax)
	buildRightMax(height, rightMax)

	total := int64(0)

	// El agua atrapada en cada columna es el mínimo entre el muro más alto a su
	// izquierda y el muro más alto a su derecha, menos la altura de la columna actual[cite: 5].
	for i := 0; i < n; i++ {
		minMax := leftMax[i]
		if rightMax[i] < minMax {
			minMax = rightMax[i]
		}

		amount := minMax - height[i]
		if amount > 0 {
			total += int64(amount)
		}
	}
	return total
}

func main() {
	// Pruebas del Ejercicio 1[cite: 6]
	fmt.Println("--- Ejercicio 1 ---")
	a, b := 8, -3
	swap(&a, &b)
	fmt.Printf("a=%d, b=%d\n\n", a, b)

	// Pruebas del Ejercicio 2[cite: 6]
	fmt.Println("--- Ejercicio 2 ---")
	arrE2 := []int{4, -2, 0, 7, -5}
	stats(arrE2)
	fmt.Println("\n")

	// Pruebas del Ejercicio 3[cite: 6]
	fmt.Println("--- Ejercicio 3 ---")
	original := []int{4, -2, 0, 7, -5}
	copia := copyArray(original)
	copia[0] = 99
	fmt.Printf("Original: %v\nCopia: %v\n\n", original, copia)

	// Pruebas del Ejercicio 4[cite: 6]
	fmt.Println("--- Ejercicio 4 ---")
	copiaParaOrdenar := copyArray(original)
	sortAscending(copiaParaOrdenar)
	fmt.Printf("Ordenado: %v\n\n", copiaParaOrdenar)

	// Pruebas del Ejercicio 5[cite: 6]
	fmt.Println("--- Ejercicio 5 ---")
	findPair([]int{-4, -1, 0, 2, 5}, 1)
	findPair([]int{-4, -1, 0, 2, 5}, 8)
	findPair([]int{0, 0}, 0)
	findPair([]int{3}, 6)
	fmt.Println()

	// Pruebas del Ejercicio 6[cite: 6]
	fmt.Println("--- Ejercicio 6 (3Sum) ---")
	casos3Sum := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{0, 1, 1},
		{0, 0, 0},
		{0, 0, 0, 0},
		{-2, 0, 0, 2, 2},
	}
	for _, caso := range casos3Sum {
		fmt.Printf("Entrada %v: ", caso)
		threeSum(caso)
	}
	fmt.Println()

	// Pruebas del Ejercicio 7 y 8[cite: 7]
	fmt.Println("--- Ejercicio 7 y 8 (Agua Atrapada) ---")
	casosAgua := [][]int{
		{5}, {0, 1, 2, 3}, {3, 2, 1, 0}, {2, 2, 2},
		{3, 0, 3}, {3, 0, 2, 0, 4}, {4, 2, 0, 3, 2, 5},
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{2, 0, 2, 0, 2}, // Caso propio introducido para validación[cite: 7].
	}

	detalleAgua := []int{3, 0, 2, 0, 4}
	lMax := make([]int, len(detalleAgua))
	rMax := make([]int, len(detalleAgua))
	buildLeftMax(detalleAgua, lMax)
	buildRightMax(detalleAgua, rMax)
	fmt.Println("Detalle para:", detalleAgua)
	fmt.Println("LeftMax:", lMax)
	fmt.Println("RightMax:", rMax)

	for _, caso := range casosAgua {
		fmt.Printf("Alturas %v | Agua total: %d\n", caso, trappedWater(caso))
	}
}
