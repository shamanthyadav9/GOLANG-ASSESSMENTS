package main

import (
	"fmt"
)

// Slice of product names
var products = []string{"Laptop", "Mouse", "Keyboard", "Monitor"}

// Map to store stock quantity
var inventory = make(map[string]int)

func main() {
	// Initialize inventory with zero stock
	for _, product := range products {
		inventory[product] = 0
	}

	for {
		fmt.Println("\nInventory System")
		fmt.Println("1. Add Stock")
		fmt.Println("2. Reduce Stock")
		fmt.Println("3. Show All Products")
		fmt.Println("4. Show Low Stock Alerts (<5)")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addStock()
		case 2:
			reduceStock()
		case 3:
			showAllStock()
		case 4:
			showLowStock()
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addStock() {
	var name string
	var amount int
	fmt.Print("Enter product name: ")
	fmt.Scan(&name)
	if _, exists := inventory[name]; exists {
		fmt.Print("Enter amount to add: ")
		fmt.Scan(&amount)
		inventory[name] += amount
		fmt.Println("Stock updated.")
	} else {
		fmt.Println("Product does not exist.")
	}
}

func reduceStock() {
	var name string
	var amount int
	fmt.Print("Enter product name: ")
	fmt.Scan(&name)
	if _, exists := inventory[name]; exists {
		fmt.Print("Enter amount to reduce: ")
		fmt.Scan(&amount)
		if inventory[name] >= amount {
			inventory[name] -= amount
			fmt.Println("Stock reduced.")
		} else {
			fmt.Println("Not enough stock to reduce.")
		}
	} else {
		fmt.Println("Product does not exist.")
	}
}

func showAllStock() {
	fmt.Println("\nCurrent Inventory:")
	for _, product := range products {
		fmt.Printf("%s: %d\n", product, inventory[product])
	}
}

func showLowStock() {
	fmt.Println("\nLow Stock Alerts (less than 5):")
	alertFound := false
	for _, product := range products {
		if inventory[product] < 5 {
			fmt.Printf("%s: %d\n", product, inventory[product])
			alertFound = true
		}
	}
	if !alertFound {
		fmt.Println("All products have sufficient stock.")
	}
}
