package animal

type Animal struct {
	Name      string
	ownerName string
}

func (a Animal) GetName() string {
	return a.Name

}
func (a Animal) GetOwner() string {
	return a.ownerName

}

func (a *Animal) SetName(an string, hn string) {
	a.Name = an
	a.ownerName = hn
}
