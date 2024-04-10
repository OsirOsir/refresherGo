package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//helper function

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createInvoice() invoice {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create invoice name: ")
	// name, _ := reader.ReadString('\n')

	// name = strings.TrimSpace(name)
	name, _ := getInput("Create Invoice name: ", reader)

	i := newInvoice(name)

	fmt.Println("New Invoice created", i.name)

	return i
}

func promptOptions(i invoice) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose from a = Add Item, s = Save, t = Add Tip :  ", reader)
	// fmt.Println(opt)

	switch opt {
	case "a":
		item, _ := getInput("Item Name: ", reader)
		price, _ := getInput("Price Value", reader)
		fmt.Println(item, price)
	case "s":
		fmt.Println("You Selected s")
	case "t":
		tip, _ := getInput("Tip: $", reader)
		fmt.Println("$", tip)
	default:
		fmt.Print("That was invalid\n")
		promptOptions(i)
	}

}

func main() {
	myInvoice := createInvoice()
	promptOptions(myInvoice)
	// fmt.Println(myInvoice)
}
