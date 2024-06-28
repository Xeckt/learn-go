package main

import (
	"github.com/Xeckt/learn-go/player"
)

func main() {
	inventory := player.Inventory{
		Items:    make(map[string]player.ItemManager),
		Capacity: 10,
	}
	inventory.AddItem(player.NewItemManager().GenerateItem(player.HealthPotion, 10))
}
