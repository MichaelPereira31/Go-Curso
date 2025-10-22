package tax

import "github.com/stretchr/testify/mock"

type TaxRepository struct {
	mock.Mock
}

func (m *TaxRepository) SaveTax(amount float64) error {
	args := m.Called(amount)
	return args.Error(0)
}