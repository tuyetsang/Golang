package entities

import(
)

type Category struct{
	Id,Name string
	Products []Product
}

func (category Category) ToString() (result string){
	result = "Id:" + category.Id
	result += "\nName:" + category.Name
	return
}