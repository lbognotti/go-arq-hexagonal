package application_test

import (
	"github.com/LucasMatosBognotti/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Arroz"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = ""
	_, err = product.IsValid()
	require.Equal(t, "disabled", product.Status)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -1
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Arroz"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Arroz"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Arroz"
	product.Status = application.DISABLED
	product.Price = 10

	id := product.GetID()
	require.Equal(t, id, product.ID)
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Arroz"
	product.Status = application.DISABLED
	product.Price = 10

	name := product.GetName()
	require.Equal(t, name, product.Name)
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Arroz"
	product.Status = application.DISABLED
	product.Price = 10

	price := product.GetPrice()
	require.Equal(t, price, product.Price)
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Arroz"
	product.Status = application.DISABLED
	product.Price = 10

	status := product.GetStatus()
	require.Equal(t, status, product.Status)
}
