package lab7

import "fmt"

type Clothes struct {
	Description string
	Fabric string
	Brand string
	Price float64
}

func (f Clothes) GetDescription() {
	fmt.Println("Это", f.Description, "из", f.Fabric, "от бренда", f.Brand ,"стоимостью", f.Price)
}

func (f Clothes) GetPrice() float64 {
	return f.Price
}

func (f *Clothes) ApplyDiscount(x float64) {
	(*f).Price = (f.Price) * (x)
}

func (f *Clothes) SetNewPrice(x float64) {
	(*f).Price = x
}

func (f *Clothes) SetNewDescription(x string) {
	(*f).Fabric = x
}
