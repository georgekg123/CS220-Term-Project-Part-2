/*
Authors: Finn Norsen, George Gordon, and Will Bentley
Assignment: Term Project Checkpoint 2
File Name: large.go
Program: Utilizes a struct and interfaces for a finals week simulation game

Source: https://pkg.go.dev/os
Desc: explains the os library to write to a file

Source: https://pkg.go.dev/fmt
Desc: explains the fmt library to format strings and get input or give output

Souce: https://www.tutorialspoint.com/go/index.htm
Desc: Go book used to learn language basic

Source: https://go.dev/tour/welcome/1
Desc: walkthrough of lots of different go functionality

How to Run:
1. Install Go from golang.org.
2. Open a terminal and navigate to the folder containing this file
3. Run the command: go run large.go
*/

package main

// import fmt package for IO
// import os package to write to a file
import (
	"fmt"
	"os"
)

// In Go, structs replace classes. They are just a collection data.
// Variables that are capitalized are public.
type Student struct {
	Name   string
	Stress int
	Sleep  int
	Grades int
	Day    int
}

// Interface gives requirements a struct must have to join the interface
// In this case:
// 1. A method called Title() that returns a string.
// 2. A method called Apply() that takes a pointer to a Student.
//
// If a struct has those two methods, Go automatically
// lets it join the interface. We can then treat structs within the interface as if they were
// the same type

type Action interface {
	Apply(s *Student)
	Title() string
}

// Study is just an empty struct. We use it as a placeholder to attach methods.
type Study struct{}

// By defining these two methods with for Study, 
// Study is now part of the Action interface automatically.
func (a Study) Title() string {
	return "Study at the Library"
}

func (a Study) Apply(s *Student) {
	fmt.Println(">> You hit the library hard. You are tired and stressed but your Grades are looking up.")
	s.Grades += 15
	if s.Grades > 100 {
		s.Grades = 100
	}
	s.Stress += 25
	s.Sleep -= 30
}

type SleepAction struct{}

func (a SleepAction) Title() string {
	return "Sleep all day"
}

func (a SleepAction) Apply(s *Student) {
	fmt.Println(">> You sleep deeply. You dream of the beach.")
	s.Sleep += 35
	if s.Sleep > 100 {
		s.Sleep = 100
		fmt.Println(">> You've slept enough")
	}
	s.Stress -= 15
	if s.Stress < 0 {
		s.Stress = 0
	}
}

type Socialize struct{}

func (a Socialize) Title() string {
	return "Hang out with Friends"
}

func (a Socialize) Apply(s *Student) {
	fmt.Println(">> You played cards with friends and then went out for lunch. Your stress is down, but so is your focus.")
	s.Stress -= 25
	if s.Stress < 0 {
		s.Stress = 0
	}
	s.Grades -= 5
}

//this is the overall game function, and the main which is automatically called when the program is run
func main() {
	fmt.Print("Enter Student Name: ")
	var name string
	
	// Error handling for input
	// _ ignores the first output, and err accepts the error
	// fmt.Scanln automatically read input into the given address
	_, err := fmt.Scanln(&name)
	if err != nil {
		name = "Anonymous Student"
	}

	// Initializing a new student object
	student := &Student{
		Name:   name, 
		Sleep:  80, 
		Grades: 35,
		Stress: 25,
		Day:    1,
	}

	// go print function
	// %s indicates a variable that is a string, which is then passed
	fmt.Printf("\nWelcome, %s. It's Finals Week. Survive 7 days.\n", student.Name)

	//Note: all loops are for loops in go
	for student.Day <= 7 {
		performDay(student)
		student.Day++
	}

	saveFinalReport(student)
}

//This is the function that gets terminal input and handles the action
func performDay(s *Student) {
	// 'defer' and 'recover' handle a panic
	// if the code hits a panic, the defer gets called, with its error passed to 
	// recover. 
	// defer causes a given function to be run
	// func() { ... }() is the notation for an anonymous function, which is the function
	//defer runs here.

	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("\n!!! CRITICAL COLLAPSE: %v !!!\n", r)
			fmt.Println("You woke up in the campus clinic. You missed progress.")
			s.Stress, s.Sleep = 50, 50
			s.Grades -= 10
		}
	}()

	fmt.Printf("\n--- DAY %d | Stress: %d | Sleep: %d | Grades: %d ---\n", s.Day, s.Stress, s.Sleep, s.Grades)

	// Here we use the interface, we define a slice, which is the same as a list 
	//in other languages for this use case. Action is the type of the data in the slice.
	actions := []Action{Study{}, SleepAction{}, Socialize{}}
	
	for i, act := range actions {
		// we can call .Title() on a struct in the interface, and Go knows which struct's method to run.
		fmt.Printf("(%d) %s\n", i+1, act.Title())
	}

	fmt.Print("Choose Action: ")
	var choice int
	_, scanErr := fmt.Scanln(&choice)

	if scanErr == nil && choice > 0 && choice <= len(actions) {
		//The interface can now handle the logic regardless of which choice was made
		actions[choice-1].Apply(s)
	} else {
		fmt.Println(">> You stared at a wall and did nothing.")
		s.Stress += 10
	}

	//when these stats reach a threshold, trigger a manual panic
	//this is not when a panic would usually be used, as it can handle all kind
	// of memory errors than can be fixed with defer
	if s.Stress >= 100 {
		panic("STRESS OVERLOAD")
	}
	if s.Sleep <= 0 {
		panic("SLEEP DEPRIVATION FAINT")
	}
}

//this is a function to write the final result to a file
func saveFinalReport(s *Student) {

	report := fmt.Sprintf("Final Grade for %s: %d\n", s.Name, s.Grades)

	//saving to a text file using the os library
	//must write bytes 
	// 0644 represents the permissions of the file
	err := os.WriteFile("final_report.txt", []byte(report), 0644)
	
	if err != nil {
		fmt.Printf("Error saving report: %v\n", err)
	} else {
		fmt.Println("\nWeek completed. Results saved to final_report.txt")
	}
}