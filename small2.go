package main
import ("fmt")


func reverse_slice(my_slice string) {

}






func main() {
	my_array := [7]string{"I", "am", "a", "slice", "in", "an", "array"}
	start_slice := my_array[0:4]


	fmt.Println("The entire array contains: ", my_array)

	// Slice size in go is determined by the starting index and the size of whatever it is slicing
	// So, the slice starting at index of 0 has a size of 7, but a size of 4.
	fmt.Println("The slice of the array starting at zero contains:", start_slice, "and has a capacity of", cap(start_slice), "with an occupied length of", len(start_slice))

	// An example of size determination slicing in the middle:
	middle_slice := my_array[2:4]
	fmt.Println("The slice of the array starting at zero contains:", middle_slice, "and has a capacity of", cap(middle_slice), "with an occupied length of", len(middle_slice))

}