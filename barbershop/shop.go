package barbershop

// a shop has got barber and a waiting room
type Shop struct {
	Barbers     *Barber
	WaitingRoom chan *Client
}

// create pointer on shop
func NewShop(barber *Barber) *Shop {
	shop := new(Shop)
	shop.Barbers = barber
	// we create 4 seats for clients in the waiting room
	shop.WaitingRoom = make(chan *Client, 4)
	return shop
}
