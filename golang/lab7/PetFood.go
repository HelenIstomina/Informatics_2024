package lab7

import "fmt"

type PetFood struct {
	Name      string
	Brand     string
	Price     float64
}

func (f PetFood) GetDescription() {
	fmt.Println("Корм для кошек", f.Name, "от бренда", f.Brand, "он стоит", f.Price)
}

func (f PetFood) GetPrice() float64 {
	return f.Price
}

func (f *PetFood) ApplyDiscount(x float64) {
	(*f).Price = (f.Price) * (x)
}

func (f *PetFood) SetNewPrice(x float64) {
	(*f).Price = x
}

func (f *PetFood) SetNewDescription(x string) {
	(*f).Name = x
}
