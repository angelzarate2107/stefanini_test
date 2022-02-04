package honey_comb

import (
	"fmt"
	"sync"
)

var once sync.Once

type honey_comb struct {
	isFilling bool   // se evita que una abeja agrega miel mientras el oso come
	index     int    // indicador del nivel de miel
	size      int    // tamaño del contenedor de miel
	container []bool // inicializa con falso el array
}

var instance *honey_comb

/*
	Coloca el tarro en modo llenado o vaciado
*/
func SetIsFilling(isFill bool) {
	instance.isFilling = isFill

}

/* Instance
Crea un singleton que representa el tarro de miel
*/
func Instance(size int) *honey_comb {
	once.Do(func() {
		instance = &honey_comb{
			isFilling: true,
			size:      size,
			index:     0,
			container: make([]bool, size),
		}
	})
	return instance
}

/* AddHoney
Cada Abeja deposita la miel en el contenedor
	@return bool el cual indica que el contenedor esta lleno
*/
func AddHoney(bee_id int) bool {
	if instance.index < instance.size && instance.isFilling {
		instance.container[instance.index] = true
		instance.index++
		// que Id de abeja puso la miel
		fmt.Printf("bee %d ", bee_id)
		printHoneyComb()
		return instance.index == instance.size
	}
	return false

}

/*
	MinusHoney
	Decrementa la cantidad del miel, la cual el oso esta comiendo
*/
func MinusHoney() bool {
	instance.index--
	instance.container[instance.index] = false
	printHoneyComb()
	return instance.index == 0
}

/*
	printHoneyComb
	imprime una representacion del tarro de miel

*/
func printHoneyComb() {
	fmt.Print("[")
	for _, v := range instance.container {
		if v {
			fmt.Print("|")
		} else {
			fmt.Print("_")
		}
	}
	fmt.Println("]")
}
