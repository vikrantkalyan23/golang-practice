// Run this lesson from the repository root:
// go run 1_hello-go/main.go

// package main tells Go that this package is a runnable program.
package main

// Import fmt so we can use its functions to display text.
import "fmt"

// main is where this program starts running.
func main() {
	// 1. COMMENTS
	// Comments are notes for people reading the code.
	// Go does not run comments as instructions.

	// This is a single-line comment.

	/*
		This is a block comment.
		It can cover more than one line.
		Block comments cannot be nested inside each other.
	*/

	fmt.Println("Hello Go!") // This is an inline comment beside code.
	// Output: Hello Go!

	// You can temporarily disable a line by adding // before it:
	// fmt.Println("This line will not run")

	// A documentation comment sits directly above the item it explains.
	// The comment above func main is an example.

	// 2. fmt.Println(): PRINT AND MOVE TO THE NEXT LINE
	// Text goes inside double quotes: "Hello".
	// A function performs an action. We call it using parentheses: (...).
	// Println adds spaces between values and a newline at the end.
	// A newline moves the next output to the next line.
	fmt.Println("Hello", "Go")
	fmt.Println("I am learning Go.")
	// Output:
	// Hello Go
	// I am learning Go.

	// 3. fmt.Print(): PRINT WITHOUT MOVING TO THE NEXT LINE
	// These two calls continue on the same line.
	fmt.Print("Hello ") // The space inside the quotes is intentional.
	fmt.Print("Go!")
	fmt.Print("\n") // \n means "start a new line".
	// Output: Hello Go!

	fmt.Print("Hello", "Go", "\n")
	// Output: HelloGo
	// Print does not add a space between these strings.
	// Detail for later: between two non-string values, it does add a space.
	fmt.Print(10, 20, "\n")
	// Output: 10 20

	// 4. fmt.Printf(): PUT VALUES INTO A MESSAGE
	// A variable is a named value. := creates a variable and gives it a value.
	name := "Vikrant"
	age := 25

	// %s is a placeholder for text; %d is a placeholder for an integer.
	// Values after the message fill the placeholders in the same order.
	// Printf does not add a newline automatically, so we include \n.
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)
	// Output: My name is Vikrant and I am 25 years old.

	// Start with these common placeholders:
	// %s   text (a string)
	// %d   a whole number (an integer)
	// %.2f a decimal number displayed with two digits after the decimal point
	// %t   true or false (a boolean)
	// %v   a value in its default format
	// %T   the type of a value
	// %%   a literal percent sign
	fmt.Printf("Price: %.2f\n", 12.5)
	// Output: Price: 12.50
	fmt.Printf("Learning Go: %t\n", true)
	// Output: Learning Go: true
	fmt.Printf("Value: %v, type: %T\n", age, age)
	// Output: Value: 25, type: int
	fmt.Printf("Progress: %d%%\n", 50)
	// Output: Progress: 50%

	// 5. fmt.Sprintf(): MAKE A STRING FOR LATER
	// Sprintf formats text but does not display it.
	// Save the result in a variable, then print it when needed.
	message := fmt.Sprintf("Welcome, %s!", name)
	fmt.Println(message)
	// Output: Welcome, Vikrant!

	// Sprint and Sprintln also make strings instead of displaying them.
	joined := fmt.Sprint("Hello, ", name)
	line := fmt.Sprintln("Hello,", name)
	fmt.Println(joined)
	fmt.Print(line) // Sprintln already included the newline.
	// Output:
	// Hello, Vikrant
	// Hello, Vikrant

	// 6. BUILT-IN print() AND println()
	// These work without importing fmt, but are meant for debugging.
	// Prefer fmt.Println and fmt.Printf for ordinary program output.
	// To try them, remove // from the following lines:
	// print("Hello", "Go", "\n") // HelloGo, then a newline from \n.
	// println("Hello", "Go")      // Hello Go, then an automatic newline.
	// Their exact formatting is not guaranteed by the language.
	// The current Go implementation sends them to the error output stream
	// (stderr). fmt.Print* uses the normal output stream (stdout).
	// Both can appear in the terminal, but their combined order can vary.

	// 7. OTHER PRINT FUNCTIONS: RECOGNIZE THESE FOR LATER
	// You do not need these to write your first programs.
	// Fprint, Fprintln, and Fprintf write to a destination you choose,
	// such as a file. Their spacing rules match Print, Println, and Printf.
	// Example syntax (requires importing os as well as fmt):
	// fmt.Fprintln(os.Stdout, "Hello Go!")
	// os.Stdout means the program's normal output destination.
	// These functions return a byte count and an error. Learn to check
	// that error when working with files and other output destinations.
	// Print, Println, and Printf also return a byte count and an error;
	// we omit those return values in this introductory lesson.

	// QUICK GUIDE
	// Display a simple message:          fmt.Println("Hello")
	// Continue printing on one line:     fmt.Print("Hello ")
	// Insert values into a message:      fmt.Printf("Age: %d\n", age)
	// Save a formatted message:          fmt.Sprintf("Age: %d", age)

	// PRACTICE
	// 1. Change name and age above, then run the program again.
	// 2. Use Println to display your favorite language.
	// 3. Use two Print calls to display "I love Go" on one line.
	// 4. Use Printf to display a price with two decimal places.
	// 5. Comment out one print statement and see how the output changes.
}

// Official reference: https://pkg.go.dev/fmt
