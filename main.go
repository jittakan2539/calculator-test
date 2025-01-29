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

var reader = bufio.NewReader(os.Stdin)

func getInteger(prompt string) int {
	fmt.Printf("%v", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must be an integer only.", prompt)
		panic(message) 
	}

	return int(value)
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

	fmt.Println("----------------------------------------")

	quantity := make([]int, 7)
	fmt.Println("Please type how many sets you wpuld like to order. (Type 0 if you don't want that set. )")

	for i, food := range menu {
		quantity[i] = getInteger(fmt.Sprintf("%s: ", food.Name))
	}




	// fmt.Println("Please type the number you would like order for each set. If you don't, please type 0.")
}