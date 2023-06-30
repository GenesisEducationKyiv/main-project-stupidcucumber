package providers

type PriceProvider interface {
	GetPrice() (float64, error)
}
