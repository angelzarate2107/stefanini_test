package entity

import (
	"bear_and_bees/honey_comb"
	"fmt"
	"sync"
	"time"
)

// Bear .- reporesentación del oso
type Bear struct {
	BearCanEat    *sync.WaitGroup //
	BeeCanProduce *sync.WaitGroup //
	Delay         int
}

// Eat.- Trabajo del oso
func (bear *Bear) Eat() int {
	var isEmpty bool
	for {
		bear.BearCanEat.Wait() // espera a que se llene el tarro de miel
		fmt.Print("bear eat: ")
		isEmpty = honey_comb.MinusHoney()
		if isEmpty { // si el tarro esta vacio
			honey_comb.SetIsFilling(true) // set el tarro en modo llenado
			bear.BeeCanProduce.Done()     // las abejas ya pueden producir
			bear.BearCanEat.Add(1)        // el oso estara en espera de volver a comer
		}
		time.Sleep(time.Duration(bear.Delay) * time.Millisecond) // duerme el oso
	}
}
