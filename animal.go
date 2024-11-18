package animal

type Animal interface {
	Eat() string
	Sound() string
	Move() string
	Age() int
}

type Giraffe struct {
	AgeValue int
}

func (g *Giraffe) Sound() string {
	return "Жираф фыркает"
}

func (g *Giraffe) Move() string {
	return "Жираф шагает"
}

func (g *Giraffe) Age() int {
	return g.AgeValue
}

func (g *Giraffe) Eat() string {
	return "Жираф ест листья и траву"
}

type Cougar struct {
	AgeValue int
}

func (c *Cougar) Sound() string {
	return "Пума рычит"
}

func (c *Cougar) Move() string {
	return "Пума шагает"
}

func (c *Cougar) Age() int {
	return c.AgeValue
}

func (c *Cougar) Eat() string {
	return "Пума ест других животных"
}

type Kangaroo struct {
	AgeValue int
}

func (k *Kangaroo) Sound() string {
	return "Кенгуру рычит и кричит"
}

func (k *Kangaroo) Move() string {
	return "Кенгуру прыгает"
}

func (k *Kangaroo) Age() int {
	return k.AgeValue
}

func (k *Kangaroo) Eat() string {
	return "Кенгуру ест траву и фрукты"
}
