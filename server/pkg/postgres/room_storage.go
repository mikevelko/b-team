package postgres

import (
	"context"
	"fmt"
	"log"

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
func NewRoomStorage(pool *pgxpool.Pool) *RoomStorage {
	return &RoomStorage{
		connPool: pool,
	}
}

// CreateRoom implements business logic of create room
func (r *RoomStorage) CreateRoom(ctx context.Context, room bookly.Room, hotelID int64) (int64, error) {
	const queryCheck = `
    SELECT *
	FROM rooms
	WHERE room_number = $1 AND hotel_id = $2
`
	list, err := r.connPool.Exec(ctx, queryCheck,
		room.RoomNumber,
		hotelID,
	)
	if err != nil {
		return 0, fmt.Errorf("postgres: could not insert offer: %w", err)
	}
	if list.RowsAffected() > 0 {
		return 0, bookly.ErrRoomAlreadyExists
	}

	const query = `
    INSERT INTO rooms(
        room_number,
		hotel_id
        )
    VALUES ($1,$2)
    RETURNING id;
`
	var id int64
	err = r.connPool.QueryRow(ctx, query,
		room.RoomNumber,
		hotelID,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("postgres: could not insert offer: %w", err)
	}
	return id, nil
}

// DeleteRoom implements business logic of delete room
func (r *RoomStorage) DeleteRoom(ctx context.Context, roomID int64, hotelID int64) error {
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
			return bookly.ErrRoomNotFound
		}
		return fmt.Errorf("postgres: could not find room with delete func: %w", err)
	}

	if room.HotelID != hotelID {
		return bookly.ErrRoomNotOwnedByHotel
	}

	// todo: Find in documentation example with occupy room

	const queryDelete = `
    DELETE
	FROM rooms
	WHERE id = $1
`
	_, err = r.connPool.Exec(ctx, queryDelete,
		roomID,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return bookly.ErrRoomNotFound
		}
		return fmt.Errorf("postgres: could not delete room: %w", err)
	}

	return nil
}

// GetAllHotelRooms implements business logic of getting all ofer belong to hotel
func (r *RoomStorage) GetAllHotelRooms(ctx context.Context, hotelID int64) ([]*bookly.Room, error) {
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
		if err != nil {
			return nil, fmt.Errorf("postgres: could not insert offer: %w", err)
		}

		room.OfferID, err = r.GetOffersRelatedWithRoom(ctx, room.ID)
		if err != nil {
			return nil, fmt.Errorf("postgres: could not insert offer: %w", err)
		}

		result = append(result, room)
	}
	errFinal := list.Err()
	if errFinal != nil {
		return nil, fmt.Errorf("postgres: could not retrieve hotel's offers: %w", err)
	}

	return result, nil
}

// GetRoomByName implements getting one room by roomNumber
func (r *RoomStorage) GetRoomByName(ctx context.Context, roomNumber string, hotelID int64) (bookly.Room, error) {
	const queryCheck = `
    SELECT *
	FROM rooms
	WHERE room_number = $1 AND hotel_id = $2
`

	row := r.connPool.QueryRow(ctx, queryCheck,
		roomNumber,
		hotelID,
	)
	room := bookly.Room{}
	err := row.Scan(&room.ID, &room.RoomNumber, &room.HotelID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return bookly.Room{}, bookly.ErrRoomNotFound
		}
		return bookly.Room{}, fmt.Errorf("postgres: could not get room: %w", err)
	}

	room.OfferID, err = r.GetOffersRelatedWithRoom(ctx, room.ID)
	if err != nil {
		return bookly.Room{}, fmt.Errorf("postgres: could not insert offer: %w", err)
	}

	return room, nil
}

// GetRoom implements database feat to get room by ID
func (r *RoomStorage) GetRoom(ctx context.Context, roomID int64) (bookly.Room, error) {
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
			return bookly.Room{}, bookly.ErrRoomNotFound
		}
		return bookly.Room{}, fmt.Errorf("postgres: could not get room: %w", err)
	}

	return room, nil
}

