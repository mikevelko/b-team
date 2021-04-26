package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"

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

// CreateOffer implements business logic of adding new offer to database
func (o *OfferStorage) CreateOffer(ctx context.Context, offer *bookly.Offer, hotelID int64) (int64, error) {
	const query = `
    INSERT INTO offers(
        is_active,
        offer_title, 
        cost_per_child, 
        cost_per_adult, 
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

// IsOfferActive checks if offer is still marked as active
func (o *OfferStorage) IsOfferActive(ctx context.Context, offerID int64) (bool, error) {
	const queryCheck = `
	SELECT is_active FROM offers
	WHERE id=$1
`
	var isActive bool
	errCheck := o.connPool.QueryRow(ctx, queryCheck,
		offerID,
	).Scan(&isActive)
	if errCheck != nil {
		return false, fmt.Errorf("postgres: could not check if offer is active: %w", errCheck)
	}
	return isActive, nil
}

// IsOfferOwnedByHotel checks if hotel with hotelID is an owner of particular offer
func (o *OfferStorage) IsOfferOwnedByHotel(ctx context.Context, hotelID int64, offerID int64) (bool, error) {
	const queryCheck = `
	SELECT hotel_id FROM offers
	WHERE id=$1
`
	var offerHotelID int64
	errCheck := o.connPool.QueryRow(ctx, queryCheck,
		offerID,
	).Scan(&offerHotelID)
	if errCheck != nil {
		return false, fmt.Errorf("postgres: could not check ownership of offer: %w", errCheck)
	}
	return offerHotelID == hotelID, nil
}

// IsOfferMarkedAsDeleted checks if offer was marked as deleted in 'is_deleted' field
func (o *OfferStorage) IsOfferMarkedAsDeleted(ctx context.Context, offerID int64) (bool, error) {
	const queryCheck = `
	SELECT is_deleted FROM offers
	WHERE id=$1
`
	var isDeleted bool
	errCheck := o.connPool.QueryRow(ctx, queryCheck,
		offerID,
	).Scan(&isDeleted)
	if errCheck != nil {
		return false, fmt.Errorf("postgres: could not check deletion status of offer: %w", errCheck)
	}
	return isDeleted, nil
}

// SetOfferDeletionStatus sets flag is_deleted in offer
func (o *OfferStorage) SetOfferDeletionStatus(ctx context.Context, offerID int64, isDeleted bool) error {
	const query = `
    UPDATE offers
    SET is_deleted = $2 
    WHERE id = $1
`
	_, errOffer := o.connPool.Exec(ctx, query, offerID, isDeleted)
	if errOffer != nil {
		return fmt.Errorf("postgres: could not update offer deletion status: %w", errOffer)
	}
	return nil
}

// GetAllOffers implements logic related to retrieving all offers for given hotel
func (o *OfferStorage) GetAllOffers(ctx context.Context, hotelID int64, isActive *bool) ([]*bookly.Offer, error) {
	const queryAny = `
    SELECT * FROM offers
	WHERE hotel_id = $1 AND is_deleted = false
`
	const queryFilter = `
    SELECT * FROM offers
	WHERE hotel_id = $1 AND is_active = $2 AND is_deleted = false
`
	var list pgx.Rows
	var err error
	if isActive == nil {
		list, err = o.connPool.Query(ctx, queryAny, hotelID)
	} else {
		list, err = o.connPool.Query(ctx, queryFilter, hotelID, *isActive)
	}
	if err != nil {
		return nil, fmt.Errorf("postgres: could not retrieve hotel's offers: %w", err)
	}
	result := []*bookly.Offer{}
	defer list.Close()
	for list.Next() {
		var id int64
		var hID int64
		var deleted bool
		// todo: find better way to ignore those ids if exists
		offer := &bookly.Offer{}
		errScan := list.Scan(&id, &hID, &offer.IsActive, &offer.OfferTitle,
			&offer.CostPerChild, &offer.CostPerAdult, &offer.MaxGuests, &offer.Description, &offer.OfferPreviewPicture, &deleted)
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

// GetSpecificOffer returns specific offer from database
func (o *OfferStorage) GetSpecificOffer(ctx context.Context, offerID int64) (*bookly.Offer, error) {
	const query = `
    SELECT hotel_id, is_active, offer_title, cost_per_child, cost_per_adult, max_guests, description FROM offers
	WHERE id=$1
`
	// todo: retrieve pictures
	r := bookly.Offer{}
	var roomHotelID int64
	err := o.connPool.QueryRow(ctx, query,
		offerID,
	).Scan(&roomHotelID, &r.IsActive, &r.OfferTitle, &r.CostPerChild, &r.CostPerAdult, &r.MaxGuests, &r.Description)
	if err != nil {
		return nil, fmt.Errorf("postgres: could not get specific offer: %w", err)
	}
	return &r, nil
}

// UpdateOfferDetails updates specific offer from database with new details
func (o *OfferStorage) UpdateOfferDetails(ctx context.Context, offerID int64, newOffer bookly.Offer) error {
	const queryUpdate = `
	UPDATE offers
    SET is_active=$2,
	offer_title=$3, 
	cost_per_child=$4, 
	cost_per_adult=$5,
	max_guests=$6,
	description=$7
	WHERE id=$1 
`
	_, errUpdate := o.connPool.Exec(ctx, queryUpdate,
		offerID,
		newOffer.IsActive,
		newOffer.OfferTitle,
		newOffer.CostPerChild,
		newOffer.CostPerAdult,
		newOffer.MaxGuests,
		newOffer.Description,
	)
	// todo: update pictures
	if errUpdate != nil {
		return fmt.Errorf("postgres: could not update offer details: %w", errUpdate)
	}
	return nil
}
