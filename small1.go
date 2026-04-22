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

func PlayerComparison() {
	var name1 string
	fmt.Print("Name your first player (last name only): ")
	fmt.Scan(&name1)
	var club1 string
	fmt.Print("Name that player's club (one word only): ")
	fmt.Scan(&club1)
	var goals1 int
	fmt.Print("How many goals have they scored? ")
	fmt.Scan(&goals1)
	var app1 bool
	fmt.Print("Have they made an appearance (true/false)? ")
	fmt.Scan(&app1)

	var name2 string
	fmt.Print("Name your second player (last name only): ")
	fmt.Scan(&name2)
	var club2 string
	fmt.Print("Name that player's club (one word only): ")
	fmt.Scan(&club2)
	var goals2 int
	fmt.Print("How many goals have they scored? ")
	fmt.Scan(&goals2)
	var app2 bool
	fmt.Print("Have they made an appearance (true/false)? ")
	fmt.Scan(&app2)

	var player1 = PlayerInfo{name1, club1, goals1, app1}
	var player2 = PlayerInfo{name2, club2, goals2, app2}

	var i string
	fmt.Print("What would you like to compare (goals, app) ")
	fmt.Scan(&i)

	if i == "goals" {
		if player1.goals < player2.goals {
			fmt.Println(player2.name + " has scored more goals than " + player1.name + "!")
		} else if player1.goals > player2.goals {
			fmt.Println(player1.name + " has scored more goals than " + player2.name + "!")
		} else {
			fmt.Println(player1.name + " and " + player2.name + " have scored the same number of goals!")
		}
	} else if i == "app" {
		if player1.app == true && player2.app == true {
			fmt.Println(player1.name + " and " + player2.name + " have both made appearances!")
		} else if player1.app == true && player2.app == false {
			fmt.Println(player1.name + " has made an appearance but " + player2.name + " has not.")
		} else if player2.app == true && player1.app == false {
			fmt.Println(player2.name + " has made an appearance but " + player1.name + " has not.")
		} else {
			fmt.Println("Neither " + player1.name + " nor " + player2.name + " have made an appearance.")
		}
	}
}


func main() {
	PlayerComparison()
}

