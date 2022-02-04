package worker

import (
	"beer_and_bees/entity"
	"beer_and_bees/honey_comb"
	"sync"
)

/*
	BearBeesWorker: Representa el ecosistema del problema
*/
type BearBeesWorker struct {
	bees          []*entity.Bee   // colmena de abejas
	bear          *entity.Bear    // oso
	mutex         *sync.Mutex     // se utiliza en los bloqueos
	bearCanEat    *sync.WaitGroup // Espera del oso
	beeCanProduce *sync.WaitGroup // Espera de las abejas
}

/*
	NewEcosystem.- Constructor del ecosistema
	sizeHoneyComb: tamaño del tarro de miel
	totalBees: total de abejas
	delay: intervalo entre cada trabajo
*/
func NewEcosystem(sizeHoneyComb, totalBees int, delay int) *BearBeesWorker {

	honey_comb.Instance(sizeHoneyComb)

	instance := &BearBeesWorker{
		mutex:         &sync.Mutex{},
		bearCanEat:    &sync.WaitGroup{},
		beeCanProduce: &sync.WaitGroup{},
		bees:          make([]*entity.Bee, totalBees),
	}
	instance.bearCanEat.Add(1)
	//
	instance.bear = &entity.Bear{
		BearCanEat:    instance.bearCanEat,
		BeeCanProduce: instance.beeCanProduce,
		Delay:         delay,
	}

	for i := 0; i < totalBees; i++ {
		instance.bees[i] = &entity.Bee{
			Id:            i,
			M:             instance.mutex,
			BearCanEat:    instance.bearCanEat,
			BeeCanProduce: instance.beeCanProduce,
			Delay:         delay,
		}
	}

	return instance
}

/*
	Run: ejecucion del ecosistema
	Invoca a las abejas y oso a trabajar
*/
func (w *BearBeesWorker) Run() int {
	go w.bear.Eat()

	for _, b := range w.bees {
		go b.Produce()
	}
	return 0
}
