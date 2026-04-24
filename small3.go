/*
Names: Finn Norsen, George Gordon, and Will Bentley
Assignment: Term Project Checkpoint 2
File Name: small3.go
Program: Tracks grades of inputted students and creates an average using closures and higher order functions
*/

package main
import ("fmt")


func studentGradeAccumulator() func(int) string {
	total, ct := 0, 0
	// This code block "remembers" the running total through the program
	// (Closure)
	return func(curr_score int) string {
		total += curr_score
		ct ++
		average := total / ct

		if average <= 64 {
			return "This might not be a good professor!"
		} 
		if average >= 65 && average <= 100 {
			return "Most people are passing."
		} else {
			return "This average is impossible!!!"
		}
	}
}

func main() {

	// Acts as a proxy to showcase HOFs
	proxy := studentGradeAccumulator()

	// Input students
	fmt.Printf("Adam was added. %s\n", proxy(68))
	fmt.Printf("Ella was added. %s\n", proxy(11))
	fmt.Printf("Zach was added. %s\n", proxy(91))
	fmt.Printf("John was added. %s\n", proxy(97))
	fmt.Printf("Ezekiel was added. %s\n", proxy(99))
	fmt.Printf("Mackenzie was added. %s\n", proxy(88))
	fmt.Printf("Zoe was added. %s\n", proxy(56))
	fmt.Printf("Carmine was added. %s\n", proxy(999))
}