package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1
	fmt.Println("Shopping list", shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Shopping list:: delete milk", shoppingList)

	fmt.Println("Need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	fmt.Println("Need cereal ?")
	if !found {
		fmt.Println("node")
	} else {
		fmt.Println("Yup", cereal, "boxes.")
	}

	totalItems := 0
	for _, items := range shoppingList {
		totalItems += items
	}

	fmt.Println("there are ", totalItems, "on the shopping list.")
}
