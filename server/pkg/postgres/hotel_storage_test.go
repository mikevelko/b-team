package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CleanHotelTestStorage(t *testing.T, pool *pgxpool.Pool, ctx context.Context) {
	queries := []string{
		"DELETE FROM hotels",
	}
	for _, q := range queries {
		_, err := pool.Exec(ctx, q)
		require.NoError(t, err)
	}
}

func TestHotelStorage_CreateHotel(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewHotelStorage(initDb(t))
	hotel := bookly.Hotel{
		Name:        "Novotel",
		Description: "Live costam",
		City:        "Warsaw",
		Country:     "Poland",
	}
	ctx := context.Background()
	CleanHotelTestStorage(t, storage.connPool, ctx)

	_, err := storage.CreateHotel(ctx, hotel)
	require.NoError(t, err)
}

func TestHotelStorage_GetHotelPreviews(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewHotelStorage(initDb(t))
	ctx := context.Background()
	CleanHotelTestStorage(t, storage.connPool, ctx)

	h1 := bookly.Hotel{
		Name:        "Novotel",
		Description: "Live costam",
		City:        "New York",
		Country:     "USA",
	}
	h2 := bookly.Hotel{
		Name:        "1Star Hotel",
		Description: "Never gonna give you up",
		City:        "Los Angeles",
		Country:     "USA",
	}
	h3 := bookly.Hotel{
		Name:        "Kangaroo Hotel",
		Description: "Never gonna let you down",
		City:        "Sydney",
		Country:     "Australia",
	}
	h1Id, err1 := storage.CreateHotel(ctx, h1)
	require.NoError(t, err1)
	h2Id, err2 := storage.CreateHotel(ctx, h2)
	require.NoError(t, err2)
	h3Id, err3 := storage.CreateHotel(ctx, h3)
	require.NoError(t, err3)

	emptyList, errEmpty := storage.GetHotelPreviews(ctx, bookly.HotelFilter{City: "Alabama"})
	require.NoError(t, errEmpty)
	assert.Empty(t, emptyList)

	nameList, errName := storage.GetHotelPreviews(ctx, bookly.HotelFilter{HotelName: "Kangaroo"})
	require.NoError(t, errName)
	require.Equal(t, 1, len(nameList))
	assert.Equal(t, h3Id, nameList[0].HotelID)

	cityList, errCity := storage.GetHotelPreviews(ctx, bookly.HotelFilter{City: "New"})
	require.NoError(t, errCity)
	require.Equal(t, 1, len(cityList))
	assert.Equal(t, h1Id, cityList[0].HotelID)

	countryList, errCountry := storage.GetHotelPreviews(ctx, bookly.HotelFilter{Country: "USA"})
	require.NoError(t, errCountry)
	require.Equal(t, 2, len(countryList))
	assert.Equal(t, h1Id, countryList[0].HotelID)
	assert.Equal(t, h2Id, countryList[1].HotelID)
}

func TestHotelStorage_GetHotelDetails(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewHotelStorage(initDb(t))
	ctx := context.Background()
	CleanHotelTestStorage(t, storage.connPool, ctx)

	h1 := bookly.Hotel{
		Name:        "Novotel",
		Description: "Live costam",
		City:        "New York",
		Country:     "USA",
	}
	id, errAdd := storage.CreateHotel(ctx, h1)
	require.NoError(t, errAdd)

	_, errDetailsEmpty := storage.GetHotelDetails(ctx, 1000000)
	require.Error(t, errDetailsEmpty)

	details, errDetails := storage.GetHotelDetails(ctx, id)
	require.NoError(t, errDetails)
	// todo: tests for pictures
	assert.Equal(t, h1.Name, details.Name)
	assert.Equal(t, h1.Description, details.Description)
	assert.Equal(t, h1.City, details.City)
	assert.Equal(t, h1.Country, details.Country)
}

func TestHotelStorage_UpdateHotelDetails(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewHotelStorage(initDb(t))
	ctx := context.Background()
	CleanHotelTestStorage(t, storage.connPool, ctx)

	pre := bookly.Hotel{
		Name:        "Novotel",
		Description: "Live costam",
		City:        "New York",
		Country:     "USA",
	}
	post := bookly.Hotel{
		Name:        "Kangaroo",
		Description: "Live costam",
		City:        "New York",
		Country:     "USA",
	}
	id, errAdd := storage.CreateHotel(ctx, pre)
	require.NoError(t, errAdd)

	errEdit := storage.UpdateHotelDetails(ctx, id, post)
	require.NoError(t, errEdit)

	nameList, errName := storage.GetHotelPreviews(ctx, bookly.HotelFilter{HotelName: "Kangaroo"})
	require.NoError(t, errName)
	require.Equal(t, 1, len(nameList))
}
