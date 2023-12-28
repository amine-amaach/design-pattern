package main

import (
	"fmt"

	G "github.com/amine-amaach/design-pattern/factory/contracts"
	F "github.com/amine-amaach/design-pattern/factory/factory"
)

type Client struct {
	g G.IGun
}

func main() {
	// * Client A wants an AK47
	gun, err := F.NewGun("ak47")
	if err != nil {
		fmt.Println(err)
		return
	}

	clientA := Client{
		g: gun,
	}
	fmt.Println("Client A has", clientA.g.GetName(), "with power", clientA.g.GetPower())

	// * Client B wants a musket
	gun, err = F.NewGun("musket")
	if err != nil {
		fmt.Println(err)
		return
	}

	clientB := Client{
		g: gun,
	}
	fmt.Println("Client B has", clientB.g.GetName(), "with power", clientB.g.GetPower())

	// * Client C wants a bazooka
	gun, err = F.NewGun("bazooka")
	if err != nil {
		fmt.Println(err)
		return
	}

	clientC := Client{
		g: gun,
	}
	fmt.Println("Client C has", clientC.g.GetName(), "with power", clientC.g.GetPower())
}
