package main

import "fmt"

// Интерфейс для всех животных
type Animal interface {
	MakeSound() string
	Move() string
	GetInfo() string
	GetType() string
}

// Интерфейс для животных, которые умеют плавать
type Swimmer interface {
	Swim() string
	CanSwim() bool
}

// волк
type Wolf struct {
	Name string
}

func (f Wolf) MakeSound() string {
	return "ГАВ-ГАВ!"
}

func (f Wolf) Move() string {
	return "волк притаился"
}

func (f Wolf) GetInfo() string {
	return fmt.Sprintf("Волк %s - опасный и злой!", f.Name)
}

func (f Wolf) GetType() string {
	return "Млекопитающее"
}

func (f Wolf) Bark() string {
	return "Тяф-тяф-тяф!"
}

func (f Wolf) CanSwim() bool {
	return false
}

// Ежик
type Hedgehog struct {
	Name string
}

func (r Hedgehog) MakeSound() string {
	return "Фырк-Фырк!"
}

func (r Hedgehog) Move() string {
	return "Ежик сворачивается в клубок"
}

func (r Hedgehog) GetInfo() string {
	return fmt.Sprintf("Ежик %s - милый и добрый!", r.Name)
}

func (r Hedgehog) GetType() string {
	return "Млекопитающее"
}

func (r Hedgehog) Chatter() string {
	return "Фырк-фырк-фырк!"
}

func (r Hedgehog) CanSwim() bool {
	return true
}

func (r Hedgehog) Swim() string {
	return "Ежик плавает, чтобы перейти лужу"
}

// бегемот
type hippopotamus struct {
	Name string
}

func (e hippopotamus) MakeSound() string {
	return "Топ-топ-топ!"
}

func (e hippopotamus) Move() string {
	return "бегемот ходит медленно"
}

func (e hippopotamus) GetInfo() string {
	return fmt.Sprintf("Бегемот %s - один из самых больших наземных млекопитающий!", e.Name)
}

func (e hippopotamus) GetType() string {
	return "Млекопитающее"
}

func (e hippopotamus) Trumpet() string {
	return "большой рот!"
}

func (e hippopotamus) CanSwim() bool {
	return false
}

// жаба
type Frog struct {
	Name string
}

func (f Frog) MakeSound() string {
	return "Ква-ква!"
}

func (f Frog) Move() string {
	return "Лягушка прыгает"
}

func (f Frog) GetInfo() string {
	return fmt.Sprintf("Лягушка %s - земноводное!", f.Name)
}

func (f Frog) GetType() string {
	return "Земноводное"
}

func (f Frog) Ribbit() string {
	return "Ква-ква-ква!"
}

func (f Frog) CanSwim() bool {
	return true
}

func (f Frog) Swim() string {
	return "Лягушка плавает"
}

// кенгуру
type kangaroo struct {
	Name string
}

func (h kangaroo) MakeSound() string {
	return "прыг-скок!"
}

func (h kangaroo) Move() string {
	return "кенгуру быстро бегает"
}

func (h kangaroo) GetInfo() string {
	return fmt.Sprintf("кенгуру %s - пушистый и быстрый!", h.Name)
}

func (h kangaroo) GetType() string {
	return "Млекопитающее"
}

func (h kangaroo) Hop() string {
	return "кенгуру прыгает"
}

func (h kangaroo) CanSwim() bool {
	return false
}

func main() {
	lisa := Wolf{"Волк"}
	enot := Hedgehog{"Ежикт"}
	slon := hippopotamus{"бегемот"}
	lyagushka := Frog{"Лягушка"}
	zayac := kangaroo{"Заяц"}

	animals := []Animal{lisa, enot, slon, lyagushka, zayac}

	for _, animal := range animals {
		fmt.Println(animal.GetInfo())
		fmt.Println("Звук:", animal.MakeSound())
		fmt.Println("Движение:", animal.Move())
		fmt.Println("Тип:", animal.GetType())
		if swimmer, ok := animal.(Swimmer); ok {
			fmt.Println("Умеет плавать:", swimmer.CanSwim())
			fmt.Println("Способ плавания:", swimmer.Swim())
		}
		fmt.Println("----")
	}
}
