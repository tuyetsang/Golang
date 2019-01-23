package entities

type Student struct {
	Id, Name string
	Address  Address
}
type Address struct {
	Street, Ward string
}
