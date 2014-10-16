package model

import (
	"fmt"
)

type Address struct {
	Name    string `json:"name"`
	Att     string `json:"att"`
	Street  string `json:"street"`
	Zip     int    `json:"zip"`
	State   string `json:"state"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func (a Address) Valid() bool {
	return true
}

func (a Address) String() string {

	addr := a.Name + "\n"

	if len(a.Att) > 0 {
		addr += fmt.Sprintf("Att: %s\n", a.Att)
	}

	addr += fmt.Sprintf("%s\n%v %s\n", a.Street, a.Zip, a.City)

	if len(a.State) > 0 {
		addr += a.State + "\n"
	}

	addr += a.Country

	return addr
}
