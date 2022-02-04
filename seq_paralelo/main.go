package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//crea un arreglo con n elementos
	size_array := 100000000
	// cantidad de sub rutinas
	sub_procesos := 4
	// creacion del arreglo
	array := make([]int, size_array)
	fmt.Printf("Suma Secuencial vs Paralelo (tamaño del arreglo: %d)\nCantidad de subprocesos: %d\n\n", size_array, sub_procesos)
	rand.Seed(time.Now().UnixNano())
	// llena el arreglo con numero aleatorios del 1 al 10
	for i := range array {
		array[i] = (rand.Int() % 10) + 1
	}
	// lama al metodo secuencial
	secuencial(array)
	// lama el metodo paralelo o con go rutinas
	paralelo(array, sub_procesos)
}

func secuencial(array []int) int64 {
	fmt.Println("Metodo Secuencial")
	// inicio del temporizador
	t := time.Now()
	var total int64 = 0
	// recorre todo el arreglo, iterando cada posicion y sumando
	for _, v := range array {
		total += int64(v)
	}
	// imprime el resultado final y el tiempo que le tomo en milisegundos
	fmt.Printf("\ttotal de la suma: %d \n\ttiempo transcurrido: %d ms\n", total, time.Since(t).Milliseconds())
	return total
}

/*
	Metodo concurrente o paralelo (dependiendo del hw)

	Divide el arreglo principal en arreglos mas pequeños,
	y este es enviado a procesos para que calcule la suma de los arrays
	una vez que todos terminen de sumar el proceso principal recolecta los
	resultados y los suma
*/
func paralelo(array []int, n_process int) int64 {
	fmt.Println("Metodo Concurrente / Paralelo")
	t := time.Now()
	var total int64
	var size int = int(len(array) / n_process) // tamaño prom de los mini arrays
	var inicio int
	var fin int
	// canal en donde se enviaran los resultados de cada proceso
	c_acum := make(chan int)
	for i := 0; i < n_process; i++ {
		// calcula el inicio y fin de cada arreglo
		inicio = (i * size)
		fin = inicio + size
		// si es el ultimo le tocan un poco mas de items
		if i == n_process-1 {
			go sum_subarray(array[inicio:], c_acum)
		} else {
			go sum_subarray(array[inicio:fin], c_acum)
		}

	}
	// obtiene los resultados mediante los canaels
	for i := 0; i < n_process; i++ {
		total += int64(<-c_acum)
	}
	fmt.Printf("\ttotal de la suma: %d \n\ttiempo transcurrido: %d ms \n", total, time.Since(t).Milliseconds())
	return total
}

func sum_subarray(arr []int, c chan int) {
	var total int = 0
	// itera un arreglo mas pequeño
	for _, v := range arr {
		total += v
	}
	// envia la informacion de regreso
	c <- total

}
