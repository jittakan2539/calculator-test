package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type foodSet struct {
	Name	string;
	Price	int;
}

func getInteger(prompt string) int {
	var reader = bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := strconv.Atoi(input)
		if err == nil {
			return value
		}

		fmt.Println("Invalid input. Please enter an integer.")
	}
}

func getYesNoInput(prompt string) string {
    var input string
    for {
        fmt.Print(prompt)
        fmt.Scanln(&input)
        input = strings.ToLower(input)

        if input == "y" || input == "n" {
            return input
        } else {
            fmt.Println("Invalid input. Please enter 'y' for Yes or 'n' for No.")
        }
    }
}

func main() {
	fmt.Println("Welcome! This is the Food Store Calculator at your service.")
	fmt.Println("Here are the list of our scrumptious menu.")
	fmt.Println("----------------------------------------")

	menu := []foodSet{
		{"Red set", 50},
		{"Green set", 40},
		{"Blue set", 30},
		{"Yellow set", 50},
		{"Pink set", 80},
		{"Purple set", 90},
		{"Orange set", 120},
	}

	discountedItems := map[string]bool{
		"Green set":  true,
		"Pink set":   true,
		"Orange set": true,
	}

	for i, food := range menu {
		fmt.Printf("%d. %-15s  %-3d THB/set\n", i+1, food.Name, food.Price)
	}

	fmt.Println("----------------------------------------")

	var normalQuantity []float64
	var discountQuantity []float64

	fmt.Println("Please type how many sets you would like to order. (Type 0 if you don't want that set. )")

	for _, food := range menu {
		quantity := getInteger(fmt.Sprintf("%s: ", food.Name))

		if quantity > 0 {
			if discountedItems[food.Name] && quantity >= 2 {
				discountQuantity = append(discountQuantity, float64(quantity * food.Price))
			} else {
				normalQuantity = append(normalQuantity, float64(quantity*food.Price))
			}
		}
	}

	var normalTotal float64
	for _, price := range normalQuantity {
		normalTotal += price
	}

	var discountTotal float64
	for _, price := range discountQuantity {
		discountTotal += price
	}

	afterDiscountTotal := float64(discountTotal)*0.95
	discount := float64(discountTotal)*0.05
	subtotal := afterDiscountTotal + normalTotal

	fmt.Println("\nOrder Summary:")
	fmt.Println("----------------------------------------")
	fmt.Printf("Total price for normal sets:                 	%10.2f THB\n", normalTotal)
	fmt.Printf("Total price for sets with (5 percent discount): %10.2f THB\n", discountTotal)
	fmt.Printf("Discount applied (5 percent):                	%10.2f THB\n", discount)
	fmt.Printf("Total price after discount:                 	%10.2f THB\n", afterDiscountTotal)
	fmt.Printf("Subtotal:                                    	%10.2f THB\n", subtotal)

	fmt.Println("----------------------------------------")
	
	var total float64
	var memberDiscount float64
	haveMember := getYesNoInput("Do you have a member card? (y/n): ")

	if haveMember == "y" {
		total = subtotal*0.9
		memberDiscount = subtotal*0.1
		fmt.Printf("You received 10 percent discount: %.2f THB", memberDiscount)
	}
	fmt.Println("----------------------------------------")
	fmt.Printf("Your Total is: %.2f THB\n", total)
	fmt.Println("----------------------------------------")
	fmt.Println("Thank you for using our service.")

	var exitInput string
	fmt.Println("\nType anything to quit.")
	fmt.Scanln(&exitInput)
}