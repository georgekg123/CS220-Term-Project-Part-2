package main
import ("fmt")

func main() {
	my_array := [7]string{"I", "am", "a", "slice", "in", "an", "array"}
	my_slice := my_array[0:4]

	fmt.Println("The entire array contains: ", my_array)
	fmt.Println("The slice of the array contains: ", my_slice)

	my_slice = append(my_slice, my_array[0])
}