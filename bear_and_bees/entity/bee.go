package entity

import (
	"bear_and_bees/honey_comb"
	"sync"
	"time"
)

// Representacion de la entidad Abeja
type Bee struct {
	Id            int             // identificador de la abeja
	M             *sync.Mutex     // lock and unlock process
	BearCanEat    *sync.WaitGroup // Espera del oso
	BeeCanProduce *sync.WaitGroup // Espera de las abajas
	Delay         int             // intervalo entre trabajos
}

func (b *Bee) Produce() {
	var isFull bool
	for {
		b.BeeCanProduce.Wait()             // Espera a que pueda producir
		b.M.Lock()                         // bloqua el proceso
		isFull = honey_comb.AddHoney(b.Id) // Agrega miel
		if isFull {
			b.BearCanEat.Done()            // si esta lleno le avisa al oso
			b.BeeCanProduce.Add(1)         // se pone en espera y todas las abejas
			honey_comb.SetIsFilling(false) // los procesos pendientes ya no pueden depositar miel
		}
		b.M.Unlock()                                          // desbloquea el proceso
		time.Sleep(time.Duration(b.Delay) * time.Millisecond) // sleep

	}

}
