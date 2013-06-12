package main

import (
	"./barbershop"
	"fmt"
	"time"
)

func main() {
	// create a barber
	barber := barbershop.NewBarber("John")

	// and create a shop, a barber will manage thid
	// we give 4 seats to the shop
	shop := barbershop.NewShop(barber, 4)

	// let's barber manage the shop... and continue
	go barber.ManageShop(shop)

	// because clients arrive later :)
	time.Sleep(time.Second * 2)

	// let's make a list of client's name
	clients := []string{"Alain"}
	clients = append(clients, "Bernard")
	clients = append(clients, "Claude")
	clients = append(clients, "Dan")
	clients = append(clients, "Eric")
	clients = append(clients, "Franck")
	clients = append(clients, "Gérard")
	clients = append(clients, "Henri")
	clients = append(clients, "Ignasse")

	// create clients
	for _, c := range clients {
		client := barbershop.NewClient(c)
		// let them enter the shop, but use "go"
		// to create other without waiting the end
		// of the "entrance"
		go client.EnterShop(shop)

	}

	// press a key to finish
	fmt.Scanln()

}
