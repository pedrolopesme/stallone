package api

import "fmt"

type Api struct{}

func NewAPI() *Api {
	return &Api{}
}

func (a *Api) Run() {
	fmt.Println("You, me or nobody is gonna hit as hard as life.")
}
