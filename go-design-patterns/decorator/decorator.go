package decorator

type (
	Food interface {
		// Eat 食用主食
		Eat() string
		// Cost 计话费
		Cost() float32
	}

	Rice struct {
	}

	Noodle struct {
	}
	Decorator Food

	LaoGanMaDecorator struct {
		Decorator
	}
	HamSausageDecorator struct {
		Decorator
	}
	FriedEggDecorator struct {
		Decorator
	}
)

func (n *Noodle) Eat() string {

	return "吃面条..."
}

func (n *Noodle) Cost() float32 {
	return 1.5
}

func NewRice() Food {
	return &Rice{}
}

func NewNoodle() Food {
	return &Noodle{}
}

func NewFriedEggDecorator(d Decorator) Decorator {
	return &FriedEggDecorator{
		Decorator: d,
	}
}

func NewDecorator(f Food) Decorator {
	return f
}

func NewLaoGanMaDecorator(d Decorator) Decorator {
	return &LaoGanMaDecorator{
		Decorator: d,
	}
}

func (l *LaoGanMaDecorator) Eat() string {
	// 加入老干妈配料
	return "加入一份老干妈~..." + l.Decorator.Eat()
}

func (l *LaoGanMaDecorator) Cost() float32 {
	// 价格增加 0.5 元
	return 0.5 + l.Decorator.Cost()
}

func NewHamSausageDecorator(d Decorator) Decorator {
	return &HamSausageDecorator{
		Decorator: d,
	}
}

func (h *HamSausageDecorator) Eat() string {
	// 加入火腿肠配料
	return "加入一份火腿~..." + h.Decorator.Eat()
}

func (h *HamSausageDecorator) Cost() float32 {
	// 价格增加 1.5 元
	return 1.5 + h.Decorator.Cost()
}

func (r *Rice) Eat() string {

	return "吃米饭..."
}

func (r *Rice) Cost() float32 {
	return 1
}

func (f *FriedEggDecorator) Eat() string {
	// 加入煎蛋配料
	return "加入一份煎蛋~..." + f.Decorator.Eat()
}

func (f *FriedEggDecorator) Cost() float32 {
	// 价格增加 1 元
	return 1 + f.Decorator.Cost()
}
