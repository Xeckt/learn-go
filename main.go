package main

import (
	"github.com/Xeckt/learn-go/inventory"
	"github.com/Xeckt/learn-go/item"
)

func main() {
	playerInventory := inventory.Inventory{
		Items:    make(map[string]item.ItemManager),
		Capacity: 10,
	}
	playerInventory.AddItem(item.GenerateItem(item.HealthPotion, 10))
}
