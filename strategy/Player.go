package main

import (
	"fmt"
)

type Player struct {
	weapon IDoDamage
	name   string
	health int
}

type Gun struct {
	damage int
}

func (g Gun) DoDamage(enemy *Player) {
	g.damage = 10
	enemy.health -= g.damage
	check(*enemy, 1)
}

type Fork struct {
	damage int
}

func (f Fork) DoDamage(enemy *Player) {
	f.damage = 4
	enemy.health -= f.damage
	check(*enemy, 2)
}

type Knife struct {
	damage int
}

func (k Knife) DoDamage(enemy *Player) {
	k.damage = 5
	enemy.health -= k.damage
	check(*enemy, 3)
}
func check(enemy Player, q_type int) {
	if enemy.health <= 0 {
		fmt.Println(enemy.name, "is died")
		SayQuote(q_type)
	} else {
	}
}
func SayQuote(i int) {
	switch i {
	case 1:
		fmt.Println("One shot one opportunity")

	case 2:
		fmt.Println("Не бойся ножа, а бойся вилки: один удар – 4 дырки")
	case 3:
		fmt.Println("Нож в печень и ты не вечен")
	default:
		fmt.Println("")
	}
}
