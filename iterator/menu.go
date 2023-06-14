package iterator

type DinerMenuIterator struct {
	Iterator
	Items    []MenuItem
	Position int
}

var _ Iterator = (*DinerMenuIterator)(nil)

func (dmi *DinerMenuIterator) Constructor(items *[]MenuItem) {
	dmi.Items = *items
}

func (dmi *DinerMenuIterator) HasNext() bool {
	if dmi.Position >= len(dmi.Items) || dmi.Items[dmi.Position] == (MenuItem{}) {
		return false
	}
	return true
}

func (dmi *DinerMenuIterator) Next() MenuItem {
	item := dmi.Items[dmi.Position]
	dmi.Position++
	return item
}

type DinerMenu struct {
	MaxItems      int // Go는 상수를 구조체 필드로 넣을 수 없다
	NumberOfItems int
	MenuItems     []MenuItem
}

func (dm *DinerMenu) Constructor() {
	dm.MaxItems = 6 //
	dm.MenuItems = []MenuItem{}
	dm.AddItem()
	dm.AddItem()
	dm.AddItem()
}

func (dm *DinerMenu) AddItem() {
	// ...
}

func (dm *DinerMenu) CreateIterator() Iterator {
	dmi := new(DinerMenuIterator)
	dmi.Constructor(&dm.MenuItems)
	return dmi
}
