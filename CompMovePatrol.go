package main

/*
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"*/
//import	//("github.com/hajimehoshi/ebiten/v2")

type CompMovePatrol struct {
	entity   *Entity
	Velocity float64
	goals    []float64
}

/*
func NewCompCollider(enemy *Enemy, events []event) CompCollider{
	new := CompCollider{}
	new.enemy = enemy
	new.events = events
	return new
}*/

func NewCompMovePatrol(entity *Entity, data *dataArgs) {
	new := CompMovePatrol{}
	new.entity = entity
	new.goals = data.floats
	g.compMovePatrol = append(g.compMovePatrol, &new)
}

func (s *CompMovePatrol) MovePatrol() {

}