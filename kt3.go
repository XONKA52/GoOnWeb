package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gen2brain/beeep"
)

type Giraffe struct {
	Height int
	NeckLength int
	Size  int
}

type Puma struct {
	HighSpeed int
	ClimbTree bool
	Stealth   bool
}

type Kangaroo struct {
	Size       int
	JumpHeight int
	PouchCapacity int
}

func processGiraffe(wg *sync.WaitGroup) {
	defer wg.Done()

	giraffe := Giraffe{Height: 500, NeckLength: 200, Size: 800}
	fmt.Printf("Жираф - Высота: %d см, Длина шеи: %d см, Размер: %d кг\n", giraffe.Height, giraffe.NeckLength, giraffe.Size)

	go func() {
		beeep.Notify("Жираф", "Обработана информация о жирафе", "")
	}()
	time.Sleep(1 * time.Second) 
}

func processPuma(wg *sync.WaitGroup) {
	defer wg.Done()

	puma := Puma{HighSpeed: 80, ClimbTree: true, Stealth: true}
	fmt.Printf("Пума - Скорость: %d км/ч, Залезает на деревья: %v, Тихая: %v\n", puma.HighSpeed, puma.ClimbTree, puma.Stealth)

	go func() {
		beeep.Notify("Пума", "Обработана информация о пуме", "")
	}()
	time.Sleep(1 * time.Second) 
}

func processKangaroo(wg *sync.WaitGroup) {
	defer wg.Done()

	kangaroo := Kangaroo{Size: 90, JumpHeight: 300, PouchCapacity: 5}
	fmt.Printf("Кенгуру - Размер: %d кг, Высота прыжка: %d см, Вместимость сумки: %d кг\n", kangaroo.Size, kangaroo.JumpHeight, kangaroo.PouchCapacity)

	go func() {
		beeep.Notify("Кенгуру", "Обработана информация о кенгуру", "")
	}()
	time.Sleep(1 * time.Second) 
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go processGiraffe(&wg)
	go processPuma(&wg)
	go processKangaroo(&wg)

	wg.Wait()
	fmt.Println("Все животные обработаны.")
}
