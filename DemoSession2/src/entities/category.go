package entities

type Category struct {
	Id, Name string
	Products []Product
}

func (category Category) ToString() (result string) {
	result = "Id:" + category.Id
	result += "\nName:" + category.Name
	for _, product := range category.Products {
		result += product.ToString()
	}
	return
}
func (category Category) Total() (result float64) {
	result = 0
	for _, product := range category.Products {
		result += product.Total()
	}
	return
}
func (category Category) Max() (m Product) {
	m = category.Products[0]
	for _, product := range category.Products {
		if m.Price < product.Price {
			m = product
		}
	}
	return
}
