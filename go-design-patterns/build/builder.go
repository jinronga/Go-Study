package build

import "errors"

type (
	Food struct {
		// 种类
		category string
		// 名称
		name string
		// 重量
		weight float64
		// 品牌
		brand string
		// 价格
		cost float64
	}

	FoodBuilder struct {
		Food
	}
)

func NewFood(category string, name string, weight float64, brand string, cost float64) *Food {
	return &Food{
		category: category,
		name:     name,
		weight:   weight,
		brand:    brand,
		cost:     cost,
	}
}

func NewFoodBuilder() *FoodBuilder {
	return &FoodBuilder{}
}

func (f *FoodBuilder) Category(category string) *FoodBuilder {
	f.category = category
	return f
}

func (f *FoodBuilder) Name(name string) *FoodBuilder {
	f.name = name
	return f
}

func (f *FoodBuilder) Weight(weight float64) *FoodBuilder {
	f.weight = weight
	return f
}

func (f *FoodBuilder) Brand(brand string) *FoodBuilder {
	f.brand = brand
	return f
}

func (f *FoodBuilder) Cost(cost float64) *FoodBuilder {
	f.cost = cost
	return f
}

func (f *FoodBuilder) Build() (*Food, error) {
	if f.category == "" {
		return nil, errors.New("miss type info")
	}
	if f.name == "" {
		return nil, errors.New("miss name info")
	}

	return &Food{
		category: f.category,
		name:     f.name,
		brand:    f.brand,
		weight:   f.weight,
		cost:     f.cost,
	}, nil
}
