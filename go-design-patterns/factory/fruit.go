package factory

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Fruit interface {
	Eat()
}

type Orange struct {
	name string
}

func (o *Orange) Eat() {
	//TODO implement me
	panic("implement me")
}

func NewOrange(name string) Fruit {
	return &Orange{
		name: name,
	}
}

type Strawberry struct {
	name string
}

func NewStrawberry(name string) Fruit {
	return &Strawberry{
		name: name,
	}
}

func (s *Strawberry) Eat() {
	fmt.Printf("i am strawberry: %s, i am about to be eaten...", s.name)
}

type Cherry struct {
	name string
}

func NewCherry(name string) Fruit {
	return &Cherry{
		name: name,
	}
}

func (c *Cherry) Eat() {
	fmt.Printf("i am cherry: %s, i am about to be eaten...", c.name)
}

type FruitFactory struct {
	creators map[string]fruitCreator
}

type fruitCreator func(name string) Fruit

func NewFruitFactory() *FruitFactory {
	return &FruitFactory{
		creators: map[string]fruitCreator{
			"orange":     NewOrange,
			"strawberry": NewStrawberry,
			"cherry":     NewCherry,
		},
	}
}

func (f *FruitFactory) CreateFruit(typ string) (Fruit, error) {
	fruitCreator, ok := f.creators[typ]
	if !ok {
		return nil, fmt.Errorf("fruit typ: %s is not supported yet", typ)
	}

	src := rand.NewSource(time.Now().UnixNano())
	rander := rand.New(src)
	name := strconv.Itoa(rander.Int())
	return fruitCreator(name), nil
}
