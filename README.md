# Taller de Programación Imperativa en Go

**Integrantes:** Ian Alejandro Bonilla Mena
**Lenguaje:** Go
**Curso:** Paradigmas de Programación

## Instrucciones de Ejecución
El proyecto está contenido en un único archivo fuente. Para compilar y ejecutar los casos de prueba, abra la terminal en el directorio raíz del proyecto y ejecute el siguiente comando:
`go run main.go`

## Resultados de las Pruebas
El programa ejecuta exitosamente la totalidad de los escenarios base estipulados en el documento del taller.
* Las métricas y recorridos del Ejercicio 2 muestran los datos correctos para `[4, -2, 0, 7, -5]`.
* Las validaciones de los Ejercicios 5 y 6 descartan los casos sin soluciones válidas (ej. tamaños menores al necesario).
* El Ejercicio 8 procesa correctamente los 8 escenarios provistos en el documento y un escenario de validación adicional, gestionando la complejidad espacial y temporal en $O(n)$.

## Respuestas Teóricas

### Ejercicio 1: Punteros y Alcance
* **Por qué la función recibe punteros:** La función necesita modificar el estado de las variables originales. Si Go pasara las variables por valor, se enviarían copias independientes a la función, y las asignaciones se perderían al finalizar la rutina. Pasar direcciones de memoria (`*int`) permite alterar los datos subyacentes.
* **Diferencia entre `a` y `*a`:** La variable `a` es el puntero que contiene la dirección de memoria, mientras que `*a` accede y manipula el valor numérico almacenado en dicha dirección.
* **Alcance de variable local:** La variable `aux` definida en `swap` posee un alcance local. Solo existe dentro del bloque de la función `swap` y es destruida de la memoria una vez que la ejecución retorna a `main`.

### Ejercicio 3: Copia Independiente
* **¿Por qué `other := original` comparte elementos?** En Go, un descriptor de *slice* consta de un puntero al arreglo subyacente, un contador de longitud y un contador de capacidad. Al utilizar `:=`, Go realiza una copia superficial del descriptor, por lo que ambas variables referencian el mismo espacio de memoria en el arreglo. Para crear copias independientes, es necesario utilizar `make` y copiar los datos explícitamente.

### Ejercicio 5: Búsqueda de Pares
* **¿Por qué se descarta un extremo?** El algoritmo utiliza un arreglo ordenado. Si la suma de los extremos es mayor al objetivo, es matemáticamente imposible que el elemento de la derecha forme un par válido con cualquier otro número a su izquierda, ya que todos sumarán de más. Por ello, se descarta el índice derecho. La misma lógica aplica para descartar el izquierdo si la suma es menor.

### Ejercicio 6: 3Sum - Evitar Duplicados
* **Mecanismo:** Primero, el arreglo base se ordena. Durante la iteración principal, si el elemento actual es idéntico al anterior `(sorted[i] == sorted[i-1])`, se omite la iteración. De manera similar, cuando se encuentra una terna válida en el ciclo interno, los punteros `left` y `right` avanzan en un bucle para saltar cualquier número idéntico consecutivo. Esto garantiza que las ternas devueltas sean únicas.

### Ejercicio 8: Trapping Rain Water
* **Dependencia del menor máximo:** Para que el agua se acumule en una columna, necesita paredes en ambos flancos. El nivel del agua retenida está limitado por el muro más bajo entre el máximo histórico izquierdo y el máximo histórico derecho. Cualquier excedente de agua rebosará por ese extremo inferior.
* **Caso Propio Justificado:** Se introdujo el arreglo `[2, 0, 2, 0, 2]`. Al existir tres columnas de altura dos separadas por depresiones de altura cero, el agua queda atrapada en ambas depresiones hasta un límite de dos unidades cada una. El resultado esperado validado por el programa es de 4 unidades de volumen de agua.

### Comparación del Manejo de Memoria: C vs Go
* **Lenguaje C:** Utiliza gestión de memoria manual. El programador emplea `malloc` para reservar memoria y está obligado a liberarla utilizando `free` para evitar fugas de memoria. 
* **Lenguaje Go:** Cuenta con un recolector de basura (*Garbage Collector*) automático. Al utilizar `make`, el motor rastrea las referencias. Cuando una variable sale de alcance y ningún puntero la referencia, el recolector recupera la memoria automáticamente sin necesidad de llamar a funciones de liberación de memoria.