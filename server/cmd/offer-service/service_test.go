package main

import (
	"testing"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestIsCreatedOfferValid(t *testing.T) {
	// todo: implement tests for rooms once they are implemented
	// todo: also handle pictures
	type args struct {
		offer *bookly.Offer
	}
	tests := []struct {
		name  string
		args  args
		check func(t *testing.T, err error)
	}{
		{
			name: "Offer with bad cost per child",
			args: args{
				offer: &bookly.Offer{
					CostPerChild: decimal.NewFromFloat(-12.12),
					CostPerAdult: decimal.NewFromFloat(250.00),
					MaxGuests:    10,
					OfferTitle:   "bad cost per child offer",
				},
			},
			check: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "Offer with bad cost per adult",
			args: args{
				offer: &bookly.Offer{
					CostPerChild: decimal.NewFromFloat(12.12),
					CostPerAdult: decimal.NewFromFloat(-50),
					MaxGuests:    10,
					OfferTitle:   "bad adult cost offer",
				},
			},
			check: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "Offer with no title",
			args: args{
				offer: &bookly.Offer{
					CostPerChild: decimal.NewFromFloat(120.12),
					CostPerAdult: decimal.NewFromFloat(500),
					MaxGuests:    10,
				},
			},
			check: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "Offer with bad max guests number",
			args: args{
				offer: &bookly.Offer{
					CostPerChild: decimal.NewFromFloat(120.12),
					CostPerAdult: decimal.NewFromFloat(500),
					OfferTitle:   "negative room",
					MaxGuests:    -2137,
				},
			},
			check: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "Valid offer",
			args: args{
				offer: &bookly.Offer{
					CostPerChild: decimal.NewFromFloat(120.12),
					CostPerAdult: decimal.NewFromFloat(500),
					MaxGuests:    5,
					OfferTitle:   "Perfectly valid offer",
				},
			},
			check: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsCreatedOfferValid(tt.args.offer)
			tt.check(t, err)
		})
	}
}
