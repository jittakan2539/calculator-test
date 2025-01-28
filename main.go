package main

import "fmt"

type foodSet struct {
	Name	string;
	Price	int;
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

	for i, food := range menu {
		fmt.Printf("%d. %-15s  %-3d THB/set\n", i+1, food.Name, food.Price)
	}


	// fmt.Println("Please type the number you would like order for each set. If you don't, please type 0.")
}