package main

/*
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"*/
import	"image"

type CompCollider struct {
	enemy *Enemy
	events []event
	Velocity float64
}

func NewCompCollider(enemy *Enemy, events []event) CompCollider{
	new := CompCollider{}
	new.enemy = enemy
	new.events = events
	return new
}

func (s *CompCollider) collide(rectPlayer image.Rectangle) {
	
	rectEnemy := image.Rect(int(s.enemy.xpos), int(s.enemy.ypos), int(s.enemy.xpos) + 30, int(s.enemy.ypos) + 30)
	//rectPlayer := image.Rect(int(player.xpos), int(player.ypos), int(player.xpos) + 30, int(player.ypos) + 30)
	if rectEnemy.Overlaps(rectPlayer) {
		for _,z := range s.events {
			z.execute(s.enemy)
		}
	}
}