// IsRoomOwnedByHotel implements database feat to get room by ID
func (r *RoomStorage) IsRoomOwnedByHotel(ctx context.Context, roomID int64, hotelID int64) (bool, error) {
	const query = `
    SELECT EXISTS(
		SELECT *
		FROM rooms
		WHERE id = $1 AND hotel_id = $2
	);
`
	row := r.connPool.QueryRow(ctx, query,
		roomID,
		hotelID,
	)
	var exist bool
	err := row.Scan(&exist)
	if err != nil {
		return false, fmt.Errorf("postgres: could not find room: %w", err)
	}
	return exist, nil
}

// GetOffersRelatedWithRoom implements database feat to get offers IDs related with room
func (r *RoomStorage) GetOffersRelatedWithRoom(ctx context.Context, roomID int64) ([]int64, error) {
	const query = `
    SELECT *
	FROM offers_rooms
	WHERE room_id = $1
`

	list, err := r.connPool.Query(ctx, query,
		roomID,
	)
	if err != nil {
		return nil, fmt.Errorf("postgres: could not get offers related with room: %w", err)
	}

	result := []int64{}
	defer list.Close()
	for list.Next() {
		var offerID int64
		var _roomID int64
		err = list.Scan(&offerID, &_roomID)
		result = append(result, offerID)
	}
	errFinal := list.Err()
	if errFinal != nil {
		return nil, fmt.Errorf("postgres: could not get offers related with room: %w", err)
	}

	return result, nil
}

// GetRoomsRelatedWithOffer implements database feat to get rooms IDs related with offer
func (r *RoomStorage) GetRoomsRelatedWithOffer(ctx context.Context, offerID int64) ([]int64, error) {
	const query = `
    SELECT *
	FROM offers_rooms
	WHERE offer_id = $1
`

	list, err := r.connPool.Query(ctx, query,
		offerID,
	)
	if err != nil {
		return nil, fmt.Errorf("postgres: could not get rooms related with offer: %w", err)
	}

	result := []int64{}
	defer list.Close()
	for list.Next() {
		var _offerID int64
		var roomID int64
		err = list.Scan(&_offerID, &roomID)
		result = append(result, roomID)
	}
	errFinal := list.Err()
	if errFinal != nil {
		return nil, fmt.Errorf("postgres: could not get rooms related with offer: %w", err)
	}

	return result, nil
}

// IsExistLinkWithRoomAndOffer implements database feat to check link with room and offer
func (r *RoomStorage) IsExistLinkWithRoomAndOffer(ctx context.Context, offerID int64, roomID int64) (bool, error) {
	const queryCheck = `
    SELECT *
	FROM offers_rooms
	WHERE offer_id = $1 AND room_id = $2
`

	row := r.connPool.QueryRow(ctx, queryCheck,
		offerID,
		roomID,
	)
	var _offerID, _roomID int64
	err := row.Scan(&_offerID, &_roomID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("postgres: could not find link with offer and room: %w", err)
	}
	return true, nil
}

// AddLinkWithRoomAndOffer implements database feat to add link with room and offer
func (r *RoomStorage) AddLinkWithRoomAndOffer(ctx context.Context, offerID int64, roomID int64) error {
	const query = `
    INSERT INTO offers_rooms(
        offer_id,
		room_id
        )
    VALUES ($1,$2)
    RETURNING offer_id;
`
	var id int64
	err := r.connPool.QueryRow(ctx, query,
		offerID,
		roomID,
	).Scan(&id)
	if err != nil {
		return fmt.Errorf("postgres: could not insert offer: %w", err)
	}
	return nil
}

// DeleteLinkWithRoomAndOffer implements database feat to delete link with room and offer
func (r *RoomStorage) DeleteLinkWithRoomAndOffer(ctx context.Context, offerID int64, roomID int64) error {
	const queryDelete = `
    DELETE
	FROM offers_rooms
	WHERE offer_id = $1 AND room_id = $2
`
	log.Println("database func")
	rows, err := r.connPool.Exec(ctx, queryDelete,
		offerID,
		roomID,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Println("database func, exit 1")
			return bookly.ErrLinkOfferRoomNotFound
		}
		return fmt.Errorf("postgres: could not delete record from offers_rooms: %w", err)
	}
	if rows.RowsAffected() == 0 {
		log.Println("database func, exit 2")
		return bookly.ErrLinkOfferRoomNotFound
	}

	return nil
}
