package service

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"

	"mensageria-go/internal/config"
)

type StripeService struct{}

func NewStripeService(cfg config.Config) *StripeService {

	stripe.Key = cfg.StripeKey

	return &StripeService{}
}

func (s *StripeService) Pagar(valor float64) error {

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(valor * 100)),
		Currency: stripe.String(string(stripe.CurrencyBRL)),
	}

	_, err := paymentintent.New(params)

	return err
}