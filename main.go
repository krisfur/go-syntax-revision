package main

import (
	"bufio"
	"fmt" //format library for printing
	"os"  //os library for operating system functionality
)

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
	default: // default case if no other case matches - optional
		fmt.Println("You are a teenager or older.")
	}

	// there's no need in Go for a branchless programming approach, the compiler already avoids jumps and implements cmove for you

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

	//recursion
	// you can use recursion in defined functions outside of main, or in anonymous functions

	// to do it in anonymous functions, you need to declare the function variable first, then assign it a function that refers to itself
	var factorial func(n int) int // declaring a variable of type function that takes an int and returns an int

	factorial = func(n int) int {
		if n == 0 {
			return 1 // base case
		}
		return n * factorial(n-1) // recursive case
	}

	fmt.Println("Factorial of 5 is:", factorial(5)) // calling the recursive function

	// #################
	// Defer

	// Defer is used to ensure that a function call is performed later in the program's or function's execution, usually for cleanup purposes
	printFile := func(filename string) { // anonymous function to read and print a file
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer f.Close() // defer ensures the file is closed after the function completes
		fmt.Println("File opened successfully:", filename)
		fmt.Println("--- File Content ---")

		// create a new scanner for the file
		scanner := bufio.NewScanner(f)

		// loop through each line of the file and print it
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		// check for any errors that occurred during scanning
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}
		fmt.Println("--- End of File ---")
	}

	printFile("example.txt") // calling the function to read and print a file

	// ###############
	// Error handling

	// Go uses multiple return values to handle errors, typically returning a value and an error
	divide := func(a, b int) (int, error) {
		if b == 0 {
			return 0, fmt.Errorf("division by zero") // returning an error if b is zero
		}
		return a / b, nil // returning the result and nil for no error
	}
	result, err := divide(10, 0) // calling the divide function
	if err != nil {
		fmt.Println("Error:", err) // printing the error if it exists
	} else {
		fmt.Println("Result of division:", result) // printing the result if no error
	}

	// ##########
	// Pointers

	x := 10                                                            // declaring an integer variable
	y := &x                                                            // y is a pointer to x, it holds the memory address of x
	fmt.Println("Value of x:", x)                                      // prints the value of x
	fmt.Println("Address of x:", &x)                                   // prints the memory address of x
	fmt.Println("Value of y (pointer to x):", *y)                      // dereferencing y to get the value of x
	*y = 20                                                            // changing the value of x through the pointer y
	fmt.Println("New value of x after changing through pointer y:", x) // prints the new value of x

	// ####################
	// Structs and methods

	// #######
	// Enums

	// #############
	// Interfaces

	// ###########################
	// Concurrency and Goroutines

	// #########
	// Channels
}
