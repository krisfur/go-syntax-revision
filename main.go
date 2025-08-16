package main

import "fmt" //format library for printing

func greet(name string) { // function that takes a string and prints a greeting
	fmt.Println("Hello,", name) // calling the function
}

func add2(a, b int) int { // function that takes two integers and returns their sum
	return a + b // returns the sum of a and b
}

// Regular function returning an anonymous function
func funception() func(i, j string) string {
	myf := func(i, j string) string {
		return i + j + "!"
	}
	return myf
}

// Variadic function that takes a variable number of integers and prints them along with their sum
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	fmt.Println("Hello, world! 👋")

	// ###########
	// Variables

	var name string = "Alice" // declaring a variable with type string, swap var for cosnt to define an immutable
	surname := "Smith"        // shorthand declaration, type inferred

	fmt.Println("Name:", name+",", "Surname:", surname) // comma separation adds a space, to append a comma after a name we just + it since they're strings
	age := 30
	fmt.Printf("Name: %s, Surname: %s, Age: %d\n", name, surname, age) // formatted output using Printf

	// #############
	// Flow control

	//if else
	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age < 65 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior citizen.")
	}

	// switch case
	switch age {
	case 0, 1, 2, 3, 4, 5: // multiple cases can be grouped
		fmt.Println("You are a toddler.")
	case 6, 7, 8, 9, 10, 11, 12: // another group
		fmt.Println("You are a child.")
	default: // default case if no other case matches
		fmt.Println("You are a teenager or older.")
	}

	// switch with type assertion
	// This is a way to check the type of a variable at runtime
	whatAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	// #######
	// Loops

	//for loop
	for i := range 5 { //same as for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// while loop (using for)
	i := 0
	for i < 5 { // same as while i < 5 {
		fmt.Println("While iteration:", i)
		i++ // incrementing i
	}

	// infinite loop (use with caution)
	for {
		fmt.Println("This will run forever unless stopped!")
		// break or return can be used to exit the loop
		break
	}

	// ##################
	// Arrays etc.

	arr := [5]int{1, 2, 3, 4, 5} // fixed-size array
	fmt.Println("Array:", arr)

	slice := []int{1, 2, 3, 4, 5} // dynamic-size slice
	fmt.Println("Slice:", slice)

	slice = append(slice, 6) // appending to a slice
	fmt.Println("Slice after append:", slice)
	fmt.Println("First element of slice:", slice[0]) // accessing elements by index
	fmt.Println("Length of slice:", len(slice))      // getting the length of a slice
	fmt.Println("Capacity of slice:", cap(slice))    // getting the capacity of a slice

	// for each loop over a slice
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Maps (like dictionaries in Python)
	myMap := map[string]int{"Alice": 30, "Bob": 25} // creating a map
	fmt.Println("Map:", myMap)
	myMap["Charlie"] = 35 // adding a new key-value pair
	fmt.Println("Map after adding Charlie:", myMap)
	fmt.Println("Value for Alice:", myMap["Alice"]) // accessing a value by key

	// #########
	// Functions

	// named functions are defined outside of main as functions that capture variables from their surrounding context
	greet(name) // calling the greet function with a string argument

	// anonymous functions with no name - these can be used for callbacks or immediate execution, the function is assigned to a variable
	add := func(a, b int) int { // name the function, takes two integers, retruns an integer
		return a + b
	}
	fmt.Println("Sum of 5 and 3 is:", add(5, 3))
	fmt.Println("Sum of 5 and 3 is:", add2(5, 3))

	fmt.Println("Funception result:", funception()("Hello", "World")) // calling the funception function, which returns an anonymous function and then calling that function with arguments

	// Variadic function call
	sum(1, 2, 3, 4, 5) // can have any number of arguments
	sum(1, 2, 3)       // calling the sum function with a variable number of arguments
}
