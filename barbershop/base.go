package barbershop

// Base structure for humans (barber and client)
type Human struct {
	name string // minuscule, so not exported outside package
}

// Return the human name
func (h *Human) GetName() string {
	return h.name
}
