package main

/*
   Key Concepts Covered:

   1. Arrays: Fixed-size, ordered collections of elements of a single type. The script shows how to declare, initialize, and manipulate arrays, including multi-dimensional arrays.

   2. Slices: Dynamic and flexible segments of an array. The script explores slice creation, appending elements, and creating sub-slices, illustrating their dynamic nature compared to arrays.

   3. Maps: Key-value pairs that provide a versatile way to store and manage data. The demonstration includes map creation, adding and removing elements, and iterating over maps.

   4. Structs: Custom data types that allow the grouping of multiple fields under a single unit. The script illustrates defining structs, creating instances, and modifying their fields.
*/

import "fmt"

// Arrays in Go
func demonstrateArrays() {
	// In GoLang arrays are a fixed size.
	// A type can be declared for each of the elements in the array.
		// var numbers [10]any can be used to hold any type, such as when an array has inconsistently typed elements.
			// any is an alias for interface{}
	var numbers [10]int
	for i := 0; i < len(numbers); i++ {
		//Initialize each element in the array.
		numbers[i] = i * 2
	}
	// Print the entire array to show the intialized values.
	fmt.Println("Numbers array:", numbers)

	// Demonstrate a multi-dimensional array. 3x3 matrix.
		// Arrays within arrays
	var matrix [3][3]int
	// To intialize the values here we need a loop within a loop, since this array is multi level.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Assign a value based on the sum of the row and column indices
			matrix[i][j] = i+j
		}
	}
	// Print the multi-dimensional array to show its structure
	fmt.Println("Matrix:", matrix)
}

// Slices in Go
func demonstrateSlices() {
	// Slices are similar to arrays, but are resizeable and more flexible.
	colors := []string{"Red", "Black", "Purple"}
	// Print our intial slices
	fmt.Println("Intial Colors Slice:", colors)

	// Demonstrating the flexibility of slices
	moreColors := []string{"Yellow", "Blue", "Orange"}
	// This will assign a new value to the colors slice, which will be the items of colors and the items of moreColors
		// `moreColors...`, uses the spread-operator to empty the contents of one thing into another.
	colors = append(colors, moreColors...)
	fmt.Println("Extended colors slice:", colors)

	// Create sub-slice from an existing slice. Syntax is `slice[low:high]`
	// This selects a range starting at 'low' and the ending is just before the `high` index
	susbsetColors := colors[1:4]
	// Print the sub-slice to demonstrate
	fmt.Println("Sub-set of colors:", susbsetColors)
}

// Maps in Go
func demonstrateMap() {
	// Maps are key-value pairs. They are similar to dictionaries in Python and Objects in JS.
	userAges := map[string]int{
		"Bob Marley": 33,
		"Jimmi Hendrix": 27,
	}
	// Adding new key-value pair to the map.
		// Maps are dynamic, pairs can be added at any time.
	userAges["Tupac"] = 23
	// Print the map to show the newly added pair
	fmt.Println("User ages with Tupac added:", userAges)

	// Delete key-value pairs from the map using the `delete` built-in function
	delete(userAges, "Bob Marley")
	// Print the map showing the state after deletion
	fmt.Println("User ages after deleting Bob Marley", userAges)

	// Make an empty map
		// make(map[<Key_TYPE>]<VALUE_TYPE>)
	// fruitMap := make(map[string]int)
}

// Structs in Go
func demonstrateStructs() {
	// A struct is a compsosite type that groups together valirables under a named unit.
	type Person struct {
		Name string
		Age int
	}

	// Create a slice of 'Person' structs with initial values.
	people := []Person{
		{Name: "Alice", Age:18},
		{Name: "Mad Hatter", Age: 34},
		{Name: "Cheshire Cat", Age: 40},
	}

	people[1].Age = 26

	// Print the slice of structs to show all persons
	fmt.Println("People slice:", people)
}

func main() {
    // Call various functions that will be used to explore different concepts. Each function will print to the console.

	demonstrateArrays() // Shows fixed-sized arrays and multi-dimensional arrays.
	demonstrateSlices() // Shows slices, which do not have fixed sizes. Append and create sub-slices
	demonstrateMap() // Maps are key-value pairs; will show adding, and deleting elements.
	demonstrateStructs() // Explains struct definition and manipulation in a slice of structs.
}
