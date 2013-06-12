package barbershop

import "fmt"
import "time"

type Barber struct {
	Human               // Barber is composed as an Human
	WakeMe chan *Client // and can be waked by a client
}

// barber manage waitin room of the shop
func (b *Barber) ManageShop(shop *Shop) {
	for {
		select {
		// try to get a client if any in waiting room
		case c := <-shop.WaitingRoom:
			// ok there was at least one client
			fmt.Printf("%s cuts %s's hair\n", b.GetName(), c.GetName())
			time.Sleep(time.Millisecond * 6)
			fmt.Printf("%s finished %s\n", b.GetName(), c.GetName())

			// pas de client,on fait une action par défaut
		default:
			fmt.Printf("%s want to sleep... zzZZZzzz\n", b.GetName())
			c := <-b.WakeMe // not a "select case" so it blocks there
			fmt.Printf("%s waked %s\n", c.GetName(), b.GetName())
		}
	}
}

// Return a new barber
func NewBarber(name string) *Barber {
	b := &Barber{}
	b.name = name
	b.WakeMe = make(chan *Client)
	return b
}
