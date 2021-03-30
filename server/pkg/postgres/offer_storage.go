package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

// OfferStorage is responsible for storing and retrieving offers
type OfferStorage struct {
	connPool *pgxpool.Pool
}

// this is very nice way of ensuring, that *OfferStorage{} implements bookly.OfferStorage.
// if it does not - program won't compile and you'll see red error in IDE
var _ bookly.OfferStorage = &OfferStorage{}

// NewOfferStorage initializes OfferStorage
func NewOfferStorage(conf Config) (*OfferStorage, func(), error) {
	pool, cleanup, err := newPool(conf)
	if err != nil {
		return nil, nil, fmt.Errorf("postgres: could not intitialize postgres pool: %w", err)
	}
	storage := &OfferStorage{
		connPool: pool,
	}
	return storage, cleanup, nil
}

// CreateOffer implements business logic of
func (o *OfferStorage) CreateOffer(ctx context.Context, offer *bookly.Offer, hotelID int) (int64, error) {
	const query = `
    INSERT INTO offers(
        is_active,
        offer_title, 
        cost_per_child, 
        cost_par_adult, 
        max_guests, 
        description, 
        offer_preview_picture_url,
		hotel_id
        )
    VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
    RETURNING id;
`
	var id int64
	err := o.connPool.QueryRow(ctx, query,
		offer.IsActive,
		offer.OfferTitle,
		offer.CostPerChild,
		offer.CostPerAdult,
		offer.MaxGuests,
		offer.Description,
		offer.OfferPreviewPicture,
		hotelID,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("postgres: could not insert offer: %w", err)
	}
	return id, nil
}

// UpdateOfferStatus implements business logic of updating offer status
func (o *OfferStorage) UpdateOfferStatus(ctx context.Context, id int64, isActive bool) error {
	const query = `
    UPDATE offers
    SET is_active = $2 
    WHERE id = $1
`
	// todo: implement changing isDeleted when appropriate field will be available in database
	// todo: also when marking as deleted, remove offer-room relations from db when they will be implemented
	_, err := o.connPool.Exec(ctx, query, id, isActive)
	if err != nil {
		return fmt.Errorf("postgres: could not update offer status: %w", err)
	}
	return nil
}

// GetAllOffers implements business logic related to retrieving all offers for given hotel
func (o *OfferStorage) GetAllOffers(ctx context.Context, hotelID int) ([]*bookly.Offer, error) {
	const query = `
    SELECT * FROM offers
	WHERE hotel_id = $1
`
	list, err := o.connPool.Query(ctx, query, hotelID)
	if err != nil {
		return nil, fmt.Errorf("postgres: could not retrieve hotel's offers: %w", err)
	}
	result := []*bookly.Offer{}
	defer list.Close()
	for list.Next() {
		var id int64
		var hID int64
		// todo: find better way to ignore those ids if exists
		offer := &bookly.Offer{}
		errScan := list.Scan(&id, &hID, &offer.IsActive, &offer.OfferTitle,
			&offer.CostPerChild, &offer.CostPerAdult, &offer.MaxGuests, &offer.Description, &offer.OfferPreviewPicture)
		if errScan != nil {
			return nil, fmt.Errorf("postgres: could not retrieve hotel's offers: %w", err)
		}
		result = append(result, offer)
	}
	errFinal := list.Err()
	if errFinal != nil {
		return nil, fmt.Errorf("postgres: could not retrieve hotel's offers: %w", err)
	}
	return result, nil
}
