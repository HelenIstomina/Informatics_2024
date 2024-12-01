package lab7

import "fmt"

type HouseholdGoods struct {
	Name      string
	Price     float64
}

func (f HouseholdGoods) GetDescription() {
	fmt.Println("Микроволновая печь", f.Name, "стоимостью", f.Price)
}

func (f HouseholdGoods) GetPrice() float64 {
	return f.Price
}

func (f *HouseholdGoods) ApplyDiscount(x float64) {
	(*f).Price = (f.Price) * (x)
}

func (f *HouseholdGoods) SetNewPrice(x float64) {
	(*f).Price = x
}

func (f *HouseholdGoods) SetNewDescription(x string) {
	(*f).Name = x
}
