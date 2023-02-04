package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// Create the main function
func main() {

	//Create an infinite loop that will run until the user enters the word "exit".
	// for {
	// 	//Ask the user to enter a symbol.
	// 	fmt.Println("Enter a symbol or type 'help': ")
	// 	//Create a new scanner that reads from the standard input.
	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	//Scan the input and save it as a variable called input.
	// 	scanner.Scan()
	// 	input := scanner.Text()
	// 	//If the user enters the word "exit" then break out of the loop.
	// 	if input == "exit" {
	// 		break
	// 	}

	// 	//If the user enters the word "help" then print the help message.
	// 	if input == "help" {
	// 		fmt.Println("Enter a symbol to get the data for that symbol.")
	// 		fmt.Println("Enter 'exit' to exit the program.")
	// 		fmt.Println("Enter 'help' to see this message again.")
	// 		continue
	// 	}

	// 	//If the user enters a symbol then print the symbol.
	// 	fmt.Println("You entered: " + input)
	// }

	fmt.Println("Extracting Data...")
	fmt.Println()

	downloadData()

	fmt.Println("Data Extracted!")

	matrix := getSymbolMatrix()

	//print Symbol and Name as two separate columns using tabwriter
	//ask user for symbol to search for
	boolean := false
	var symbol string
	fmt.Println("\n" + "Enter a symbol to search for, or type help to see a full list: ")
	fmt.Scanln(&symbol)

	//if user enters help, print the full list of symbols
	if symbol == "help" {

		fmt.Print("\n" + "Here is the data: " + "\n" + "\n")
		//print matrix[1][i] and  matrix[0][i] as two separate columns using tabwriter
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
		fmt.Fprintln(w, "Symbol\tName")
		for i := 0; i < len(matrix[0]); i++ {
			fmt.Fprintln(w, matrix[0][i]+"\t"+"  "+matrix[1][i])
		}
		w.Flush()

		fmt.Println("\n" + "Enter a symbol to search for: ")
		fmt.Scanln(&symbol)

		//check if symbol is found in matrix[0][i]
		fmt.Println("\n" + "Checking if symbol is valid: ")
		fmt.Println()

		//check if symbol is found in matrix[0][i]
		for i := 0; i < len(matrix[0]); i++ {
			if symbol == matrix[0][i] {
				boolean = true
				fmt.Println("Symbol found!")
				fmt.Println()
				break
			}
		}
	}

	//check if symbol is found in matrix[0][i]
	fmt.Println("\n" + "Checking if symbol is valid: ")
	fmt.Println()

	//check if symbol is found in matrix[0][i]
	for i := 0; i < len(matrix[0]); i++ {
		if symbol == matrix[0][i] {
			boolean = true
			break
		}
	}

	if boolean == true {
		//if symbol is found, print the data for that symbol
		fmt.Println("\n" + "Symbol found!")
		fmt.Println()

		data := getDailyStatements(symbol)
		printDailyStatements(data)

		fmt.Println()
		fmt.Println("Overview Statements: ")
		fmt.Println()
		data1 := getOverview(symbol)
		printOverviewStatement(data1)
	} else {
		fmt.Println("Symbol not found!")
	}

}

// TitleCard struct
type TitleCard struct {
	Title  string
	Field1 string
	Field2 string
	Field3 string
	Field4 string
}
