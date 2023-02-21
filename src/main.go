package main

import (
	"fmt"

	"github.com/domin191013/Go-Hitbtc-With-Cache/config"
	"github.com/domin191013/Go-Hitbtc-With-Cache/handlers"
)

func main() {
	fmt.Println("Server starts at port :", config.GetPort())
	handlers.Start()
}
