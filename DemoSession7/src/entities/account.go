package entities

import (
	"fmt"
)

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Status   bool   `json:"status"`
}

func (account Account) ToString() string {
	return fmt.Sprintf("Username:%s\nPassword:%s\nAge:%d\nStatus:%v", account.Username, account.Password, account.Age, account.Status)
}
