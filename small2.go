package main
import ("fmt"
		"slices"
)


func reverse_slice(my_slice []string) {

for i, j :=0, len(my_slice) -1; i < j; i, j = i + 1, j - 1 {
	my_slice[i], my_slice[j] = my_slice[j], my_slice[i]
}

}

func main() {
	// Declare the array to be sliced from and the first slice
	my_array := [7]string{"I", "am", "a", "slice", "in", "an", "array"}
	start_slice := my_array[0:4]

	fmt.Println("The entire array contains: ", my_array)

	// Slice size in go is determined by the starting index and the size of whatever it is slicing
	// So, the slice starting at index of 0 has a size of 7, but a size of 4.
	fmt.Println("The slice of the array starting at zero contains:", start_slice, "and has a capacity of", cap(start_slice), "with an occupied length of", len(start_slice))

	// An example of size determination slicing in the middle:
	middle_slice := my_array[2:4]
	fmt.Println("The slice of the array starting at zero contains:", middle_slice, "and has a capacity of", cap(middle_slice), "with an occupied length of", len(middle_slice))

	slices.Reverse(middle_slice)
	fmt.Println("Slice reversed via inbuilt function:", middle_slice)

	// Slices also have many inbuilts, which make some code non necessary, like this reversing function I just wrote:
	reverse_slice(middle_slice)
	fmt.Println("Re-reversed slice using looped statement:", middle_slice)
	
	// Since a slice is basically a pointer list to the main array, it actually edits the main array! Lets see what happens before and after reversing...
	fmt.Println("Before reversal:", my_array, "\n")
	reverse_slice(start_slice)
	fmt.Println("After reversal:", my_array, "\n")

}