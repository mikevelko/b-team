package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

// ReservationStorage is responsible for storing and retrieving Reservations
type ReservationStorage struct {
	connPool *pgxpool.Pool
}

// this is very nice way of ensuring, that *ReservationStorage{} implements bookly.ReservationStorage.
// if it does not - program won't compile and you'll see red error in IDE
var _ bookly.ReservationStorage = &ReservationStorage{}

// NewReservationStorage initializes ReservationStorage
func NewReservationStorage(pool *pgxpool.Pool) *ReservationStorage {
	return &ReservationStorage{
		connPool: pool,
	}
}

// CreateReservation creates reservation in database
func (r *ReservationStorage) CreateReservation(ctx context.Context, reservation *bookly.Reservation) (int64, error) {
	const query = `
    INSERT INTO reservations(
        client_id,
        hotel_id, 
        offer_id, 
        from_time, 
        to_time, 
        child_count, 
        adult_count
        )
    VALUES ($1,$2,$3,$4,$5,$6,$7)
    RETURNING id;
`
	var id int64
	err := r.connPool.QueryRow(ctx, query,
		reservation.ClientID,
		reservation.HotelID,
		reservation.OfferID,
		reservation.FromTime,
		reservation.ToTime,
		reservation.ChildCount,
		reservation.AdultCount,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("postgres: could not insert reservation: %w", err)
	}
	return id, nil
}

// DeleteReservation deletes reservation from database
func (r *ReservationStorage) DeleteReservation(ctx context.Context, reservationID int64) error {
	const query = `
    DELETE FROM reservations
	WHERE id=$1
`
	_, err := r.connPool.Exec(ctx, query,
		reservationID)
	if err != nil {
		return fmt.Errorf("postgres: could not delete reservation: %w", err)
	}
	return nil
}

// GetClientReservations retrieves reservation made by particular client from database
func (r *ReservationStorage) GetClientReservations(ctx context.Context, clientID int64) ([]*bookly.Reservation, error) {
	const query = `
    SELECT * FROM reservations
	WHERE client_id=$1
`

	list, err := r.connPool.Query(ctx, query, clientID)
	if err != nil {
		return nil, fmt.Errorf("postgres: could not retrieve client reservations: %w", err)
	}
	result := []*bookly.Reservation{}
	defer list.Close()
	for list.Next() {
		reser := &bookly.Reservation{}
		errScan := list.Scan(&reser.ID, &reser.ClientID, &reser.HotelID, &reser.OfferID, &reser.FromTime,
			&reser.ToTime, &reser.ChildCount, &reser.AdultCount)
		if errScan != nil {
			return nil, fmt.Errorf("postgres: could not retrieve client reservations: %w", err)
		}
		result = append(result, reser)
	}
	errFinal := list.Err()
	if errFinal != nil {
		return nil, fmt.Errorf("postgres: could not retrieve client reservations: %w", err)
	}
	return result, nil
}

// GetHotelReservations retrieves reservation made in particular hotel
func (r *ReservationStorage) GetHotelReservations(ctx context.Context, hotelID int64) ([]*bookly.Reservation, error) {
	const query = `
    SELECT * FROM reservations
	WHERE hotel_id=$1
`
	list, err := r.connPool.Query(ctx, query, hotelID)
	if err != nil {
		return nil, fmt.Errorf("postgres: could not retrieve hotel reservations: %w", err)
	}
	result := []*bookly.Reservation{}
	defer list.Close()
	for list.Next() {
		reser := &bookly.Reservation{}
		errScan := list.Scan(&reser.ID, &reser.ClientID, &reser.HotelID, &reser.OfferID, &reser.FromTime,
			&reser.ToTime, &reser.ChildCount, &reser.AdultCount)
		if errScan != nil {
			return nil, fmt.Errorf("postgres: could not retrieve hotel reservations: %w", err)
		}
		result = append(result, reser)
	}
	return result, nil
}

// IsReservationOwnedByClient checks if reservation was made by particular client
func (r *ReservationStorage) IsReservationOwnedByClient(ctx context.Context, clientID int64, reservationID int64) (bool, error) {
	const queryCheck = `
	SELECT client_id FROM reservations
	WHERE id=$1
`
	var reservationClientID int64
	errCheck := r.connPool.QueryRow(ctx, queryCheck,
		reservationID,
	).Scan(&reservationClientID)
	if errCheck != nil {
		return false, fmt.Errorf("postgres: could not healthCheck ownership of reservation: %w", errCheck)
	}
	return clientID == reservationClientID, nil
}

// GetSpecificReservation retrieves specific reservation from database
func (r *ReservationStorage) GetSpecificReservation(ctx context.Context, reservationID int64) (*bookly.Reservation, error) {
	const query = `
    SELECT * FROM reservations
	WHERE id=$1
`
	reservation := bookly.Reservation{}
	err := r.connPool.QueryRow(ctx, query,
		reservationID,
	).Scan(&reservation.ID, &reservation.ClientID, &reservation.HotelID, &reservation.OfferID,
		&reservation.FromTime, &reservation.ToTime, &reservation.ChildCount, &reservation.AdultCount)
	if err != nil {
		return nil, fmt.Errorf("postgres: could not get specific reservation: %w", err)
	}
	return &reservation, nil
}
