package main

import "bear_and_bees/worker"

/*
	Autor: Angel Zarate
	Fecha: 03 Febrero 2022
	Problema: El Oso y las abejas
	Descripcion:
		El oso y las abejas.
		Existen n abejas y un hambriento oso que comparten un tarro de miel.
		Inicialmente el tarro de miel esta vacio,
		su capacidad es M porciones de miel (M << N).
		El oso duerme hasta que el tarro de miel se llene,
		entonces se come toda la miel y vuelve a dormir.
		Cada abeja produce una porcion de miel que coloca en el tarro,
		la abeja que llena el tarro de miel despierta al oso.
*/

func main() {
	// tamaño del panal de abejas
	size_honey_comb := 20
	// numero de abejas en la colmena
	total_bees := 5
	// intervalo de tiempo en el que duermen las abejas
	delay := 500
	// Ecosistema del problema
	w := worker.NewEcosystem(size_honey_comb, total_bees, delay)
	// se utiliza para que no finalice el programacion
	// Si desea detener el programa presionar Ctl + C
	result := make(chan int)
	result <- w.Run()
}
