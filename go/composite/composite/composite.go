package composite

type Item interface {
	Add(Item)
	String() string
	GetPrice() int
}
