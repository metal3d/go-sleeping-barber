package barbershop

// a shop has got barber and a waiting room
type Shop struct {
	Barber     *Barber
	WaitingRoom chan *Client
}

// create pointer on shop
func NewShop(barber *Barber, seats int) *Shop {
	shop := new(Shop)
	shop.Barber = barber
	// we create 4 seats for clients in the waiting room
	shop.WaitingRoom = make(chan *Client, seats)
	return shop
}
