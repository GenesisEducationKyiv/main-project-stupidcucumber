package mocks

type PriceProviderMock struct {
	Error error
	Price float64
}

func (p *PriceProviderMock) GetPrice() (float64, error) {
	return p.Price, p.Error
}
