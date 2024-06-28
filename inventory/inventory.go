package inventory

import "github.com/Xeckt/learn-go/item"

type Inventory struct {
	Items    map[string]item.ItemManager
	Capacity int
}

func NewInventory() *Inventory {
	return &Inventory{
		Items:    make(map[string]item.ItemManager),
		Capacity: 10,
	}
}

func (i *Inventory) AddItem(item item.ItemManager) {
	i.Items[item.Name] = item
}

func (i *Inventory) RemoveItem(item item.ItemManager) {
	delete(i.Items, item.Name)
}

func (i *Inventory) GetItem(item item.ItemManager) item.ItemManager {
	return i.Items[item.Name]
}
