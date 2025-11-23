package main

import (
	"fmt"
	"time"
)

func doWork(done chan bool) {
	fmt.Println("Worker: Starting long job...")
	time.Sleep(1 * time.Second)
	fmt.Println("Worker: Long job done.")
	done <- true
}

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

func checkStock() error {
	for i := 0; i < len(inventory); i++ {
		if (inventory[i].Stock < 5) {
			return fmt.Errorf("stock of %s is running low. only %d left in stock", inventory[i].Name, inventory[i].Stock)
		}
	}
	return nil
}

func main() {
	done := make(chan bool)
	go doWork(done)

	addProduct(Product{ Name: "Laptop", Price: 1000.0, Stock: 10 })
	addProduct(Product{ Name: "Mouse", Price: 25.5, Stock: 20 })
	addProduct(Product{ Name: "Keyboard", Price: 50.0, Stock: 2 })

	fmt.Println(displayInventory())

	error := checkStock()
	
	if error != nil {
		fmt.Println(error)
	}


	<-done
	fmt.Println("Main: Worker confirmed done, program exiting.")


	fmt.Print("Enter your name: ")

	var name string

	fmt.Scanln(&name)

	fmt.Println("Hello", name)
}
