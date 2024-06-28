package item

type ItemManager struct {
	Id       int
	Name     string
	Quantity int
}

type Item int

const (
	HealthPotion Item = iota
	ManaPotion
	Sword
	Shield
)

func GenerateItem(item Item, quantity int) ItemManager {
	switch item {
	case HealthPotion:
		return ItemManager{
			Id:       int(HealthPotion),
			Name:     "Mana Potion",
			Quantity: quantity,
		}
	case ManaPotion:
		return ItemManager{
			Id:       int(ManaPotion),
			Name:     "Mana Potion",
			Quantity: quantity,
		}
	case Sword:
		return ItemManager{
			Id:       int(Sword),
			Name:     "Sword",
			Quantity: quantity,
		}
	case Shield:
		return ItemManager{
			Id:       int(Shield),
			Name:     "Shield",
			Quantity: quantity,
		}
	}
	return ItemManager{}
}
