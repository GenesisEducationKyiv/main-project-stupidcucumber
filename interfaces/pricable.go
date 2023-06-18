package interfaces

type Pricable interface {
	GetPrice() (float64, error)
}
