package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

// SessionStorage is responsible for storing and retrieving user sessions
type SessionStorage struct {
	connPool *pgxpool.Pool
	// global token expiration time in minutes
	expirationTime time.Duration
}

var _ bookly.SessionStorage = &SessionStorage{}

// NewSessionStorage initializes SessionStorage
func NewSessionStorage(conf Config, expireTime time.Duration) (*SessionStorage, func(), error) {
	pool, cleanup, err := newPool(conf)
	if err != nil {
		return nil, nil, fmt.Errorf("postgres: could not intitialize postgres pool: %w", err)
	}
	storage := &SessionStorage{
		connPool:       pool,
		expirationTime: expireTime,
	}
	return storage, cleanup, nil
}

// CreateNew update
func (o *SessionStorage) CreateNew(ctx context.Context, userID int64) (bookly.Token, error) {
	const query = `
    INSERT INTO sessions(
	creation_date,
	expire_date,
	user_id)
	VALUES($1,$2,$3)
`
	creationTime := time.Now()
	expireTime := creationTime.Local().Add(o.expirationTime)
	_, err := o.connPool.Exec(ctx, query, creationTime, expireTime, userID)
	if err != nil {
		return bookly.Token{}, fmt.Errorf("postgres: could not create new token: %w", err)
	}
	token := bookly.Token{ID: userID, CreatedAt: creationTime.Format(time.RFC3339)}
	return token, nil
}

// Update updates or closes session if expired
func (o *SessionStorage) Update(ctx context.Context, Token bookly.Token) error {
	const query = `
    SELECT id,expire_date FROM sessions
	WHERE user_id=$1
`
	const updateQuery = `
    UPDATE sessions
	SET expire_date=$1
	WHERE id=$2
`
	const deleteQuery = `
    DELETE FROM sessions
	WHERE id=$1
`
	var id int64
	var expireDate time.Time
	nowTime := time.Now()
	err := o.connPool.QueryRow(ctx, query, Token.ID).Scan(&id, &expireDate)
	if err != nil {
		if err != pgx.ErrNoRows {
			return fmt.Errorf("postgres: could not find token: %w", err)
		}
		return bookly.ErrSessionNotFound
	}
	if nowTime.Before(expireDate) {
		newExpireTime := nowTime.Local().Add(o.expirationTime)
		_, errUpdate := o.connPool.Exec(ctx, updateQuery, newExpireTime, id)
		if errUpdate != nil {
			return fmt.Errorf("postgres: could not update token: %w", errUpdate)
		}
		return nil
	}
	_, errDelete := o.connPool.Exec(ctx, deleteQuery, id)
	if errDelete != nil {
		return fmt.Errorf("postgres: could not delete token: %w", errDelete)
	}
	return bookly.ErrSessionExpired
}

// GetSession gets session data based on token
func (o *SessionStorage) GetSession(ctx context.Context, token bookly.Token) (*bookly.Session, error) {
	const query = `
    SELECT hotel_id FROM users
	WHERE id=$1
`
	var hotelID *int64
	err := o.connPool.QueryRow(ctx, query, token.ID).Scan(&hotelID)
	if err != nil {
		return nil, fmt.Errorf("postgres: could not get session: %w", err)
	}
	session := &bookly.Session{UserID: token.ID, HotelID: 0}
	if hotelID != nil {
		session.HotelID = *hotelID
	}
	return session, nil
}
