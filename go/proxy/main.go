package main

import (
	"fmt"
	"time"

	"gitlab.com/chicho69-cesar/proxy/local"
)

var loc local.Proxy

func main() {
	loc = local.New()

	GetByID(2)
	GetByID(2)
	GetByID(1)
	GetByID(2)
	GetByID(3)
	GetByID(2)
	GetByID(1)
	GetAll()
}

func GetByID(ID uint) {
	start := time.Now()
	fmt.Printf("%+v", loc.GetByID(ID))
	elapsed := time.Since(start)
	fmt.Printf("Tiempo usado: %v\n", elapsed)
}

func GetAll() {
	start := time.Now()
	fmt.Printf("%+v", loc.GetAll())
	elapsed := time.Since(start)
	fmt.Printf("Tiempo usado: %v\n", elapsed)
}
