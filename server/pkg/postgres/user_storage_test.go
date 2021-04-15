package postgres

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"gopkg.in/guregu/null.v3"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/testutils"
	"github.com/stretchr/testify/require"
)

func TestUserStorage_AddUserForce(t *testing.T) {
	testutils.SetIntegration(t)
	storage, cleanup, err := NewUserStorage(conf)
	require.NoError(t, err)
	t.Cleanup(cleanup)
	correctUser := bookly.User{
		ID:        0,
		FirstName: "Adam",
		Surname:   "Nowak",
		Email:     "adam@rmail.com",
		UserName:  "adam1",
		HotelID:   null.Int{},
		UserRole:  bookly.UserRoleClientCustomer,
	}
	ctx := context.Background()
	_, err = storage.AddUserForce(ctx, correctUser, "haslo")
	require.NoError(t, err)
}

func TestUserStorage_GetUser(t *testing.T) {
	testutils.SetIntegration(t)
	storage, cleanup, err := NewUserStorage(conf)
	require.NoError(t, err)
	t.Cleanup(cleanup)
	var users []*bookly.User
	users = append(users, &bookly.User{
		ID:        0,
		FirstName: "Adam",
		Surname:   "Nowak",
		Email:     "adam1@rmail.com",
		UserName:  "adam1",
		HotelID:   null.Int{},
		UserRole:  bookly.UserRoleClientCustomer,
	})
	users = append(users, &bookly.User{
		ID:        0,
		FirstName: "Adam",
		Surname:   "Nowak",
		Email:     "adam2@rmail.com",
		UserName:  "adam2",
		HotelID:   null.Int{},
		UserRole:  bookly.UserRoleClientCustomer,
	})
	users = append(users, &bookly.User{
		ID:        0,
		FirstName: "Adam",
		Surname:   "Nowak",
		Email:     "adam3@rmail.com",
		UserName:  "adam3",
		HotelID:   null.Int{},
		UserRole:  bookly.UserRoleClientCustomer,
	})

	ctx := context.Background()
	CleanTestStorage(t, storage.connPool, ctx)

	for i, u := range users {
		id, errAdd := storage.AddUserForce(ctx, *u, "abc1")
		require.NoError(t, errAdd)
		users[i].ID = id
	}

	result, errGet := storage.GetUser(ctx, users[0].ID)
	require.NoError(t, errGet)

	assert.Equal(t, result, *users[0])

	_, errGet2 := storage.GetUser(ctx, 2850)
	require.Error(t, errGet2)
}

func TestUserStorage_UpdateUserInformationUser(t *testing.T) {
	testutils.SetIntegration(t)
	storage, cleanup, err := NewUserStorage(conf)
	require.NoError(t, err)
	t.Cleanup(cleanup)
	var users []*bookly.User
	users = append(users, &bookly.User{
		ID:        0,
		FirstName: "Adam",
		Surname:   "Nowak",
		Email:     "adam1@rmail.com",
		UserName:  "adam1",
		HotelID:   null.Int{},
		UserRole:  bookly.UserRoleClientCustomer,
	})
	users = append(users, &bookly.User{
		ID:        0,
		FirstName: "Adam",
		Surname:   "Nowak",
		Email:     "adam2@rmail.com",
		UserName:  "adam2",
		HotelID:   null.Int{},
		UserRole:  bookly.UserRoleClientCustomer,
	})
	users = append(users, &bookly.User{
		ID:        0,
		FirstName: "Adam",
		Surname:   "Nowak",
		Email:     "adam3@rmail.com",
		UserName:  "adam3",
		HotelID:   null.Int{},
		UserRole:  bookly.UserRoleClientCustomer,
	})

	ctx := context.Background()
	CleanTestStorage(t, storage.connPool, ctx)

	for i, u := range users {
		id, errAdd := storage.AddUserForce(ctx, *u, "abc1")
		require.NoError(t, errAdd)
		users[i].ID = id
	}

	err = storage.UpdateUserInformation(ctx, users[0].ID, "adam2", "adam4@rmail.com")
	require.Error(t, err)

	err = storage.UpdateUserInformation(ctx, users[0].ID, "adam4", "adam2@rmail.com")
	require.Error(t, err)

	err = storage.UpdateUserInformation(ctx, users[0].ID, "adam4", "adam4@rmail.com")
	require.NoError(t, err)
}

func TestUserStorage_UserVerify(t *testing.T) {
	testutils.SetIntegration(t)
	storage, cleanup, err := NewUserStorage(conf)
	require.NoError(t, err)
	t.Cleanup(cleanup)
	var users []*bookly.User
	users = append(users, &bookly.User{
		ID:        0,
		FirstName: "Adam",
		Surname:   "Nowak",
		Email:     "adam1@rmail.com",
		UserName:  "adam1",
		HotelID:   null.Int{},
		UserRole:  bookly.UserRoleClientCustomer,
	})
	users = append(users, &bookly.User{
		ID:        0,
		FirstName: "Adam",
		Surname:   "Nowak",
		Email:     "adam2@rmail.com",
		UserName:  "adam2",
		HotelID:   null.Int{},
		UserRole:  bookly.UserRoleClientCustomer,
	})
	users = append(users, &bookly.User{
		ID:        0,
		FirstName: "Adam",
		Surname:   "Nowak",
		Email:     "adam3@rmail.com",
		UserName:  "adam3",
		HotelID:   null.Int{},
		UserRole:  bookly.UserRoleClientCustomer,
	})

	ctx := context.Background()
	CleanTestStorage(t, storage.connPool, ctx)

	for i, u := range users {
		id, errAdd := storage.AddUserForce(ctx, *u, "abc1")
		require.NoError(t, errAdd)
		users[i].ID = id
	}

	result, user, errVerify := storage.UserVerify(ctx, "adam1", "abc1")
	require.NoError(t, errVerify)

	assert.Equal(t, user, *users[0])
	assert.Equal(t, result, true)

	result, user, errVerify = storage.UserVerify(ctx, "adam1", "abc")
	require.NoError(t, errVerify)

	assert.Equal(t, user, bookly.User{})
	assert.Equal(t, result, false)
}
