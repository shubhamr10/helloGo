//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	Health, MaxHealth int
	Energy, MaxEnergy int
	Name              string
}

func (player *Player) modifyHealth(newHealth int) {
	if player.Health > player.MaxHealth {
		player.Health = newHealth
	}
}

func (player *Player) modifyEnergy(newEnergy int) {
	if player.Energy > player.MaxEnergy {
		player.Energy = newEnergy
	}
}

func main() {
	player1 := Player{90, 150, 80, 100, "Mario"}
	player2 := Player{80, 150, 60, 100, "Luiggi"}

	fmt.Println("Initial health of Player", player1.Name)
	fmt.Println(player1)
	fmt.Println("Initial health of Player", player2.Name)
	fmt.Println(player2)

	player1.modifyHealth(92)
	player1.modifyEnergy(40)

	player2.modifyEnergy(40)
	player2.modifyHealth(92)

	fmt.Println("Modified health of Player", player1.Name)
	fmt.Println(player1)
	fmt.Println("Modified health of Player", player2.Name)
	fmt.Println(player2)

}
