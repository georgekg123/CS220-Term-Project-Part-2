/*
Names: Finn Norsen, George Gordon, and Will Bentley
Assignment: Term Project Checkpoint 2
File Name: small1.go
Program: Utilizes a struct to compare sports player data
*/

package main

import (
	"fmt"
)

type PlayerInfo struct {
	name, club string
	goals int
	app bool
}

func PlayerSetup() {
	var name1 string
	fmt.Print("Name your first player: ")
	fmt.Scan(&name1)
	var club1 string
	fmt.Print("Name that player's club: ")
	fmt.Scan(&name1)
	var goals1 int
	fmt.Print("How many goals have they scored? ")
	fmt.Scan(&name1)
	var app1 bool
	fmt.Print("Have they made an appearance (true/false)? ")
	fmt.Scan(&name1)

	var name2 string
	fmt.Print("Name your second player: ")
	fmt.Scan(&name1)
	var club2 string
	fmt.Print("Name that player's club: ")
	fmt.Scan(&name1)
	var goals2 int
	fmt.Print("How many goals have they scored? ")
	fmt.Scan(&name1)
	var app2 bool
	fmt.Print("Have they made an appearance (true/false)? ")
	fmt.Scan(&name1)

	var player1 = PlayerInfo{name1, club1, goals1, app1}
	var player2 = PlayerInfo{name2, club2, goals2, app2}
}


func main() {

}

