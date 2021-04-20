package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

// RoomStorage is responsible for storing and retrieving offers
type RoomStorage struct {
	connPool *pgxpool.Pool
}

// this is very nice way of ensuring, that *RoomStorage{} implements bookly.RoomStorage.
// if it does not - program won't compile and you'll see red error in IDE
var _ bookly.RoomStorage = &RoomStorage{}

// NewRoomStorage initializes RoomStorage
func NewRoomStorage(conf Config) (*RoomStorage, func(), error) {
	pool, cleanup, err := newPool(conf)
	if err != nil {
		return nil, nil, fmt.Errorf("postgres: could not intitialize postgres pool: %w", err)
	}
	storage := &RoomStorage{
		connPool: pool,
	}
	return storage, cleanup, nil
}

// CreateRoom implements business logic of create room
func (r RoomStorage) CreateRoom(ctx context.Context, room bookly.Room, hotelID int64) (int64, error) {
	const query = `
    INSERT INTO rooms(
        room_number,
		hotel_id
        )
    VALUES ($1,$2)
    RETURNING id;
`
	var id int64
	err := r.connPool.QueryRow(ctx, query,
		room.RoomNumber,
		hotelID,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("postgres: could not insert offer: %w", err)
	}
	return id, nil
}

// DeleteRoom implements business logic of delete room
func (r RoomStorage) DeleteRoom(ctx context.Context, roomID int64, hotelID int64) (bookly.DeleteResponse, error) {
	const queryCheck = `
    SELECT *
	FROM rooms
	WHERE id = $1
`

	row := r.connPool.QueryRow(ctx, queryCheck,
		roomID,
	)
	room := bookly.Room{}
	err := row.Scan(&room.ID, &room.RoomNumber, &room.HotelID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return bookly.RoomNotFound, nil
		}
		return bookly.RoomError, fmt.Errorf("postgres: could not insert offer: %w", err)
	}

	if room.HotelID != hotelID {
		return bookly.RoomNotBelongToHotel, nil
	}

	const queryDelete = `
    DELETE
	FROM rooms
	WHERE id = $1
`
	_, err = r.connPool.Exec(ctx, queryDelete,
		roomID,
	)
	if err == pgx.ErrNoRows {
		return bookly.RoomNotFound, nil
	}

	return bookly.RoomSuccess, nil
}

// GetAllHotelRooms implements business logic of getting all ofer belong to hotel
func (r RoomStorage) GetAllHotelRooms(ctx context.Context, hotelID int) ([]*bookly.Room, error) {
	const query = `
    SELECT *
	FROM rooms
	WHERE hotel_id = $1
`

	list, err := r.connPool.Query(ctx, query,
		hotelID,
	)
	if err != nil {
		return nil, fmt.Errorf("postgres: could not insert offer: %w", err)
	}

	result := []*bookly.Room{}
	defer list.Close()
	for list.Next() {
		room := &bookly.Room{}
		err = list.Scan(&room.ID, &room.RoomNumber, &room.HotelID)
		result = append(result, room)
	}
	errFinal := list.Err()
	if errFinal != nil {
		return nil, fmt.Errorf("postgres: could not retrieve hotel's offers: %w", err)
	}

	return result, nil
}