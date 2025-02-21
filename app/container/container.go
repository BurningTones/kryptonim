package container

import (
	"kryptonim/app/services"
)

type Container struct {
	RatesService    services.RatesService
	ExchangeService services.ExchangeService
}

func BuildContainer() *Container {
	return &Container{
		RatesService:    services.NewRatesService(),
		ExchangeService: services.NewExchangeService(),
	}
}
