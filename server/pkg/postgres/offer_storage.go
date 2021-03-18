package postgres

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v4/pgxpool"
    "github.com/pw-software-engineering/b-team/server/pkg/rently"
)

//OfferStorage is responsible for storing and retrieving offers
type OfferStorage struct{
    connPool *pgxpool.Pool
}

// this is very nice way of ensuring, that *OfferStorage{} implements rently.OfferStorage.
// if it does not - program won't compile and you'll see red error in IDE
var _ rently.OfferStorage = &OfferStorage{}

//NewOfferStorage initializes OfferStorage
func NewOfferStorage(conf Config) (*OfferStorage, func(), error){
    pool, cleanup, err := newPool(conf)
    if err != nil {
        return nil, nil, fmt.Errorf("postgres: could not intitialize postgres pool: %w", err)
    }
    storage := &OfferStorage{
        connPool: pool,
    }
    return storage, cleanup, nil
}


func (o *OfferStorage) CreateOffer(ctx context.Context, offer *rently.Offer) (int64, error) {
    const query = `
    INSERT INTO offers(
        is_active, 
        offer_title, 
        cost_per_child, 
        cost_par_adult, 
        max_guests, 
        description, 
        offer_preview_picture_url
        )
    VALUES ($1,$2,$3,$4,$5,$6,$7)
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
    ).Scan(&id)
    if err != nil{
        return 0, fmt.Errorf("postgres: could not insert offer: %w", err)
    }
    return id, nil
}

func (o *OfferStorage) DeleteOffer(ctx context.Context, id int) error {
    panic("implement me")
}

func (o *OfferStorage) GetAllOffers(ctx context.Context, hotelID int) ([]*rently.Offer, error) {
    panic("implement me")
}
