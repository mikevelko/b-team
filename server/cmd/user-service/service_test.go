package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmailValid(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name  string
		args  args
		check func(t *testing.T, err error)
	}{
		{
			name: "Email with wrong length",
			args: args{email: "sd"},
			check: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "Email with wrong sign",
			args: args{email: "ala@gmail_com"},
			check: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "Valid email",
			args: args{email: "jan123@rmail.com"},
			check: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsEmailValid(tt.args.email)
			tt.check(t, err)
		})
	}
}

func TestIsUserNameValid(t *testing.T) {
	type args struct {
		userName string
	}
	tests := []struct {
		name  string
		args  args
		check func(t *testing.T, err error)
	}{
		{
			name: "UserName with wrong length",
			args: args{userName: "sd"},
			check: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "Valid userName",
			args: args{userName: "adam1"},
			check: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsUserNameValid(tt.args.userName)
			tt.check(t, err)
		})
	}
}
