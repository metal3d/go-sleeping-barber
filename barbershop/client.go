package barbershop

import "fmt"
import "time"

// AFAIK a client is an Human
type Client struct {
	Human
}

// Client enters shop and tries to find a seat
func (c *Client) EnterShop(shop *Shop) {
	// try to find a seat 3 times...
	for i := 0; i < 3; i++ {
		if i > 0 {
			// client is back
			fmt.Printf(">> %s is back\n", c.GetName())
		}
		select {
		case shop.WaitingRoom <- c:
			// yes ! client found a seat
			fmt.Printf(":) %s found a seat\n", c.GetName())
			select {
			// we check that the barber is not sleeping
			case shop.Barbers.WakeMe <- c:
				// lazy barber ! wake him up !
				fmt.Printf("WAKE UP ! screams %s\n", c.GetName())
			default:
				// default... do nothing (default is necessary there)
			}
			return // stop the problem, client is now managed by barber

		default:
			// no seat (write on channel failed)
			fmt.Printf("... %s will be back later ...\n", c.GetName())
			time.Sleep(time.Millisecond * 100)
		}
	}
	// no seat, get out
	fmt.Printf(":( %s didn't found a seat, get out...\n", c.GetName())
}

// create a new client (pointer)
func NewClient(name string) *Client {
	c := new(Client) // returns a pointer, equiv. &Client{}
	c.name = name
	return c
}
