package player

type Inventory struct {
	Items    map[string]ItemManager
	Capacity int
}

func NewInventory() *Inventory {
	return &Inventory{
		Items:    nil,
		Capacity: 10,
	}
}

func (i *Inventory) AddItem(item ItemManager) {
	i.Items[item.Name] = item
}

func (i *Inventory) RemoveItem(item ItemManager) {
	delete(i.Items, item.Name)
}

func (i *Inventory) GetItem(item ItemManager) ItemManager {
	return i.Items[item.Name]
}
