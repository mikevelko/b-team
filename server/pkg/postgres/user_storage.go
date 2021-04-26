package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

// UserStorage is responsible for storing and retrieving user
type UserStorage struct {
	connPool *pgxpool.Pool
}

// this is very nice way of ensuring, that *UserStorage{} implements bookly.UserStorage.
// if it does not - program won't compile and you'll see red error in IDE
var _ bookly.UserStorage = &UserStorage{}

// NewUserStorage initializes UserStorage
func NewUserStorage(conf Config) (*UserStorage, func(), error) {
	pool, cleanup, err := newPool(conf)
	if err != nil {
		return nil, nil, fmt.Errorf("postgres: could not intitialize postgres pool: %w", err)
	}
	storage := &UserStorage{
		connPool: pool,
	}
	return storage, cleanup, nil
}

// AddUserForce implements force adding without business logic
func (u *UserStorage) AddUserForce(ctx context.Context, user bookly.User, password string) (int64, error) {
	const query = `
    INSERT INTO users(
    	first_name,
    	surname,
    	email,
    	user_name,
    	password,
    	hotel_id,
		user_role
		)
    VALUES ($1,$2,$3,$4,$5,$6,$7)
    RETURNING id;
`
	var id int64
	err := u.connPool.QueryRow(ctx, query,
		user.FirstName,
		user.Surname,
		user.Email,
		user.UserName,
		password,
		user.HotelID,
		user.UserRole,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("postgres: could not insert user: %w", err)
	}
	return id, nil
}

// GetUser implements business logic of getting information about user
func (u *UserStorage) GetUser(ctx context.Context, userID int64) (bookly.User, error) {
	const queryAny = `
    SELECT * FROM users
	WHERE id = $1
`

	var list pgx.Rows
	var err error
	list, err = u.connPool.Query(ctx, queryAny, userID)

	if err != nil {
		return bookly.User{}, fmt.Errorf("postgres: could not retrieve user info: %w", err)
	}

	user := bookly.User{}

	count := 0

	defer list.Close()
	for list.Next() {
		var pass string
		// todo: find better way to ignore pass if exists

		errScan := list.Scan(&user.ID, &user.FirstName, &user.Surname, &user.Email,
			&user.UserName, &pass, &user.HotelID, &user.UserRole)
		if errScan != nil {
			return bookly.User{}, fmt.Errorf("postgres: could not retrieve user info: %w", err)
		}

		count++
	}
	errFinal := list.Err()
	if errFinal != nil {
		return bookly.User{}, fmt.Errorf("postgres: could not retrieve user info: %w", err)
	}
	if count < 1 {
		return bookly.User{}, fmt.Errorf("postgres: could not retrieve user info: There is no user with this ID")
	}
	return user, nil
}

// UpdateUserInformation implements business logic of update user user name ad e-mail
func (u *UserStorage) UpdateUserInformation(ctx context.Context, id int64, userName string, email string) error {
	user, err := u.GetUser(ctx, id)
	if err != nil {
		return fmt.Errorf("postgres: could not update user info: %w", err)
	}

	const queryName = `
    SELECT *
	FROM users
	WHERE user_name = $1
`
	list, err := u.connPool.Exec(ctx, queryName, userName)
	if err != nil {
		return fmt.Errorf("postgres: could not update user info: %w", err)
	}
	if list.RowsAffected() != 0 && user.UserName != userName {
		return fmt.Errorf("postgres: could not update user info: There is another user with this user_name")
	}

	const queryEmail = `
    SELECT *
	FROM users
	WHERE email = $1
`
	list, err = u.connPool.Exec(ctx, queryEmail, email)
	if err != nil {
		return fmt.Errorf("postgres: could not update user info: %w", err)
	}
	if list.RowsAffected() != 0 && user.Email != email {
		return fmt.Errorf("postgres: could not update user info: There is another user with this email")
	}

	const queryUpdate = `
    UPDATE users
    SET user_name = $2, email = $3
    WHERE id = $1
`
	_, err = u.connPool.Exec(ctx, queryUpdate, id, userName, email)
	if err != nil {
		return fmt.Errorf("postgres: could not update user info: %w", err)
	}
	return nil
}

// UserVerify implement business logic of checking if user exists in database
func (u *UserStorage) UserVerify(ctx context.Context, userName string, password string) (bool, bookly.User, error) {
	const queryAny = `
    SELECT * FROM users
	WHERE user_name = $1 AND password = $2
`

	var list pgx.Rows
	var err error
	list, err = u.connPool.Query(ctx, queryAny, userName, password)

	if err != nil {
		return false, bookly.User{}, fmt.Errorf("postgres: could not retrieve user info: %w", err)
	}

	user := bookly.User{}

	count := 0

	defer list.Close()
	for list.Next() {
		var pass string
		// todo: find better way to ignore pass if exists

		errScan := list.Scan(&user.ID, &user.FirstName, &user.Surname, &user.Email,
			&user.UserName, &pass, &user.HotelID, &user.UserRole)
		if errScan != nil {
			return false, bookly.User{}, fmt.Errorf("postgres: could not retrieve hotel's offers: %w", err)
		}

		count++
	}
	errFinal := list.Err()
	if errFinal != nil {
		return false, bookly.User{}, fmt.Errorf("postgres: could not retrieve hotel's offers: %w", err)
	}
	if count == 0 {
		return false, bookly.User{}, nil
	}
	return true, user, nil
}
