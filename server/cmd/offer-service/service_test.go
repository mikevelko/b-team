package main

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsCreatedOfferValid(t *testing.T) {
	//todo: implement tests for rooms once they are implemented
	//todo: also handle pictures
	type args struct {
		offer *CreateOfferRequest
	}
	tests := []struct {
		name  string
		args  args
		check func(t *testing.T, err error)
	}{
		{
			name: "Offer with bad cost per child",
			args: args{
				offer: &CreateOfferRequest{
					Costperchild: decimal.NewFromFloat(-12.12),
					Costperadult: decimal.NewFromFloat(250.00),
					Maxguests:    10,
					Offertitle:   "bad cost per child offer",
				},
			},
			check: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "Offer with bad cost per adult",
			args: args{
				offer: &CreateOfferRequest{
					Costperchild: decimal.NewFromFloat(12.12),
					Costperadult: decimal.NewFromFloat(-50),
					Maxguests:    10,
					Offertitle:   "bad adult cost offer",
				},
			},
			check: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "Offer with no title",
			args: args{
				offer: &CreateOfferRequest{
					Costperchild: decimal.NewFromFloat(120.12),
					Costperadult: decimal.NewFromFloat(500),
					Maxguests:    10,
				},
			},
			check: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "Offer with bad max guests number",
			args: args{
				offer: &CreateOfferRequest{
					Costperchild: decimal.NewFromFloat(120.12),
					Costperadult: decimal.NewFromFloat(500),
					Offertitle:   "negative room",
					Maxguests:    -2137,
				},
			},
			check: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "Valid offer",
			args: args{
				offer: &CreateOfferRequest{
					Costperchild: decimal.NewFromFloat(120.12),
					Costperadult: decimal.NewFromFloat(500),
					Maxguests:    5,
					Offertitle:   "Perfectly valid offer",
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
