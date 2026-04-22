/*
Names: Finn Norsen, George Gordon, and Will Bentley
Assignment: Term Project Checkpoint 2
File Name: small1.go
Program: Utilizes a struct to compare sports player data!
*/

package main

// Imports fmt, which is used for printing:
import (
	"fmt"
)

// Sets up our struct, which is a data type in Go that can groups different types into a single structure.
type PlayerInfo struct {
	// Types are declared along with names; any type can be accepted (int, string, bool, etc.).
	name, club string
	goals int
	app bool
}

// PlayerComparison prompts the user to enter info about two players, then lets them compare them, having entered the data into the struct.
func PlayerComparison() {

	// Takes user input using Print and Scan. No spaces are allowed for individual input!
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

	// Enters the gathered info into entries of the PlayerInfo struct:
	var player1 = PlayerInfo{name1, club1, goals1, app1}
	var player2 = PlayerInfo{name2, club2, goals2, app2}

	// Takes user input to decide what to compare:
	var i string
	fmt.Print("What would you like to compare (goals, app) ")
	fmt.Scan(&i)

	// Utilizes simple conditionals concatenation to output the correct statistical comparisons:
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

// The main function controls the flow of the program; in this case, it calls the PlayerComparison function:
func main() {
	PlayerComparison()
}

