package memento

type GameMemento struct {
	money int
}

func (gm *GameMemento) getMoney() int {
	return gm.money
}

type Game struct {
	Money int
}

func (g *Game) CreateMemento() *GameMemento {
	return &GameMemento{g.Money}
}

func (g *Game) RestoreMemento(memento *GameMemento) {
	g.Money = memento.getMoney()
}
