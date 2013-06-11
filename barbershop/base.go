package barbershop

// Base structure for humans (barber and client)
type Human struct {
    name string // minuscule, donc non exportée hors package
}

// Return the human name
func (h *Human) GetName() string {
    return h.name
}

