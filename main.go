package main

import "fmt"

var inventory []Product

func addProduct(newProduct Product) {
	inventory = append(inventory, newProduct)
}

func displayInventory() (result string) {
	for _, product := range inventory {
		result += fmt.Sprintf("Name: %s; Price %f; Stock %d\n", product.Name, product.Price, product.Stock)
	}
	return
}

func checkStock() bool {
	for i := 0; i < len(inventory); i++ {
		if (inventory[i].Stock < 5) {
			fmt.Printf("Warning: %s is running low. Only %d left in stock.\n", inventory[i].Name, inventory[i].Stock)
			return true
		}
	}
	return false
}

func main() {
	addProduct(Product{ Name: "Laptop", Price: 1000.0, Stock: 10 })
	addProduct(Product{ Name: "Mouse", Price: 25.5, Stock: 20 })
	addProduct(Product{ Name: "Keyboard", Price: 50.0, Stock: 2 })

	fmt.Println(displayInventory())
	checkStock()
}
