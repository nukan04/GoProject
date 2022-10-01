package main

import "fmt"

func main() {
	fmt.Println("Hello Tlezuhan")
	Nurdaulet := Player{Fork{}, "Nurdaulet", 10}
	Tleuzhan := Player{Fork{}, "Tleuzhan", 10}
	Ruslan := Player{Fork{}, "Ruslan", 10}
	Syndar := Player{Fork{}, "Syndar", 10}
	//-----
	Nurdaulet.weapon.DoDamage(&Tleuzhan)
	fmt.Println("Tleuzhan health is:", Tleuzhan.health)
	Nurdaulet.weapon.DoDamage(&Tleuzhan)
	fmt.Println("Tleuzhan health is:", Tleuzhan.health)
	Nurdaulet.weapon.DoDamage(&Tleuzhan)
	//-----
	Nurdaulet.weapon = Gun{}
	Nurdaulet.weapon.DoDamage(&Ruslan)
	fmt.Println("Ruslan health is:", Ruslan.health)
	//------
	Nurdaulet.weapon = Knife{}
	Nurdaulet.weapon.DoDamage(&Syndar)
	fmt.Println("Syndar health is:", Syndar.health)
	Nurdaulet.weapon.DoDamage(&Syndar)
	fmt.Println("Syndar health is: ", Syndar.health)

}
