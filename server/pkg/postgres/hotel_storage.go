package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

// HotelStorage is responsible for storing and retrieving offers
type HotelStorage struct {
	connPool *pgxpool.Pool
}

// this is very nice way of ensuring, that *HotelStorage{} implements bookly.HotelStorage.
var _ bookly.HotelStorage = &HotelStorage{}

// NewHotelStorage initializes OfferStorage
func NewHotelStorage(pool *pgxpool.Pool) *HotelStorage {
	return &HotelStorage{
		connPool: pool,
	}
}

// CreateHotel implements business logic of creating hotels in database
func (o *HotelStorage) CreateHotel(ctx context.Context, hotel bookly.Hotel) (int64, error) {
	// todo: add pictures
	const query = `
    INSERT INTO hotels(
        hotel_name,
        hotel_desc, 
        city, 
        country
        )
    VALUES ($1,$2,$3,$4)
    RETURNING id
`
	var id int64
	err := o.connPool.QueryRow(ctx, query,
		hotel.Name,
		hotel.Description,
		hotel.City,
		hotel.Country,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("postgres: could not insert hotel: %w", err)
	}
	return id, nil
}

// GetHotelPreviews retrieves hotel previews from storage
func (o *HotelStorage) GetHotelPreviews(ctx context.Context, filter bookly.HotelFilter) ([]*bookly.HotelListing, error) {
	const query = `
    SELECT id, hotel_name, city, country
	FROM hotels
	WHERE
	LOWER(hotel_name) LIKE LOWER($1) AND
	LOWER(city) LIKE LOWER($2) AND
	LOWER(country) LIKE LOWER($3)
`
	hotelList, err := o.connPool.Query(ctx, query,
		"%"+filter.HotelName+"%",
		"%"+filter.City+"%",
		"%"+filter.Country+"%")
	if err != nil {
		return nil, fmt.Errorf("postgres: could not retrieve hotels: %w", err)
	}
	result := []*bookly.HotelListing{}
	defer hotelList.Close()
	for hotelList.Next() {
		hotel := &bookly.HotelListing{}
		// todo: also retrieve preview picture when implemented
		errScan := hotelList.Scan(&hotel.HotelID, &hotel.HotelName, &hotel.City, &hotel.Country)
		if errScan != nil {
			return nil, fmt.Errorf("postgres: could not retrieve hotels: %w", err)
		}
		result = append(result, hotel)
	}
	errFinal := hotelList.Err()
	if errFinal != nil {
		return nil, fmt.Errorf("postgres: could not retrieve hotels: %w", err)
	}
	return result, nil
}

// GetHotelDetails returns details about particular hotel
func (o *HotelStorage) GetHotelDetails(ctx context.Context, hotelID int64) (*bookly.Hotel, error) {
	// todo: retrieve pictures when implemented
	const query = `
    SELECT hotel_name, hotel_desc, city, country
	FROM hotels
	WHERE id=$1
`
	result := &bookly.Hotel{}
	err := o.connPool.QueryRow(ctx, query,
		hotelID).Scan(&result.Name, &result.Description, &result.City, &result.Country)
	if err != nil {
		return nil, fmt.Errorf("postgres: could not retrieve hotel details: %w", err)
	}
	return result, nil
}

// UpdateHotelDetails updates details about existing hotel
func (o *HotelStorage) UpdateHotelDetails(ctx context.Context, hotelID int64, newHotel bookly.Hotel) error {
	// todo: update pictures when implemented
	const query = `
    UPDATE hotels
	SET 
	hotel_name=$2,hotel_desc=$3
	WHERE id=$1
`
	_, err := o.connPool.Exec(ctx, query, hotelID, newHotel.Name, newHotel.Description)
	if err != nil {
		return fmt.Errorf("postgres: could not update hotel details: %w", err)
	}
	return nil
}
