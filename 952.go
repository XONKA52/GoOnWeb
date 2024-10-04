package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Интерфейс для всех животных
type creature interface {
	MakeSound() string
	Move() string
	GetInfo() string
	GetType() string
}

// Интерфейс для животных, которые умеют плавать
type Diver interface {
	Swim() string
	CanSwim() bool
}

// Пингвин
type Penguin struct {
	Name string
}

func (p Penguin) MakeSound() string {
	return "Кря-кря!"
}

func (p Penguin) Move() string {
	return "Пингвин шлепает по снегу"
}

func (p Penguin) GetInfo() string {
	return fmt.Sprintf("Пингвин %s - холодолюбивое животное!", p.Name)
}

func (p Penguin) GetType() string {
	return "Птица"
}

func (p Penguin) CanSwim() bool {
	return true
}

func (p Penguin) Swim() string {
	return "Пингвин ныряет в ледяную воду"
}

// Белка
type Squirrel struct {
	Name string
}

func (s Squirrel) MakeSound() string {
	return "Цик-цик!"
}

func (s Squirrel) Move() string {
	return "Белка прыгает по деревьям"
}

func (s Squirrel) GetInfo() string {
	return fmt.Sprintf("Белка %s - маленький пушистый зверёк!", s.Name)
}

func (s Squirrel) GetType() string {
	return "Млекопитающее"
}

func (s Squirrel) CanSwim() bool {
	return false
}

// Медведь
type Bear struct {
	Name string
}

func (b Bear) MakeSound() string {
	return "Рррр!"
}

func (b Bear) Move() string {
	return "Медведь шагает тяжело"
}

func (b Bear) GetInfo() string {
	return fmt.Sprintf("Медведь %s - крупное и сильное животное!", b.Name)
}

func (b Bear) GetType() string {
	return "Млекопитающее"
}

func (b Bear) CanSwim() bool {
	return true
}

func (b Bear) Swim() string {
	return "Медведь плывёт через реку"
}

// Черепаха
type Turtle struct {
	Name string
}

func (t Turtle) MakeSound() string {
	return "Ш-ш-ш!"
}

func (t Turtle) Move() string {
	return "Черепаха медленно ползет"
}

func (t Turtle) GetInfo() string {
	return fmt.Sprintf("Черепаха %s - древнее и долго живущее существо!", t.Name)
}

func (t Turtle) GetType() string {
	return "Пресмыкающееся"
}

func (t Turtle) CanSwim() bool {
	return true
}

func (t Turtle) Swim() string {
	return "Черепаха плывёт по воде"
}

// Дельфин
type Dolphin struct {
	Name string
}

func (d Dolphin) MakeSound() string {
	return "Ии-ии!"
}

func (d Dolphin) Move() string {
	return "Дельфин быстро плавает"
}

func (d Dolphin) GetInfo() string {
	return fmt.Sprintf("Дельфин %s - одно из самых умных существ в океане!", d.Name)
}

func (d Dolphin) GetType() string {
	return "Млекопитающее"
}

func (d Dolphin) CanSwim() bool {
	return true
}

func (d Dolphin) Swim() string {
	return "Дельфин резво скользит в воде"
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите тип животного (пингвин, белка, медведь, черепаха, дельфин):")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		var animal Animal

		switch input {
		case "пингвин":
			fmt.Println("Введите имя пингвина:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			animal = Penguin{Name: name}
		case "белка":
			fmt.Println("Введите имя белки:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			animal = Squirrel{Name: name}
		case "медведь":
			fmt.Println("Введите имя медведя:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			animal = Bear{Name: name}
		case "черепаха":
			fmt.Println("Введите имя черепахи:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			animal = Turtle{Name: name}
		case "дельфин":
			fmt.Println("Введите имя дельфина:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			animal = Dolphin{Name: name}
		default:
			fmt.Println("Неверный тип животного. Попробуйте снова.")
			continue
		}

		// Вывод информации о животном
		fmt.Println(animal.GetInfo())

		// Вывод звука
		fmt.Println("Звук:", animal.MakeSound())

		// Вывод движения
		fmt.Println("Движение:", animal.Move())

		// Вывод типа
		fmt.Println("Тип:", animal.GetType())

		// Проверка умения плавать
		if swimmer, ok := animal.(Swimmer); ok {
			fmt.Println("Умеет плавать:", swimmer.
				CanSwim())
			if swimmer.CanSwim() {
				fmt.Println("Способ плавания:", swimmer.Swim())
			}
		}

		fmt.Println("----")
	}
}
