package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/stretchr/testify/assert"

	"gopkg.in/guregu/null.v3"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/testutils"
	"github.com/stretchr/testify/require"
)

func CleanTestUserStorage(t *testing.T, pool *pgxpool.Pool, ctx context.Context) {
	queries := []string{
		"DELETE FROM users",
	}
	for _, q := range queries {
		_, err := pool.Exec(ctx, q)
		require.NoError(t, err)
	}
}

func TestUserStorage_AddUserForce(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewUserStorage(initDb(t))
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
	_, err := storage.AddUserForce(ctx, correctUser, "haslo")
	require.NoError(t, err)
}

func TestUserStorage_GetUser(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewUserStorage(initDb(t))

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
	storage := NewUserStorage(initDb(t))

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

	for i, u := range users {
		id, errAdd := storage.AddUserForce(ctx, *u, "abc1")
		require.NoError(t, errAdd)
		users[i].ID = id
	}

	err := storage.UpdateUserInformation(ctx, users[0].ID, "adam2", "adam4@rmail.com")
	require.Error(t, err)

	err = storage.UpdateUserInformation(ctx, users[0].ID, "adam4", "adam2@rmail.com")
	require.Error(t, err)

	err = storage.UpdateUserInformation(ctx, users[0].ID, "adam4", "adam4@rmail.com")
	require.NoError(t, err)

	err = storage.UpdateUserInformation(ctx, users[0].ID, "adam5", "adam4@rmail.com")
	require.NoError(t, err)

	err = storage.UpdateUserInformation(ctx, users[0].ID, "adam5", "adam5@rmail.com")
	require.NoError(t, err)
}

func TestUserStorage_UserVerify(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewUserStorage(initDb(t))
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
	CleanTestUserStorage(t, storage.connPool, ctx)

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
